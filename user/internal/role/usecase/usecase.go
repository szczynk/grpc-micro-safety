package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"
	"user/internal/domain"
	"user/internal/models"
	"user/pkg/logger"
	"user/pkg/utils"

	"github.com/opentracing/opentracing-go"
)

// Role UseCase
type roleUseCase struct {
	logger        logger.Logger
	roleRepo      domain.RoleRepository
	roleRedisRepo domain.RoleRedisRepository
}

// New Role UseCase
func NewRoleUseCase(logger logger.Logger, roleRepo domain.RoleRepository, roleRedisRepo domain.RoleRedisRepository) domain.RoleUseCase {
	return &roleUseCase{logger: logger, roleRepo: roleRepo, roleRedisRepo: roleRedisRepo}
}

// *Command

// create new role
func (u *roleUseCase) CreateRole(ctx context.Context, role *models.Role) (*models.Role, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RoleUseCase.CreateRole")
	defer span.Finish()

	newRole, err := u.roleRepo.CreateRole(ctx, role)
	if err != nil {
		u.logger.Errorf("roleRepo.CreateRole: %v", err)
		return nil, fmt.Errorf("roleRepo.CreateRole: %v", err)
	}

	return newRole, nil
}

func (u *roleUseCase) DeleteByID(ctx context.Context, ID uint32) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RoleUseCase.DeleteByID")
	defer span.Finish()

	err := u.roleRepo.DeleteByID(ctx, ID)
	if err != nil {
		u.logger.Errorf("roleRepo.DeleteByID: %v", err)
		return fmt.Errorf("roleRepo.DeleteByID: %v", err)
	}

	return nil
}

// *Query

func (u *roleUseCase) Find(ctx context.Context, filters map[string]string, paginateQuery *utils.Pagination, expire time.Duration) ([]*models.Role, uint32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RoleUseCase.Find")
	defer span.Finish()

	keys := make([]string, 0, len(filters))
	for k := range filters {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var filterKey string
	parsedFilters := make(map[string]interface{}, len(filters))
	for _, k := range keys {
		if len(filters[k]) > 0 {
			filterKey += fmt.Sprintf("%v_%v-", k, filters[k])

			if k != "limit" && k != "page" && k != "sort" {
				parsedFilters[k] = filters[k]
			}
		}
	}
	filterKey = strings.TrimSuffix(filterKey, "-")

	var cachedRoles models.RolesPaginate
	cachedByte, er := u.roleRedisRepo.FindByID(ctx, "role_list:", filterKey)
	if er != nil {
		foundRoleList, totalCount, err := u.roleRepo.Find(ctx, parsedFilters, paginateQuery)
		if err != nil {
			u.logger.Errorf("roleRepo.Find: %v", err)
			return nil, 0, fmt.Errorf("roleRepo.Find: %v", err)
		}

		foundRoles := models.RolesPaginate{
			RoleList:   foundRoleList,
			TotalCount: totalCount,
		}

		foundRoleByte, err := json.Marshal(foundRoles)
		if err != nil {
			u.logger.Errorf("json.Marshal: %v", err)
			return nil, 0, err
		}

		err = u.roleRedisRepo.CreateRole(ctx, "role_list:", filterKey, foundRoleByte, expire)
		if err != nil {
			u.logger.Errorf("roleRedisRepo.CreateRole", err)
		}

		return foundRoleList, totalCount, nil
	}

	err := json.Unmarshal(cachedByte, &cachedRoles)
	if err != nil {
		u.logger.Errorf("json.Unmarshal: %v", err)
		return nil, 0, err
	}

	return cachedRoles.RoleList, cachedRoles.TotalCount, nil
}
