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

type methodsService struct {
	pb.UnimplementedMethodServiceServer
	logger   logger.Logger
	cfg      *config.Config
	methodUC domain.MethodUseCase
}

// Method service constructor
func NewMethodGRPCService(
	logger logger.Logger,
	cfg *config.Config,
	methodUC domain.MethodUseCase,
) *methodsService {
	return &methodsService{
		logger:   logger,
		cfg:      cfg,
		methodUC: methodUC,
	}
}

// *Command

func (u *methodsService) CreateMethod(ctx context.Context, r *pb.CreateMethodRequest) (*pb.CreateMethodResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "method.CreateMethod")
	defer span.Finish()

	method := strings.ToLower(r.GetMethod())

	newMethod := &models.Method{
		Method: method,
	}

	createdMethod, err := u.methodUC.CreateMethod(ctx, newMethod)
	if err != nil {
		u.logger.Errorf("methodUC.CreateMethod: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "CreateMethod: %v", err)
	}

	err = u.SendHeader(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.CreateMethodResponse{Method: u.MethodModelToProto(createdMethod)}, nil
}

func (u *methodsService) DeleteMethodById(ctx context.Context, r *pb.DeleteMethodByIdRequest) (*pb.DeleteMethodByIdResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "method.DeleteMethodByID")
	defer span.Finish()

	methodId := r.GetId()

	err := u.methodUC.DeleteByID(ctx, methodId)
	if err != nil {
		u.logger.Errorf("methodUC.DeleteByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "methodUC.DeleteByID: %v", err)
	}

	err = u.SendHeader(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteMethodByIdResponse{Message: "Method deleted successfully"}, nil
}

// *Query

func (u *methodsService) FindMethods(ctx context.Context, r *pb.FindMethodsRequest) (*pb.FindMethodsResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "method.FindMethods")
	defer span.Finish()

	paginateQuery := utils.NewPaginationQuery(r.GetLimit(), r.GetPage(), r.GetSort())

	// filter
	filters := map[string]string{
		"method": r.GetMethod(),

		"limit": strconv.FormatUint(uint64(r.GetLimit()), 10),
		"page":  strconv.FormatUint(uint64(r.GetPage()), 10),
		"sort":  r.GetSort(),
	}

	methodList, totalCount, err := u.methodUC.Find(ctx, filters, paginateQuery, u.cfg.Cache.Expire)
	if err != nil {
		u.logger.Errorf("methodUC.Find: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "methodUC.Find: %v", err)
	}

	parsedMethodList := make([]*pb.Method, 0, totalCount)
	for _, method := range methodList {
		parsedMethodList = append(parsedMethodList, u.MethodModelToProto(method))
	}

	err = u.SendHeader(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.FindMethodsResponse{
		TotalCount: totalCount,
		TotalPages: paginateQuery.GetTotalPages(totalCount),
		Page:       paginateQuery.GetPage(),
		Limit:      paginateQuery.GetLimit(),
		HasMore:    paginateQuery.GetHasMore(totalCount),
		Methods:    parsedMethodList,
	}, nil
}

// * Utils

func (u *methodsService) MethodModelToProto(method *models.Method) *pb.Method {
	return &pb.Method{
		Id:        uint32(method.ID),
		Method:    method.Method,
		CreatedAt: timestamppb.New(method.CreatedAt),
		UpdatedAt: timestamppb.New(method.UpdatedAt),
	}
}
