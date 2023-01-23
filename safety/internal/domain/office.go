package domain

import (
	"context"
	"safety/internal/models"
	"safety/pkg/utils"
	"time"
)

// Office Repository
type OfficeRepository interface {
	CreateOffice(ctx context.Context, office *models.Office) (*models.Office, error)
	UpdateByID(ctx context.Context, ID uint32, updates models.Office) (*models.Office, error)
	DeleteByID(ctx context.Context, ID uint32) error

	Find(ctx context.Context, filters map[string]interface{}, paginateQuery *utils.Pagination) ([]*models.Office, uint32, error)
	FindByID(ctx context.Context, ID uint32) (*models.Office, error)
}

// Office Redis Repository
type OfficeRedisRepository interface {
	CreateOffice(ctx context.Context, key string, value string, dataBytes []byte, expire time.Duration) error
	DeleteByID(ctx context.Context, key string, value string) error

	FindByID(ctx context.Context, key string, value string) ([]byte, error)
}

// Office Usecase
type OfficeUseCase interface {
	CreateOffice(ctx context.Context, office *models.Office) (*models.Office, error)
	UpdateByID(ctx context.Context, ID uint32, updates models.Office) (*models.Office, error)
	DeleteByID(ctx context.Context, ID uint32) error

	Find(ctx context.Context, filters map[string]string, paginateQuery *utils.Pagination, expire time.Duration) ([]*models.Office, uint32, error)
	FindByID(ctx context.Context, ID uint32, expire time.Duration) (*models.Office, error)
}
