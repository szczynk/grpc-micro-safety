package domain

import (
	"context"
	"mail/internal/models"
	"time"

	"github.com/google/uuid"
)

// User Repository
type UserRepository interface {
	UpdateByID(ctx context.Context, ID uuid.UUID, updates models.User) (*models.User, error)
}

// User Redis Repository
type UserRedisRepository interface {
	CreateUser(ctx context.Context, key string, value string, dataBytes []byte, expire time.Duration) error
	DeleteByID(ctx context.Context, key string, value string) error

	FindByID(ctx context.Context, key string, value string) ([]byte, error)
}

// User Usecase
type UserUseCase interface {
	UpdateByID(ctx context.Context, ID uuid.UUID, updates models.User) (*models.User, error)
}
