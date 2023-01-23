package domain

import (
	"context"
	"safety/internal/models"
	"safety/pkg/utils"
	"time"
)

// Attendance Repository
type AttendanceRepository interface {
	CreateAttendance(ctx context.Context, attendance *models.Attendance) (*models.Attendance, error)
	UpdateByID(ctx context.Context, ID uint32, updates models.Attendance) (*models.Attendance, error)
	DeleteByID(ctx context.Context, ID uint32) error

	Find(ctx context.Context, filters map[string]interface{}, paginateQuery *utils.Pagination) ([]*models.Attendance, uint32, error)
	FindByID(ctx context.Context, ID uint32) (*models.Attendance, error)
}

// Attendance Redis Repository
type AttendanceRedisRepository interface {
	CreateAttendance(ctx context.Context, key string, value string, dataBytes []byte, expire time.Duration) error
	DeleteByID(ctx context.Context, key string, value string) error

	FindByID(ctx context.Context, key string, value string) ([]byte, error)
}

// Attendance Usecase
type AttendanceUseCase interface {
	CreateAttendance(ctx context.Context, attendance *models.Attendance) (*models.Attendance, error)
	UpdateByID(ctx context.Context, ID uint32, updates models.Attendance) (*models.Attendance, error)
	DeleteByID(ctx context.Context, ID uint32) error

	Find(ctx context.Context, filters map[string]string, paginateQuery *utils.Pagination, expire time.Duration) ([]*models.Attendance, uint32, error)
	FindByID(ctx context.Context, ID uint32, expire time.Duration) (*models.Attendance, error)
}
