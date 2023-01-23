package service

import (
	"context"
	"safety/config"
	"safety/internal/domain"
	"safety/internal/models"
	"safety/pb"
	"safety/pkg/grpc_errors"
	"safety/pkg/logger"
	"safety/pkg/utils"
	"strconv"
	"time"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type checksService struct {
	pb.UnimplementedCheckServiceServer
	logger       logger.Logger
	cfg          *config.Config
	attendanceUC domain.AttendanceUseCase
}

// Check service constructor
func NewCheckGRPCService(
	logger logger.Logger,
	cfg *config.Config,
	attendanceUC domain.AttendanceUseCase,
) *checksService {
	return &checksService{
		logger:       logger,
		cfg:          cfg,
		attendanceUC: attendanceUC,
	}
}

// * Command

func (u *checksService) CheckIn(ctx context.Context, r *pb.CheckInRequest) (*pb.CheckInResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "check.CheckIn")
	defer span.Finish()

	attendanceId := r.GetAttendanceId()

	temperature := r.GetTemperature()
	var checkStatus string
	if temperature >= 37.5 {
		checkStatus = "rejected"
	}

	_, err := u.attendanceUC.FindByID(ctx, attendanceId, u.cfg.Cache.Expire)
	if err != nil {
		u.logger.Errorf("attendanceUC.FindByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "attendanceUC.FindByID: %v", err)
	}

	updatedAttendance, err := u.attendanceUC.UpdateByID(ctx, attendanceId, models.Attendance{
		CheckIn:          time.Now(),
		CheckTemperature: temperature,
		CheckStatus:      checkStatus,
	})
	if err != nil {
		u.logger.Errorf("attendanceUC.UpdateByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "attendanceUC.UpdateByID: %v", err)
	}

	return &pb.CheckInResponse{CheckIn: u.AttendanceModelToProto(updatedAttendance)}, nil
}

func (u *checksService) CheckOut(ctx context.Context, r *pb.CheckOutRequest) (*pb.CheckOutResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "check.CheckOut")
	defer span.Finish()

	attendanceId := r.GetAttendanceId()

	_, err := u.attendanceUC.FindByID(ctx, attendanceId, u.cfg.Cache.Expire)
	if err != nil {
		u.logger.Errorf("attendanceUC.FindByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "attendanceUC.FindByID: %v", err)
	}

	updatedAttendance, err := u.attendanceUC.UpdateByID(ctx, attendanceId, models.Attendance{
		CheckOut: time.Now(),
	})
	if err != nil {
		u.logger.Errorf("attendanceUC.UpdateByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "attendanceUC.UpdateByID: %v", err)
	}

	return &pb.CheckOutResponse{CheckOut: u.AttendanceModelToProto(updatedAttendance)}, nil
}

// *Query

func (u *checksService) FindChecks(ctx context.Context, r *pb.FindChecksRequest) (*pb.FindChecksResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "check.FindChecks")
	defer span.Finish()

	// filter
	// TODO(Bagus): if user, userId from extractUserId, else, u can input userId
	userID, userRole, err := u.ExtractUserId(ctx)
	if err != nil {
		u.logger.Errorf("ExtractUserId: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "ExtractUserId: %v", err)
	}

	var userId string
	switch userRole {
	case "admin":
		userId = r.GetUserId()
	default:
		userId = userID
	}

	scheduleId := strconv.FormatUint(uint64(r.GetScheduleId()), 10)

	sort := r.GetSort()
	if len(sort) == 0 {
		sort = "created_at asc"
	}

	paginateQuery := utils.NewPaginationQuery(r.GetLimit(), r.GetPage(), sort)

	filters := map[string]string{
		"user_id":     userId,
		"schedule_id": scheduleId,

		"limit": strconv.FormatUint(uint64(r.GetLimit()), 10),
		"page":  strconv.FormatUint(uint64(r.GetPage()), 10),
		"sort":  sort,
	}

	checkList, totalCount, err := u.attendanceUC.Find(ctx, filters, paginateQuery, u.cfg.Cache.Expire)
	if err != nil {
		u.logger.Errorf("attendanceUC.Find: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "attendanceUC.Find: %v", err)
	}

	parsedCheckList := make([]*pb.AttendanceDetail, 0, totalCount)
	for _, check := range checkList {
		parsedCheckList = append(parsedCheckList, u.AttendanceDetailToProto(check))
	}

	return &pb.FindChecksResponse{
		TotalCount: totalCount,
		TotalPages: paginateQuery.GetTotalPages(totalCount),
		Page:       paginateQuery.GetPage(),
		Limit:      paginateQuery.GetLimit(),
		HasMore:    paginateQuery.GetHasMore(totalCount),
		Checks:     parsedCheckList,
	}, nil
}

func (u *checksService) FindCheckById(ctx context.Context, r *pb.FindCheckByIdRequest) (*pb.FindCheckByIdResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "check.FindCheckByID")
	defer span.Finish()

	attendanceId := r.GetAttendanceId()

	check, err := u.attendanceUC.FindByID(ctx, attendanceId, u.cfg.Cache.Expire)
	if err != nil {
		u.logger.Errorf("attendanceUC.FindByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "attendanceUC.FindByID: %v", err)
	}

	return &pb.FindCheckByIdResponse{Check: u.AttendanceDetailToProto(check)}, nil
}

// * Utils

func (u *checksService) AttendanceModelToProto(attendance *models.Attendance) *pb.Attendance {
	return &pb.Attendance{
		Id:               uint32(attendance.ID),
		UserId:           attendance.UserID.String(),
		ScheduleId:       attendance.ScheduleID,
		ImageUrl:         attendance.ImageUrl,
		Description:      attendance.Description,
		AdminUsername:    attendance.AdminUsername,
		Status:           attendance.Status,
		StatusAt:         timestamppb.New(attendance.StatusAt),
		StatusInfo:       attendance.StatusInfo,
		CheckTemperature: attendance.CheckTemperature,
		CheckStatus:      attendance.CheckStatus,
		CheckIn:          timestamppb.New(attendance.CheckIn),
		CheckOut:         timestamppb.New(attendance.CheckOut),
		CreatedAt:        timestamppb.New(attendance.CreatedAt),
		UpdatedAt:        timestamppb.New(attendance.UpdatedAt),
	}
}

func (u *checksService) AttendanceDetailToProto(attendance *models.Attendance) *pb.AttendanceDetail {
	return &pb.AttendanceDetail{
		Id:               uint32(attendance.ID),
		UserId:           attendance.UserID.String(),
		UserUsername:     attendance.UserUsername,
		UserAvatar:       attendance.UserAvatar,
		ScheduleId:       attendance.ScheduleID,
		ScheduleDate:     timestamppb.New(attendance.ScheduleDate),
		OfficeId:         attendance.OfficeID,
		OfficeName:       attendance.OfficeName,
		ImageUrl:         attendance.ImageUrl,
		Description:      attendance.Description,
		AdminUsername:    attendance.AdminUsername,
		Status:           attendance.Status,
		StatusAt:         timestamppb.New(attendance.StatusAt),
		StatusInfo:       attendance.StatusInfo,
		CheckTemperature: attendance.CheckTemperature,
		CheckStatus:      attendance.CheckStatus,
		CheckIn:          timestamppb.New(attendance.CheckIn),
		CheckOut:         timestamppb.New(attendance.CheckOut),
		CreatedAt:        timestamppb.New(attendance.CreatedAt),
		UpdatedAt:        timestamppb.New(attendance.UpdatedAt),
	}
}
