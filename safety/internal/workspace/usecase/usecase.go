package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"safety/internal/domain"
	"safety/internal/models"
	"safety/pkg/logger"
	"safety/pkg/utils"
	"sort"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
)

// Workspace UseCase
type workspaceUseCase struct {
	logger             logger.Logger
	workspaceRepo      domain.WorkspaceRepository
	workspaceRedisRepo domain.WorkspaceRedisRepository
}

// New Workspace UseCase
func NewWorkspaceUseCase(logger logger.Logger, workspaceRepo domain.WorkspaceRepository, workspaceRedisRepo domain.WorkspaceRedisRepository) domain.WorkspaceUseCase {
	return &workspaceUseCase{logger: logger, workspaceRepo: workspaceRepo, workspaceRedisRepo: workspaceRedisRepo}
}

// *Command

// create new workspace
func (u *workspaceUseCase) CreateWorkspace(ctx context.Context, workspace *models.Workspace) (*models.Workspace, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "WorkspaceUseCase.CreateWorkspace")
	defer span.Finish()

	newWorkspace, err := u.workspaceRepo.CreateWorkspace(ctx, workspace)
	if err != nil {
		u.logger.Errorf("workspaceRepo.CreateWorkspace: %v", err)
		return nil, fmt.Errorf("workspaceRepo.CreateWorkspace: %v", err)
	}

	return newWorkspace, nil
}

func (u *workspaceUseCase) DeleteByUserID(ctx context.Context, UserID uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "WorkspaceUseCase.DeleteByID")
	defer span.Finish()

	err := u.workspaceRepo.DeleteByUserID(ctx, UserID)
	if err != nil {
		u.logger.Errorf("workspaceRepo.DeleteByUserID: %v", err)
		return fmt.Errorf("workspaceRepo.DeleteByUserID: %v", err)
	}

	return nil
}

// *Query

func (u *workspaceUseCase) Find(ctx context.Context, filters map[string]string, paginateQuery *utils.Pagination, expire time.Duration) ([]*models.User, uint32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "WorkspaceUseCase.Find")
	defer span.Finish()

	keys := make([]string, 0, len(filters))
	for k := range filters {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var filterKey string
	parsedFilters := make(map[string]interface{}, len(filters))
	for _, k := range keys {
		if len(filters[k]) > 0 && filters[k] != "0" {
			filterKey += fmt.Sprintf("%v_%v-", k, filters[k])

			if k != "limit" && k != "page" && k != "sort" {
				parsedFilters[k] = filters[k]
			}
		}
	}
	filterKey = strings.TrimSuffix(filterKey, "-")

	var cachedWorkspaces models.WorkspacesPaginate
	cachedByte, er := u.workspaceRedisRepo.FindByID(ctx, "workspace_list:", filterKey)
	if er != nil {
		foundWorkspaceList, totalCount, err := u.workspaceRepo.Find(ctx, parsedFilters, paginateQuery)
		if err != nil {
			u.logger.Errorf("workspaceRepo.Find: %v", err)
			return nil, 0, fmt.Errorf("workspaceRepo.Find: %v", err)
		}

		foundWorkspaces := models.WorkspacesPaginate{
			WorkspaceList: foundWorkspaceList,
			TotalCount:    totalCount,
		}

		foundWorkspaceByte, err := json.Marshal(foundWorkspaces)
		if err != nil {
			u.logger.Errorf("json.Marshal: %v", err)
			return nil, 0, err
		}

		err = u.workspaceRedisRepo.CreateWorkspace(ctx, "workspace_list:", filterKey, foundWorkspaceByte, expire)
		if err != nil {
			u.logger.Errorf("workspaceRedisRepo.CreateWorkspace", err)
		}

		return foundWorkspaceList, totalCount, nil
	}

	err := json.Unmarshal(cachedByte, &cachedWorkspaces)
	if err != nil {
		u.logger.Errorf("json.Unmarshal: %v", err)
		return nil, 0, err
	}

	return cachedWorkspaces.WorkspaceList, cachedWorkspaces.TotalCount, nil
}
