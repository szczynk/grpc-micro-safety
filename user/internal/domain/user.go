package domain

import (
	"context"
	"time"
	"user/internal/models"
	"user/pkg/utils"

	"github.com/google/uuid"
)

// User Repository
type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	UpdateByID(ctx context.Context, ID uuid.UUID, updates models.User) (*models.User, error)
	UpdateByEmail(ctx context.Context, email string, updates map[string]interface{}) (*models.User, error)
	DeleteByID(ctx context.Context, ID uuid.UUID) error

	Find(ctx context.Context, filters map[string]interface{}, paginateQuery *utils.Pagination) ([]*models.User, uint32, error)
	FindByID(ctx context.Context, ID uuid.UUID) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByUsername(ctx context.Context, username string) (*models.User, error)
	FindByVerificationCode(ctx context.Context, code string) (*models.User, error)
	FindByResetPasswordToken(ctx context.Context, token string) (*models.User, error)
}

// User Redis Repository
type UserRedisRepository interface {
	CreateUser(ctx context.Context, key string, value string, dataBytes []byte, expire time.Duration) error
	DeleteByID(ctx context.Context, key string, value string) error

	FindByID(ctx context.Context, key string, value string) ([]byte, error)
}

// User Usecase
type UserUseCase interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	Login(ctx context.Context, email string, password string) (*models.User, error)
	UpdateByID(ctx context.Context, ID uuid.UUID, updates models.User) (*models.User, error)
	UpdateByEmail(ctx context.Context, email string, updates map[string]interface{}) (*models.User, error)
	DeleteByID(ctx context.Context, ID uuid.UUID) error

	Find(ctx context.Context, filters map[string]string, paginateQuery *utils.Pagination, expire time.Duration) ([]*models.User, uint32, error)
	FindByID(ctx context.Context, ID uuid.UUID, expire time.Duration) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByUsername(ctx context.Context, username string) (*models.User, error)
	FindByVerificationCode(ctx context.Context, code string) (*models.User, error)
	FindByResetPasswordToken(ctx context.Context, token string) (*models.User, error)
}
