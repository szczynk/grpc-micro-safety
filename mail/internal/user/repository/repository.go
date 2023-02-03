package repository

import (
	"context"
	"mail/internal/domain"
	"mail/internal/models"

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
