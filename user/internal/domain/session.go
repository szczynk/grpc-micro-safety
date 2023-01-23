package domain

import (
	"context"
	"time"
	"user/internal/models"

	"github.com/google/uuid"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, maxAge time.Duration, session *models.Session) error
	DeleteByID(ctx context.Context, ID uuid.UUID) error

	FindByID(ctx context.Context, ID uuid.UUID) (*models.Session, error)
}

type SessionUseCase interface {
	CreateSession(ctx context.Context, maxAge time.Duration, session *models.Session) error
	DeleteByID(ctx context.Context, ID uuid.UUID) error

	GetSessionByID(ctx context.Context, ID uuid.UUID) (*models.Session, error)
}
