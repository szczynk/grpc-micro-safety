package domain

import (
	"context"
	"safety/internal/models"
	"safety/pkg/utils"
	"time"
)

// Schedule Repository
type ScheduleRepository interface {
	CreateSchedule(ctx context.Context, schedule *models.CreateSchedule) error
	UpdateByID(ctx context.Context, ID uint32, updates models.Schedule) (*models.Schedule, error)
	DeleteByID(ctx context.Context, ID uint32) error

	Find(ctx context.Context, filters map[string]interface{}, paginateQuery *utils.Pagination) ([]*models.Schedule, uint32, error)
	FindByID(ctx context.Context, ID uint32) (*models.Schedule, error)
}

// Schedule Redis Repository
type ScheduleRedisRepository interface {
	CreateSchedule(ctx context.Context, key string, value string, dataBytes []byte, expire time.Duration) error
	DeleteByID(ctx context.Context, key string, value string) error

	FindByID(ctx context.Context, key string, value string) ([]byte, error)
}

// Schedule Usecase
type ScheduleUseCase interface {
	CreateSchedule(ctx context.Context, schedule *models.CreateSchedule) error
	UpdateByID(ctx context.Context, ID uint32, updates models.Schedule) (*models.Schedule, error)
	DeleteByID(ctx context.Context, ID uint32) error

	Find(ctx context.Context, filters map[string]string, paginateQuery *utils.Pagination, expire time.Duration) ([]*models.Schedule, uint32, error)
	FindByID(ctx context.Context, ID uint32, expire time.Duration) (*models.Schedule, error)
}
