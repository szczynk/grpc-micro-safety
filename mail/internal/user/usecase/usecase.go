package usecase

import (
	"context"
	"fmt"
	"mail/internal/domain"
	"mail/internal/models"
	"mail/pkg/logger"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
)

// User UseCase
type userUseCase struct {
	logger        logger.Logger
	userRepo      domain.UserRepository
	userRedisRepo domain.UserRedisRepository
}

// New User UseCase
func NewUserUseCase(logger logger.Logger, userRepo domain.UserRepository, userRedisRepo domain.UserRedisRepository) domain.UserUseCase {
	return &userUseCase{logger: logger, userRepo: userRepo, userRedisRepo: userRedisRepo}
}

// *Command

func (u *userUseCase) UpdateByID(ctx context.Context, ID uuid.UUID, updates models.User) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserUseCase.UpdateByID")
	defer span.Finish()

	err := u.userRedisRepo.DeleteByID(ctx, "user_id:", ID.String())
	if err != nil {
		u.logger.Errorf("userRedisRepo.DeleteByID: %v", err)
	}

	updatedUser, err := u.userRepo.UpdateByID(ctx, ID, updates)
	if err != nil {
		u.logger.Errorf("userRepo.UpdateByID: %v", err)
		return nil, fmt.Errorf("userRepo.UpdateByID: %v", err)
	}

	return updatedUser, err
}
