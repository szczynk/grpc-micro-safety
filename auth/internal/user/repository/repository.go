package repository

import (
	"auth/internal/domain"
	"auth/internal/models"
	"context"
	"time"

	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

// User repository
type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepo{db: db}
}

// *Command

func (ur *userRepo) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserRepo.CreateUser")
	defer span.Finish()

	err := ur.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepo) UpdateByEmail(ctx context.Context, email string, updates map[string]interface{}) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserRepo.UpdateByEmail")
	defer span.Finish()

	user := new(models.User)
	err := ur.db.WithContext(ctx).Model(&user).Where("email = ?", email).Updates(updates).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

// *Query

func (ur *userRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserRepo.FindByEmail")
	defer span.Finish()

	user := new(models.User)
	err := ur.db.WithContext(ctx).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepo) FindByVerificationCode(ctx context.Context, code string) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserRepo.FindByVerificationCode")
	defer span.Finish()

	user := new(models.User)
	err := ur.db.WithContext(ctx).Where("verification_code = ?", code).Take(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepo) FindByResetPasswordToken(ctx context.Context, token string) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserRepo.FindByResetPasswordToken")
	defer span.Finish()

	user := new(models.User)
	err := ur.db.WithContext(ctx).Where("password_reset_token = ? AND password_reset_at > ?", token, time.Now()).Take(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
