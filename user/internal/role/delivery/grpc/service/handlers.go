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

type rolesService struct {
	pb.UnimplementedRoleServiceServer
	logger logger.Logger
	cfg    *config.Config
	roleUC domain.RoleUseCase
}

// Role service constructor
func NewRoleGRPCService(
	logger logger.Logger,
	cfg *config.Config,
	roleUC domain.RoleUseCase,
) *rolesService {
	return &rolesService{
		logger: logger,
		cfg:    cfg,
		roleUC: roleUC,
	}
}

// *Command

func (u *rolesService) CreateRole(ctx context.Context, r *pb.CreateRoleRequest) (*pb.CreateRoleResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "role.CreateRole")
	defer span.Finish()

	role := strings.ToLower(r.GetRole())

	newRole := &models.Role{
		Role: role,
	}

	createdRole, err := u.roleUC.CreateRole(ctx, newRole)
	if err != nil {
		u.logger.Errorf("roleUC.CreateRole: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "CreateRole: %v", err)
	}

	err = u.SendHeader(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.CreateRoleResponse{Role: u.RoleModelToProto(createdRole)}, nil
}

func (u *rolesService) DeleteRoleById(ctx context.Context, r *pb.DeleteRoleByIdRequest) (*pb.DeleteRoleByIdResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "role.DeleteRoleByID")
	defer span.Finish()

	roleId := r.GetId()

	err := u.roleUC.DeleteByID(ctx, roleId)
	if err != nil {
		u.logger.Errorf("roleUC.DeleteByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "roleUC.DeleteByID: %v", err)
	}

	err = u.SendHeader(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteRoleByIdResponse{Message: "Role deleted successfully"}, nil
}

// *Query

func (u *rolesService) FindRoles(ctx context.Context, r *pb.FindRolesRequest) (*pb.FindRolesResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "role.FindRoles")
	defer span.Finish()

	paginateQuery := utils.NewPaginationQuery(r.GetLimit(), r.GetPage(), r.GetSort())

	// filter
	filters := map[string]string{
		"role": strings.ToLower(r.GetRole()),

		"limit": strconv.FormatUint(uint64(r.GetLimit()), 10),
		"page":  strconv.FormatUint(uint64(r.GetPage()), 10),
		"sort":  r.GetSort(),
	}

	roleList, totalCount, err := u.roleUC.Find(ctx, filters, paginateQuery, u.cfg.Cache.Expire)
	if err != nil {
		u.logger.Errorf("roleUC.Find: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "roleUC.Find: %v", err)
	}

	parsedRoleList := make([]*pb.Role, 0, totalCount)
	for _, role := range roleList {
		parsedRoleList = append(parsedRoleList, u.RoleModelToProto(role))
	}

	err = u.SendHeader(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.FindRolesResponse{
		TotalCount: totalCount,
		TotalPages: paginateQuery.GetTotalPages(totalCount),
		Page:       paginateQuery.GetPage(),
		Limit:      paginateQuery.GetLimit(),
		HasMore:    paginateQuery.GetHasMore(totalCount),
		Roles:      parsedRoleList,
	}, nil
}

// * Utils

func (u *rolesService) RoleModelToProto(role *models.Role) *pb.Role {
	return &pb.Role{
		Id:        uint32(role.ID),
		Role:      role.Role,
		CreatedAt: timestamppb.New(role.CreatedAt),
		UpdatedAt: timestamppb.New(role.UpdatedAt),
	}
}
