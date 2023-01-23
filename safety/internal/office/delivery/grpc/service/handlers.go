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

type officesService struct {
	pb.UnimplementedOfficeServiceServer
	logger   logger.Logger
	cfg      *config.Config
	officeUC domain.OfficeUseCase
}

// Office service constructor
func NewOfficeGRPCService(
	logger logger.Logger,
	cfg *config.Config,
	officeUC domain.OfficeUseCase,
) *officesService {
	return &officesService{
		logger:   logger,
		cfg:      cfg,
		officeUC: officeUC,
	}
}

// *Command

func (u *officesService) CreateOffice(ctx context.Context, r *pb.CreateOfficeRequest) (*pb.CreateOfficeResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "office.CreateOffice")
	defer span.Finish()

	newOffice := &models.Office{
		Name:   r.GetName(),
		Detail: r.GetDetail(),
	}

	createdOffice, err := u.officeUC.CreateOffice(ctx, newOffice)
	if err != nil {
		u.logger.Errorf("officeUC.CreateOffice: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "CreateOffice: %v", err)
	}

	return &pb.CreateOfficeResponse{Office: u.OfficeModelToProto(createdOffice)}, nil
}

func (u *officesService) UpdateOfficeById(ctx context.Context, r *pb.UpdateOfficeByIdRequest) (*pb.UpdateOfficeByIdResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "office.UpdateOfficeByID")
	defer span.Finish()

	officeId := r.GetId()

	_, err := u.officeUC.FindByID(ctx, officeId, u.cfg.Cache.Expire)
	if err != nil {
		u.logger.Errorf("officeUC.FindByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "officeUC.FindByID: %v", err)
	}

	updatedOffice, err := u.officeUC.UpdateByID(ctx, officeId, models.Office{
		Name:   r.GetName(),
		Detail: r.GetDetail(),
	})
	if err != nil {
		u.logger.Errorf("officeUC.UpdateByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "officeUC.UpdateByID: %v", err)
	}

	return &pb.UpdateOfficeByIdResponse{Office: u.OfficeModelToProto(updatedOffice)}, nil
}

func (u *officesService) DeleteOfficeById(ctx context.Context, r *pb.DeleteOfficeByIdRequest) (*pb.DeleteOfficeByIdResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "office.DeleteOfficeByID")
	defer span.Finish()

	officeId := r.GetId()

	err := u.officeUC.DeleteByID(ctx, officeId)
	if err != nil {
		u.logger.Errorf("officeUC.DeleteByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "officeUC.DeleteByID: %v", err)
	}

	return &pb.DeleteOfficeByIdResponse{Message: "Office deleted successfully"}, nil
}

// *Query

func (u *officesService) FindOffices(ctx context.Context, r *pb.FindOfficesRequest) (*pb.FindOfficesResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "office.FindOffices")
	defer span.Finish()

	// filter
	name := r.GetName()
	detail := r.GetDetail()

	paginateQuery := utils.NewPaginationQuery(r.GetLimit(), r.GetPage(), r.GetSort())

	filters := map[string]string{
		"name":   name,
		"detail": detail,

		"limit": strconv.FormatUint(uint64(r.GetLimit()), 10),
		"page":  strconv.FormatUint(uint64(r.GetPage()), 10),
		"sort":  r.GetSort(),
	}

	officeList, totalCount, err := u.officeUC.Find(ctx, filters, paginateQuery, u.cfg.Cache.Expire)
	if err != nil {
		u.logger.Errorf("officeUC.Find: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "officeUC.Find: %v", err)
	}

	parsedOfficeList := make([]*pb.Office, 0, totalCount)
	for _, office := range officeList {
		parsedOfficeList = append(parsedOfficeList, u.OfficeModelToProto(office))
	}

	return &pb.FindOfficesResponse{
		TotalCount: totalCount,
		TotalPages: paginateQuery.GetTotalPages(totalCount),
		Page:       paginateQuery.GetPage(),
		Limit:      paginateQuery.GetLimit(),
		HasMore:    paginateQuery.GetHasMore(totalCount),
		Offices:    parsedOfficeList,
	}, nil
}

func (u *officesService) FindOfficeById(ctx context.Context, r *pb.FindOfficeByIdRequest) (*pb.FindOfficeByIdResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "office.FindOfficeByID")
	defer span.Finish()

	officeId := r.GetId()

	office, err := u.officeUC.FindByID(ctx, officeId, u.cfg.Cache.Expire)
	if err != nil {
		u.logger.Errorf("officeUC.FindByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "officeUC.FindByID: %v", err)
	}

	return &pb.FindOfficeByIdResponse{Office: u.OfficeModelToProto(office)}, nil
}

// * Utils

func (u *officesService) OfficeModelToProto(office *models.Office) *pb.Office {
	return &pb.Office{
		Id:        uint32(office.ID),
		Name:      office.Name,
		Detail:    office.Detail,
		CreatedAt: timestamppb.New(office.CreatedAt),
		UpdatedAt: timestamppb.New(office.UpdatedAt),
	}
}
