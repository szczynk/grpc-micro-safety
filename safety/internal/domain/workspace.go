package domain

import (
	"context"
	"safety/internal/models"
	"safety/pkg/utils"
	"time"

	"github.com/google/uuid"
)

// Workspace Repository
type WorkspaceRepository interface {
	CreateWorkspace(ctx context.Context, workspace *models.Workspace) (*models.Workspace, error)
	DeleteByUserID(ctx context.Context, UserID uuid.UUID) error

	Find(ctx context.Context, filters map[string]interface{}, paginateQuery *utils.Pagination) ([]*models.User, uint32, error)
}

// Workspace Redis Repository
type WorkspaceRedisRepository interface {
	CreateWorkspace(ctx context.Context, key string, value string, dataBytes []byte, expire time.Duration) error
	DeleteByID(ctx context.Context, key string, value string) error

	FindByID(ctx context.Context, key string, value string) ([]byte, error)
}

// Workspace Usecase
type WorkspaceUseCase interface {
	CreateWorkspace(ctx context.Context, workspace *models.Workspace) (*models.Workspace, error)
	DeleteByUserID(ctx context.Context, UserID uuid.UUID) error

	Find(ctx context.Context, filters map[string]string, paginateQuery *utils.Pagination, expire time.Duration) ([]*models.User, uint32, error)
}
