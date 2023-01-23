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

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type schedulesService struct {
	pb.UnimplementedScheduleServiceServer
	logger     logger.Logger
	cfg        *config.Config
	scheduleUC domain.ScheduleUseCase
}

// Schedule service constructor
func NewScheduleGRPCService(
	logger logger.Logger,
	cfg *config.Config,
	scheduleUC domain.ScheduleUseCase,
) *schedulesService {
	return &schedulesService{
		logger:     logger,
		cfg:        cfg,
		scheduleUC: scheduleUC,
	}
}

// *Command

func (u *schedulesService) CreateSchedule(ctx context.Context, r *pb.CreateScheduleRequest) (*pb.CreateScheduleResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "schedule.CreateSchedule")
	defer span.Finish()

	newSchedule := &models.CreateSchedule{
		OfficeID:      r.GetOfficeId(),
		TotalCapacity: r.GetTotalCapacity(),
		Month:         r.GetMonth(),
		Year:          r.GetYear(),
	}

	err := u.scheduleUC.CreateSchedule(ctx, newSchedule)
	if err != nil {
		u.logger.Errorf("scheduleUC.CreateSchedule: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "CreateSchedule: %v", err)
	}

	// schedule return 31 id so it return message
	return &pb.CreateScheduleResponse{Message: "Schedules created successfully"}, nil
}

func (u *schedulesService) UpdateScheduleById(ctx context.Context, r *pb.UpdateScheduleByIdRequest) (*pb.UpdateScheduleByIdResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "schedule.UpdateScheduleByID")
	defer span.Finish()

	scheduleId := r.GetId()

	_, err := u.scheduleUC.FindByID(ctx, scheduleId, u.cfg.Cache.Expire)
	if err != nil {
		u.logger.Errorf("scheduleUC.FindByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "scheduleUC.FindByID: %v", err)
	}

	updatedSchedule, err := u.scheduleUC.UpdateByID(ctx, scheduleId, models.Schedule{
		TotalCapacity: r.GetTotalCapacity(),
	})
	if err != nil {
		u.logger.Errorf("scheduleUC.UpdateByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "scheduleUC.UpdateByID: %v", err)
	}

	return &pb.UpdateScheduleByIdResponse{Schedule: u.ScheduleModelToProto(updatedSchedule)}, nil
}

func (u *schedulesService) DeleteScheduleById(ctx context.Context, r *pb.DeleteScheduleByIdRequest) (*pb.DeleteScheduleByIdResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "schedule.DeleteScheduleByID")
	defer span.Finish()

	scheduleId := r.GetId()

	err := u.scheduleUC.DeleteByID(ctx, scheduleId)
	if err != nil {
		u.logger.Errorf("scheduleUC.DeleteByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "scheduleUC.DeleteByID: %v", err)
	}

	return &pb.DeleteScheduleByIdResponse{Message: "Schedule deleted successfully"}, nil
}

// *Query

func (u *schedulesService) FindSchedules(ctx context.Context, r *pb.FindSchedulesRequest) (*pb.FindSchedulesResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "schedule.FindSchedules")
	defer span.Finish()

	// filter
	officeId := strconv.FormatUint(uint64(r.GetOfficeId()), 10)
	month := strconv.FormatUint(uint64(r.GetMonth()), 10)
	year := strconv.FormatUint(uint64(r.GetYear()), 10)

	sort := r.GetSort()
	if len(sort) == 0 {
		sort = "created_at asc"
	}

	paginateQuery := utils.NewPaginationQuery(r.GetLimit(), r.GetPage(), sort)

	filters := map[string]string{
		"office_id": officeId,
		"month":     month,
		"year":      year,

		"limit": strconv.FormatUint(uint64(r.GetLimit()), 10),
		"page":  strconv.FormatUint(uint64(r.GetPage()), 10),
		"sort":  sort,
	}

	scheduleList, totalCount, err := u.scheduleUC.Find(ctx, filters, paginateQuery, u.cfg.Cache.Expire)
	if err != nil {
		u.logger.Errorf("scheduleUC.Find: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "scheduleUC.Find: %v", err)
	}

	parsedScheduleList := make([]*pb.ScheduleWithOffice, 0, totalCount)
	for _, schedule := range scheduleList {
		parsedScheduleList = append(parsedScheduleList, u.ScheduleWithOfficeToProto(schedule))
	}

	return &pb.FindSchedulesResponse{
		TotalCount: totalCount,
		TotalPages: paginateQuery.GetTotalPages(totalCount),
		Page:       paginateQuery.GetPage(),
		Limit:      paginateQuery.GetLimit(),
		HasMore:    paginateQuery.GetHasMore(totalCount),
		Schedules:  parsedScheduleList,
	}, nil
}

// TODO(Bagus): Get User by attendance
func (u *schedulesService) FindScheduleById(ctx context.Context, r *pb.FindScheduleByIdRequest) (*pb.FindScheduleByIdResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "schedule.FindScheduleByID")
	defer span.Finish()

	scheduleId := r.GetId()

	schedule, err := u.scheduleUC.FindByID(ctx, scheduleId, u.cfg.Cache.Expire)
	if err != nil {
		u.logger.Errorf("scheduleUC.FindByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "scheduleUC.FindByID: %v", err)
	}

	return &pb.FindScheduleByIdResponse{Schedule: u.ScheduleWithOfficeToProto(schedule)}, nil
}

// * Utils

func (u *schedulesService) ScheduleModelToProto(schedule *models.Schedule) *pb.Schedule {
	return &pb.Schedule{
		Id:            uint32(schedule.ID),
		Capacity:      schedule.Capacity,
		TotalCapacity: schedule.TotalCapacity,
		OfficeId:      schedule.OfficeID,
		Date:          timestamppb.New(schedule.Date),
		CreatedAt:     timestamppb.New(schedule.CreatedAt),
		UpdatedAt:     timestamppb.New(schedule.UpdatedAt),
	}
}

func (u *schedulesService) ScheduleWithOfficeToProto(schedule *models.Schedule) *pb.ScheduleWithOffice {
	return &pb.ScheduleWithOffice{
		Id:            uint32(schedule.ID),
		OfficeId:      schedule.OfficeID,
		OfficeName:    schedule.OfficeName,
		TotalCapacity: schedule.TotalCapacity,
		Capacity:      schedule.Capacity,
		Date:          timestamppb.New(schedule.Date),
		CreatedAt:     timestamppb.New(schedule.CreatedAt),
		UpdatedAt:     timestamppb.New(schedule.UpdatedAt),
	}
}
