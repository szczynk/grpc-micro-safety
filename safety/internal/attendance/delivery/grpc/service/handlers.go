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

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type attendancesService struct {
	pb.UnimplementedAttendanceServiceServer
	logger        logger.Logger
	cfg           *config.Config
	attendanceUC  domain.AttendanceUseCase
	certificateUC domain.CertificateUseCase
	scheduleUC    domain.ScheduleUseCase
}

// Attendance service constructor
func NewAttendanceGRPCService(
	logger logger.Logger,
	cfg *config.Config,
	attendanceUC domain.AttendanceUseCase,
	certificateUC domain.CertificateUseCase,
	scheduleUC domain.ScheduleUseCase,
) *attendancesService {
	return &attendancesService{
		logger:        logger,
		cfg:           cfg,
		attendanceUC:  attendanceUC,
		certificateUC: certificateUC,
		scheduleUC:    scheduleUC,
	}
}

// *Command

func (u *attendancesService) CreateAttendance(ctx context.Context, r *pb.CreateAttendanceRequest) (*pb.CreateAttendanceResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "attendance.CreateAttendance")
	defer span.Finish()

	parsedUserID, err := uuid.Parse(r.GetUserId())
	if err != nil {
		u.logger.Errorf("ConvertUserID: %v", err)
		return nil, status.Errorf(codes.Internal, "ConvertUserID: %v", err)
	}

	scheduleId := r.GetScheduleId()

	newAttendance := &models.Attendance{
		UserID:      parsedUserID,
		ScheduleID:  scheduleId,
		Description: r.GetDescription(),
		ImageUrl:    r.GetImageUrl(),
	}

	// TODO(Bagus): Check User's Vaccine Status
	certificate, err := u.certificateUC.CountByUserID(ctx, parsedUserID, u.cfg.Cache.Expire)
	if err != nil {
		u.logger.Errorf("certificateUC.CountByUserID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "The vaccine certificate that has been input has not been approved by the admin; please wait.")
	}
	if certificate < 3 {
		u.logger.Error("The number of approved vaccine certificates is less than three.")
		return nil, status.Errorf(codes.InvalidArgument, "The number of approved vaccine certificates is less than three.")
	}

	// TODO(Bagus): Check is Request's date > time.Now()
	schedule, err := u.scheduleUC.FindByID(ctx, scheduleId, u.cfg.Cache.Expire)
	if err != nil {
		u.logger.Errorf("scheduleUC.FindByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "scheduleUC.FindByID: %v", err)
	}
	checkDate := time.Now().Before(schedule.Date)
	if !checkDate {
		u.logger.Error("the number of vaccine certificates is not 3")
		return nil, status.Errorf(codes.InvalidArgument, "the number of vaccine certificates is not 3")
	}

	createdAttendance, err := u.attendanceUC.CreateAttendance(ctx, newAttendance)
	if err != nil {
		u.logger.Errorf("attendanceUC.CreateAttendance: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "CreateAttendance: %v", err)
	}

	err = u.SendHeader(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.CreateAttendanceResponse{Attendance: u.AttendanceModelToProto(createdAttendance)}, nil
}

func (u *attendancesService) UpdateAttendanceById(ctx context.Context, r *pb.UpdateAttendanceByIdRequest) (*pb.UpdateAttendanceByIdResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "attendance.UpdateAttendanceByID")
	defer span.Finish()

	attendanceId := r.GetId()
	attendanceStatus := r.GetStatus()

	attendance, err := u.attendanceUC.FindByID(ctx, attendanceId, u.cfg.Cache.Expire)
	if err != nil {
		u.logger.Errorf("attendanceUC.FindByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "attendanceUC.FindByID: %v", err)
	}

	// TODO(Bagus): check capacity by scheduleId
	schedule, err := u.scheduleUC.FindByID(ctx, attendance.ScheduleID, u.cfg.Cache.Expire)
	if err != nil {
		u.logger.Errorf("scheduleUC.FindByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "scheduleUC.FindByID: %v", err)
	}
	remainingCapacity := schedule.TotalCapacity - schedule.Capacity
	if remainingCapacity < 1 {
		u.logger.Error("the remaining capacity has been used up")
		return nil, status.Errorf(codes.Internal, "the remaining capacity has been used up")
	}

	// TODO(Bagus): increment if status == "approved"
	if attendanceStatus == "approved" {
		_, err = u.scheduleUC.UpdateByID(ctx, attendance.ScheduleID, models.Schedule{
			Capacity: schedule.Capacity + 1,
		})
		if err != nil {
			u.logger.Errorf("scheduleUC.UpdateByID: %v", err)
			return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "scheduleUC.UpdateByID: %v", err)
		}
	}

	updatedAttendance, err := u.attendanceUC.UpdateByID(ctx, attendanceId, models.Attendance{
		AdminUsername: r.GetAdminUsername(),
		Status:        attendanceStatus,
		StatusAt:      time.Now(),
		StatusInfo:    r.GetStatusInfo(),
	})
	if err != nil {
		u.logger.Errorf("attendanceUC.UpdateByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "attendanceUC.UpdateByID: %v", err)
	}

	err = u.SendHeader(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateAttendanceByIdResponse{Attendance: u.AttendanceModelToProto(updatedAttendance)}, nil
}

func (u *attendancesService) DeleteAttendanceById(ctx context.Context, r *pb.DeleteAttendanceByIdRequest) (*pb.DeleteAttendanceByIdResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "attendance.DeleteAttendanceByID")
	defer span.Finish()

	attendanceId := r.GetId()

	err := u.attendanceUC.DeleteByID(ctx, attendanceId)
	if err != nil {
		u.logger.Errorf("attendanceUC.DeleteByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "attendanceUC.DeleteByID: %v", err)
	}

	err = u.SendHeader(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteAttendanceByIdResponse{Message: "Attendance deleted successfully"}, nil
}

// *Query

func (u *attendancesService) FindAttendances(ctx context.Context, r *pb.FindAttendancesRequest) (*pb.FindAttendancesResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "attendance.FindAttendances")
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
	adminUsername := r.GetAdminUsername()
	aStatus := r.GetStatus()

	sort := r.GetSort()
	if len(sort) == 0 {
		sort = "created_at asc"
	}

	paginateQuery := utils.NewPaginationQuery(r.GetLimit(), r.GetPage(), sort)

	filters := map[string]string{
		"user_id":        userId,
		"schedule_id":    scheduleId,
		"admin_username": adminUsername,
		"status":         aStatus,

		"limit": strconv.FormatUint(uint64(r.GetLimit()), 10),
		"page":  strconv.FormatUint(uint64(r.GetPage()), 10),
		"sort":  sort,
	}

	attendanceList, totalCount, err := u.attendanceUC.Find(ctx, filters, paginateQuery, u.cfg.Cache.Expire)
	if err != nil {
		u.logger.Errorf("attendanceUC.Find: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "attendanceUC.Find: %v", err)
	}

	parsedAttendanceList := make([]*pb.AttendanceDetail, 0, totalCount)
	for _, attendance := range attendanceList {
		parsedAttendanceList = append(parsedAttendanceList, u.AttendanceDetailToProto(attendance))
	}

	err = u.SendHeader(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.FindAttendancesResponse{
		TotalCount:  totalCount,
		TotalPages:  paginateQuery.GetTotalPages(totalCount),
		Page:        paginateQuery.GetPage(),
		Limit:       paginateQuery.GetLimit(),
		HasMore:     paginateQuery.GetHasMore(totalCount),
		Attendances: parsedAttendanceList,
	}, nil
}

func (u *attendancesService) FindAttendanceById(ctx context.Context, r *pb.FindAttendanceByIdRequest) (*pb.FindAttendanceByIdResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "attendance.FindAttendanceByID")
	defer span.Finish()

	attendanceId := r.GetId()

	attendance, err := u.attendanceUC.FindByID(ctx, attendanceId, u.cfg.Cache.Expire)
	if err != nil {
		u.logger.Errorf("attendanceUC.FindByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "attendanceUC.FindByID: %v", err)
	}

	err = u.SendHeader(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.FindAttendanceByIdResponse{Attendance: u.AttendanceDetailToProto(attendance)}, nil
}

// * Utils

func (u *attendancesService) AttendanceModelToProto(attendance *models.Attendance) *pb.Attendance {
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

func (u *attendancesService) AttendanceDetailToProto(attendance *models.Attendance) *pb.AttendanceDetail {
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
