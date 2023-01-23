package repository

import (
	"context"
	"fmt"
	"user/internal/domain"
	"user/internal/models"
	"user/pkg/utils"

	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

// Role repository
type roleRepo struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) domain.RoleRepository {
	return &roleRepo{db: db}
}

// *Command

func (ur *roleRepo) CreateRole(ctx context.Context, role *models.Role) (*models.Role, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RoleRepo.CreateRole")
	defer span.Finish()

	err := ur.db.WithContext(ctx).Create(&role).Error
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (ur *roleRepo) DeleteByID(ctx context.Context, ID uint32) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RoleRepo.DeleteByID")
	defer span.Finish()

	role := new(models.Role)
	err := ur.db.WithContext(ctx).Delete(&role, ID).Error
	if err != nil {
		return err
	}

	return nil
}

// *Query

func (ur *roleRepo) Find(ctx context.Context, filter map[string]interface{}, paginateQuery *utils.Pagination) ([]*models.Role, uint32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RoleRepo.Find")
	defer span.Finish()

	var roles []*models.Role
	var count int64

	result := ur.db.WithContext(ctx)
	for k, v := range filter {
		v1 := fmt.Sprintf("%%%v%%", v)
		result = result.Where(k+" LIKE ?", v1)
	}

	result.Count(&count)

	rows := result.Offset(int(paginateQuery.GetOffset())).
		Limit(int(paginateQuery.GetLimit())).
		Order(paginateQuery.GetSort()).
		Find(&roles)
	if rows.Error != nil {
		return nil, 0, rows.Error
	}

	return roles, uint32(count), nil
}
