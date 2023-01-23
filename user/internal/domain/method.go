package domain

import (
	"context"
	"time"
	"user/internal/models"
	"user/pkg/utils"
)

// Method Repository
type MethodRepository interface {
	CreateMethod(ctx context.Context, method *models.Method) (*models.Method, error)
	DeleteByID(ctx context.Context, ID uint32) error

	Find(ctx context.Context, filters map[string]interface{}, paginateQuery *utils.Pagination) ([]*models.Method, uint32, error)
}

// Method Redis Repository
type MethodRedisRepository interface {
	CreateMethod(ctx context.Context, key string, value string, dataBytes []byte, expire time.Duration) error
	DeleteByID(ctx context.Context, key string, value string) error

	FindByID(ctx context.Context, key string, value string) ([]byte, error)
}

// Method Usecase
type MethodUseCase interface {
	CreateMethod(ctx context.Context, method *models.Method) (*models.Method, error)
	DeleteByID(ctx context.Context, ID uint32) error

	Find(ctx context.Context, filters map[string]string, paginateQuery *utils.Pagination, expire time.Duration) ([]*models.Method, uint32, error)
}
