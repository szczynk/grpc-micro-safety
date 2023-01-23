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
	"strings"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type workspacesService struct {
	pb.UnimplementedWorkspaceServiceServer
	logger      logger.Logger
	cfg         *config.Config
	workspaceUC domain.WorkspaceUseCase
}

// Workspace service constructor
func NewWorkspaceGRPCService(
	logger logger.Logger,
	cfg *config.Config,
	workspaceUC domain.WorkspaceUseCase,
) *workspacesService {
	return &workspacesService{
		logger:      logger,
		cfg:         cfg,
		workspaceUC: workspaceUC,
	}
}

// *Command

func (u *workspacesService) CreateWorkspace(ctx context.Context, r *pb.CreateWorkspaceRequest) (*pb.CreateWorkspaceResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "workspace.CreateWorkspace")
	defer span.Finish()

	officeId := r.GetOfficeId()
	parsedUserID, err := uuid.Parse(r.GetUserId())
	if err != nil {
		u.logger.Errorf("ConvertUserID: %v", err)
		return nil, status.Errorf(codes.Internal, "ConvertUserID: %v", err)
	}

	newWorkspace := &models.Workspace{
		OfficeID: officeId,
		UserID:   parsedUserID,
	}

	createdWorkspace, err := u.workspaceUC.CreateWorkspace(ctx, newWorkspace)
	if err != nil {
		u.logger.Errorf("workspaceUC.CreateWorkspace: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "CreateWorkspace: %v", err)
	}

	return &pb.CreateWorkspaceResponse{
		Workspace: &pb.Workspace{
			OfficeId: createdWorkspace.OfficeID,
			UserId:   createdWorkspace.UserID.String(),
		},
	}, nil
}

func (u *workspacesService) DeleteWorkspaceById(ctx context.Context, r *pb.DeleteWorkspaceByIdRequest) (*pb.DeleteWorkspaceByIdResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "workspace.DeleteWorkspaceByID")
	defer span.Finish()

	parsedUserID, err := uuid.Parse(r.GetUserId())
	if err != nil {
		u.logger.Errorf("ConvertUserID: %v", err)
		return nil, status.Errorf(codes.Internal, "ConvertUserID: %v", err)
	}

	err = u.workspaceUC.DeleteByUserID(ctx, parsedUserID)
	if err != nil {
		u.logger.Errorf("workspaceUC.DeleteByUserID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "workspaceUC.DeleteByUserID: %v", err)
	}

	return &pb.DeleteWorkspaceByIdResponse{Message: "Workspace deleted successfully"}, nil
}

// *Query

func (u *workspacesService) FindWorkspaces(ctx context.Context, r *pb.FindWorkspacesRequest) (*pb.FindWorkspacesResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "workspace.FindWorkspaces")
	defer span.Finish()

	_, userRole, err := u.ExtractUserId(ctx)
	if err != nil {
		u.logger.Errorf("ExtractUserId: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "ExtractUserId: %v", err)
	}

	// filter
	officeId := strconv.FormatUint(uint64(r.GetOfficeId()), 10)
	username := strings.TrimSpace(r.GetUsername())
	email := strings.ToLower(strings.TrimSpace(r.GetEmail()))
	role := strings.ToLower(r.GetRole())
	verified := strings.ToLower(r.GetVerified())

	paginateQuery := utils.NewPaginationQuery(r.GetLimit(), r.GetPage(), r.GetSort())

	var filters map[string]string
	switch userRole {
	case "admin":
		filters = map[string]string{
			"office_id": officeId,
			"username":  username,
			"email":     email,
			"role":      role,
			"verified":  verified,

			"limit": strconv.FormatUint(uint64(r.GetLimit()), 10),
			"page":  strconv.FormatUint(uint64(r.GetPage()), 10),
			"sort":  r.GetSort(),
		}
	default:
		filters = map[string]string{
			"office_id": officeId,

			"limit": strconv.FormatUint(uint64(r.GetLimit()), 10),
			"page":  strconv.FormatUint(uint64(r.GetPage()), 10),
			"sort":  r.GetSort(),
		}
	}

	workspaceList, totalCount, err := u.workspaceUC.Find(ctx, filters, paginateQuery, u.cfg.Cache.Expire)
	if err != nil {
		u.logger.Errorf("workspaceUC.Find: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "workspaceUC.Find: %v", err)
	}

	parsedUserList := make([]*pb.User, 0, totalCount)
	for _, user := range workspaceList {
		parsedUserList = append(parsedUserList, u.UserModelToProto(user))
	}

	return &pb.FindWorkspacesResponse{
		TotalCount: totalCount,
		TotalPages: paginateQuery.GetTotalPages(totalCount),
		Page:       paginateQuery.GetPage(),
		Limit:      paginateQuery.GetLimit(),
		HasMore:    paginateQuery.GetHasMore(totalCount),
		Users:      parsedUserList,
	}, nil
}

// * Utils

func (u *workspacesService) UserModelToProto(user *models.User) *pb.User {
	return &pb.User{
		Id:        user.ID.String(),
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		Avatar:    user.Avatar,
		Verified:  user.Verified,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}
