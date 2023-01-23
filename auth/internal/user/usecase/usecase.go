package usecase

import (
	"auth/internal/domain"
	"auth/internal/models"
	"auth/pkg/grpc_errors"
	"auth/pkg/logger"
	"auth/pkg/utils"
	"context"
	"fmt"
	"strings"

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

// Login user with email and password
func (u *userUseCase) Login(ctx context.Context, email string, password string) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserUseCase.Login")
	defer span.Finish()

	foundUser, err := u.userRepo.FindByEmail(ctx, email)
	if err != nil {
		u.logger.Errorf("userRepo.FindByEmail: %v", err)
		return nil, err
	}

	err = utils.CheckPasswordHash(foundUser.Password, password)
	if err != nil {
		return nil, grpc_errors.ErrInvalidPassword
	}

	return foundUser, err
}

func (u *userUseCase) UpdateByEmail(ctx context.Context, email string, updates map[string]interface{}) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserUseCase.UpdateByEmail")
	defer span.Finish()

	err := u.userRedisRepo.DeleteByID(ctx, "user_email", email)
	if err != nil {
		u.logger.Errorf("userRedisRepo.DeleteByID: %v", err)
	}

	user, err := u.userRepo.UpdateByEmail(ctx, email, updates)
	if err != nil {
		u.logger.Errorf("userRepo.UpdateByEmail: %v", err)
		return nil, fmt.Errorf("userRepo.UpdateByEmail: %v", err)
	}

	return user, err
}

// *Query

func (u *userUseCase) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserUseCase.FindByEmail")
	defer span.Finish()

	user, err := u.userRepo.FindByEmail(ctx, email)
	if err != nil {
		u.logger.Errorf("userRepo.FindByEmail: %v", err)
		return nil, fmt.Errorf("userRepo.FindByEmail: %v", err)
	}

	return user, err
}

func (u *userUseCase) FindByVerificationCode(ctx context.Context, code string) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserUseCase.FindByVerificationCode")
	defer span.Finish()

	user, err := u.userRepo.FindByVerificationCode(ctx, code)
	if err != nil {
		u.logger.Errorf("userRepo.FindByVerificationCode: %v", err)
		return nil, fmt.Errorf("userRepo.FindByVerificationCode: %v", err)
	}

	return user, err
}

func (u *userUseCase) FindByResetPasswordToken(ctx context.Context, code string) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserUseCase.FindByResetPasswordToken")
	defer span.Finish()

	user, err := u.userRepo.FindByResetPasswordToken(ctx, code)
	if err != nil {
		u.logger.Errorf("userRepo.FindByResetPasswordToken: %v", err)
		return nil, fmt.Errorf("userRepo.FindResetPasswordToken: %v", err)
	}

	return user, err
}
