package domain

import (
	"context"
	"time"
	"user/internal/models"
	"user/pkg/utils"
)

// Role Repository
type RoleRepository interface {
	CreateRole(ctx context.Context, role *models.Role) (*models.Role, error)
	DeleteByID(ctx context.Context, ID uint32) error

	Find(ctx context.Context, filters map[string]interface{}, paginateQuery *utils.Pagination) ([]*models.Role, uint32, error)
}

// Role Redis Repository
type RoleRedisRepository interface {
	CreateRole(ctx context.Context, key string, value string, dataBytes []byte, expire time.Duration) error
	DeleteByID(ctx context.Context, key string, value string) error

	FindByID(ctx context.Context, key string, value string) ([]byte, error)
}

// Role Usecase
type RoleUseCase interface {
	CreateRole(ctx context.Context, role *models.Role) (*models.Role, error)
	DeleteByID(ctx context.Context, ID uint32) error

	Find(ctx context.Context, filters map[string]string, paginateQuery *utils.Pagination, expire time.Duration) ([]*models.Role, uint32, error)
}
