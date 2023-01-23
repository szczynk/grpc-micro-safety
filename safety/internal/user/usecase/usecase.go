package usecase

import (
	"context"
	"fmt"
	"safety/internal/domain"
	"safety/internal/models"
	"safety/pkg/grpc_errors"
	"safety/pkg/logger"
	"strings"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
)

// User UseCase
type userUseCase struct {
	logger   logger.Logger
	userRepo domain.UserRepository
}

// New User UseCase
func NewUserUseCase(logger logger.Logger, userRepo domain.UserRepository) domain.UserUseCase {
	return &userUseCase{logger: logger, userRepo: userRepo}
}

// *Command

// create new user
func (u *userUseCase) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserUseCase.CreateUser")
	defer span.Finish()

	newUser, err := u.userRepo.CreateUser(ctx, user)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "idx_users_email"):
			return nil, grpc_errors.ErrEmailExists
		case strings.Contains(err.Error(), "idx_users_username"):
			return nil, grpc_errors.ErrUsernameExists
		default:
			u.logger.Errorf("userRepo.CreateUser: %v", err)
			return nil, fmt.Errorf("userRepo.CreateUser: %v", err)
		}
	}

	return newUser, nil
}

func (u *userUseCase) UpdateByID(ctx context.Context, ID uuid.UUID, updates models.User) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserUseCase.UpdateByID")
	defer span.Finish()

	updatedUser, err := u.userRepo.UpdateByID(ctx, ID, updates)
	if err != nil {
		u.logger.Errorf("userRepo.UpdateByID: %v", err)
		return nil, fmt.Errorf("userRepo.UpdateByID: %v", err)
	}

	return updatedUser, err
}

func (u *userUseCase) UpdateByEmail(ctx context.Context, email string, updates map[string]interface{}) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserUseCase.UpdateByEmail")
	defer span.Finish()

	user, err := u.userRepo.UpdateByEmail(ctx, email, updates)
	if err != nil {
		u.logger.Errorf("userRepo.UpdateByEmail: %v", err)
		return nil, fmt.Errorf("userRepo.UpdateByEmail: %v", err)
	}

	return user, err
}

func (u *userUseCase) DeleteByID(ctx context.Context, ID uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserUseCase.DeleteByID")
	defer span.Finish()

	err := u.userRepo.DeleteByID(ctx, ID)
	if err != nil {
		u.logger.Errorf("userRepo.DeleteByID: %v", err)
		return fmt.Errorf("userRepo.DeleteByID: %v", err)
	}

	return nil
}
