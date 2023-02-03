package repository

import (
	"context"
	"fmt"
	"safety/internal/domain"
	"safety/internal/models"
	"safety/pkg/utils"
	"time"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (ur *userRepo) UpdateByID(ctx context.Context, ID uuid.UUID, updates models.User) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserRepo.UpdateByID")
	defer span.Finish()

	user := new(models.User)
	err := ur.db.WithContext(ctx).Model(&user).Clauses(clause.Returning{}).
		Where("id = ?", ID.String()).Updates(updates).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepo) UpdateByEmail(ctx context.Context, email string, updates map[string]interface{}) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserRepo.UpdateByEmail")
	defer span.Finish()

	user := new(models.User)
	err := ur.db.WithContext(ctx).Model(&user).Clauses(clause.Returning{}).
		Where("email = ?", email).Updates(updates).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepo) DeleteByID(ctx context.Context, ID uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserRepo.DeleteByID")
	defer span.Finish()

	user := new(models.User)
	err := ur.db.WithContext(ctx).Delete(&user, ID.String()).Error
	if err != nil {
		return err
	}

	return nil
}

// *Query

func (ur *userRepo) Find(ctx context.Context, filter map[string]interface{}, paginateQuery *utils.Pagination) ([]*models.User, uint32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserRepo.Find")
	defer span.Finish()

	var users []*models.User
	var count int64

	result := ur.db.WithContext(ctx)
	for k, v := range filter {
		if k == "verified" {
			result = result.Where(k+" = ?", v)
		} else {
			v1 := fmt.Sprintf("%%%v%%", v)
			result = result.Where(k+" LIKE ?", v1)
		}
	}

	result.Count(&count)

	row := result.Offset(int(paginateQuery.GetOffset())).
		Limit(int(paginateQuery.GetLimit())).
		Order(paginateQuery.GetSort()).
		Find(&users)
	if row.Error != nil {
		return nil, 0, row.Error
	}

	return users, uint32(count), nil
}

func (ur *userRepo) FindByID(ctx context.Context, ID uuid.UUID) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserRepo.FindByID")
	defer span.Finish()

	user := new(models.User)
	err := ur.db.WithContext(ctx).First(&user, "id = ?", ID.String()).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserRepo.FindByEmail")
	defer span.Finish()

	user := new(models.User)
	err := ur.db.WithContext(ctx).Where("email = ?", email).Limit(1).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepo) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserRepo.FindByUsername")
	defer span.Finish()

	user := new(models.User)
	err := ur.db.WithContext(ctx).Where("username = ?", username).Take(&user).Error
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
