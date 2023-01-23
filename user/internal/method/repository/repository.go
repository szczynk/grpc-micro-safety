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

// Method repository
type methodRepo struct {
	db *gorm.DB
}

func NewMethodRepository(db *gorm.DB) domain.MethodRepository {
	return &methodRepo{db: db}
}

// *Command

func (ur *methodRepo) CreateMethod(ctx context.Context, method *models.Method) (*models.Method, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "MethodRepo.CreateMethod")
	defer span.Finish()

	err := ur.db.WithContext(ctx).Create(&method).Error
	if err != nil {
		return nil, err
	}

	return method, nil
}

func (ur *methodRepo) DeleteByID(ctx context.Context, ID uint32) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "MethodRepo.DeleteByID")
	defer span.Finish()

	method := new(models.Method)
	err := ur.db.WithContext(ctx).Delete(&method, ID).Error
	if err != nil {
		return err
	}

	return nil
}

// *Query

func (ur *methodRepo) Find(ctx context.Context, filter map[string]interface{}, paginateQuery *utils.Pagination) ([]*models.Method, uint32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "MethodRepo.Find")
	defer span.Finish()

	var methods []*models.Method
	var count int64

	result := ur.db.WithContext(ctx).Model(&models.Method{})
	for k, v := range filter {
		v1 := fmt.Sprintf("%%%v%%", v)
		result = result.Where(k+" LIKE ?", v1)
	}

	result.Count(&count)

	rows := result.Offset(int(paginateQuery.GetOffset())).
		Limit(int(paginateQuery.GetLimit())).
		Order(paginateQuery.GetSort()).
		Find(&methods)
	if rows.Error != nil {
		return nil, 0, rows.Error
	}

	return methods, uint32(count), nil
}
