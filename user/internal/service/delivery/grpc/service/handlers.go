package service

import (
	"context"
	"strconv"
	"strings"
	"user/config"
	"user/internal/domain"
	"user/internal/models"
	"user/pb"
	"user/pkg/grpc_errors"
	"user/pkg/logger"
	"user/pkg/utils"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type servicesService struct {
	pb.UnimplementedServiceServiceServer
	logger    logger.Logger
	cfg       *config.Config
	serviceUC domain.ServiceUseCase
}

// Service service constructor
func NewServiceGRPCService(
	logger logger.Logger,
	cfg *config.Config,
	serviceUC domain.ServiceUseCase,
) *servicesService {
	return &servicesService{
		logger:    logger,
		cfg:       cfg,
		serviceUC: serviceUC,
	}
}

// *Command

func (u *servicesService) CreateService(ctx context.Context, r *pb.CreateServiceRequest) (*pb.CreateServiceResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.CreateService")
	defer span.Finish()

	service := strings.ToLower(r.GetService())

	newService := &models.Service{
		Service: service,
	}

	createdService, err := u.serviceUC.CreateService(ctx, newService)
	if err != nil {
		u.logger.Errorf("serviceUC.CreateService: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "CreateService: %v", err)
	}

	return &pb.CreateServiceResponse{Service: u.ServiceModelToProto(createdService)}, nil
}

func (u *servicesService) DeleteServiceById(ctx context.Context, r *pb.DeleteServiceByIdRequest) (*pb.DeleteServiceByIdResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.DeleteServiceByID")
	defer span.Finish()

	serviceId := r.GetId()

	err := u.serviceUC.DeleteByID(ctx, serviceId)
	if err != nil {
		u.logger.Errorf("serviceUC.DeleteByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "serviceUC.DeleteByID: %v", err)
	}

	return &pb.DeleteServiceByIdResponse{Message: "Service deleted successfully"}, nil
}

// *Query

func (u *servicesService) FindServices(ctx context.Context, r *pb.FindServicesRequest) (*pb.FindServicesResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.FindServices")
	defer span.Finish()

	paginateQuery := utils.NewPaginationQuery(r.GetLimit(), r.GetPage(), r.GetSort())

	// filter
	filters := map[string]string{
		"service": r.GetService(),

		"limit": strconv.FormatUint(uint64(r.GetLimit()), 10),
		"page":  strconv.FormatUint(uint64(r.GetPage()), 10),
		"sort":  r.GetSort(),
	}

	serviceList, totalCount, err := u.serviceUC.Find(ctx, filters, paginateQuery, u.cfg.Cache.Expire)
	if err != nil {
		u.logger.Errorf("serviceUC.Find: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "serviceUC.Find: %v", err)
	}

	parsedServiceList := make([]*pb.Service, 0, totalCount)
	for _, service := range serviceList {
		parsedServiceList = append(parsedServiceList, u.ServiceModelToProto(service))
	}

	return &pb.FindServicesResponse{
		TotalCount: totalCount,
		TotalPages: paginateQuery.GetTotalPages(totalCount),
		Page:       paginateQuery.GetPage(),
		Limit:      paginateQuery.GetLimit(),
		HasMore:    paginateQuery.GetHasMore(totalCount),
		Services:   parsedServiceList,
	}, nil
}

// * Utils

func (u *servicesService) ServiceModelToProto(service *models.Service) *pb.Service {
	return &pb.Service{
		Id:        uint32(service.ID),
		Service:   service.Service,
		CreatedAt: timestamppb.New(service.CreatedAt),
		UpdatedAt: timestamppb.New(service.UpdatedAt),
	}
}
