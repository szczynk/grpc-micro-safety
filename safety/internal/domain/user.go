package domain

import (
	"context"
	"safety/internal/models"

	"github.com/google/uuid"
)

// User Repository
type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	UpdateByID(ctx context.Context, ID uuid.UUID, updates models.User) (*models.User, error)
	UpdateByEmail(ctx context.Context, email string, updates map[string]interface{}) (*models.User, error)
	DeleteByID(ctx context.Context, ID uuid.UUID) error
}

// User Usecase
type UserUseCase interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	UpdateByID(ctx context.Context, ID uuid.UUID, updates models.User) (*models.User, error)
	UpdateByEmail(ctx context.Context, email string, updates map[string]interface{}) (*models.User, error)
	DeleteByID(ctx context.Context, ID uuid.UUID) error
}
