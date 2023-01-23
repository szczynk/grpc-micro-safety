package usecase

import (
	"auth/internal/domain"
	"auth/internal/models"
	"auth/pkg/logger"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
)

// Session use case
type sessionUseCase struct {
	logger      logger.Logger
	sessionRepo domain.SessionRepository
}

// New session use case constructor
func NewSessionUseCase(logger logger.Logger, sessionRepo domain.SessionRepository) domain.SessionUseCase {
	return &sessionUseCase{logger: logger, sessionRepo: sessionRepo}
}

// *Command

// Create new session
func (u *sessionUseCase) CreateSession(ctx context.Context, maxAge time.Duration, session *models.Session) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "SessionUseCase.CreateSession")
	defer span.Finish()

	return u.sessionRepo.CreateSession(ctx, maxAge, session)
}

// Delete session by id
func (u *sessionUseCase) DeleteByID(ctx context.Context, ID uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "SessionUseCase.DeleteByID")
	defer span.Finish()

	return u.sessionRepo.DeleteByID(ctx, ID)
}

// *Query

// get session by id
func (u *sessionUseCase) GetSessionByID(ctx context.Context, ID uuid.UUID) (*models.Session, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "SessionUseCase.GetSessionByID")
	defer span.Finish()

	return u.sessionRepo.FindByID(ctx, ID)
}
