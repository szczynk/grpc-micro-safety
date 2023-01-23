//go:generate mockgen -source user.go -destination ../user/mock/user.go

package domain

import (
	"auth/internal/models"
	"context"
	"time"
)

// User Repository
type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	UpdateByEmail(ctx context.Context, email string, updates map[string]interface{}) (*models.User, error)

	FindByEmail(ctx context.Context, email string) (*models.User, error)
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
	UpdateByEmail(ctx context.Context, email string, updates map[string]interface{}) (*models.User, error)

	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByVerificationCode(ctx context.Context, code string) (*models.User, error)
	FindByResetPasswordToken(ctx context.Context, token string) (*models.User, error)
}
