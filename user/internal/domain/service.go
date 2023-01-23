package domain

import (
	"context"
	"time"
	"user/internal/models"
	"user/pkg/utils"
)

// Service Repository
type ServiceRepository interface {
	CreateService(ctx context.Context, service *models.Service) (*models.Service, error)
	DeleteByID(ctx context.Context, ID uint32) error

	Find(ctx context.Context, filters map[string]interface{}, paginateQuery *utils.Pagination) ([]*models.Service, uint32, error)
}

// Service Redis Repository
type ServiceRedisRepository interface {
	CreateService(ctx context.Context, key string, value string, dataBytes []byte, expire time.Duration) error
	DeleteByID(ctx context.Context, key string, value string) error

	FindByID(ctx context.Context, key string, value string) ([]byte, error)
}

// Service Usecase
type ServiceUseCase interface {
	CreateService(ctx context.Context, service *models.Service) (*models.Service, error)
	DeleteByID(ctx context.Context, ID uint32) error

	Find(ctx context.Context, filters map[string]string, paginateQuery *utils.Pagination, expire time.Duration) ([]*models.Service, uint32, error)
}
