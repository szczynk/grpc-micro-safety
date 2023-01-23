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

// Service repository
type serviceRepo struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) domain.ServiceRepository {
	return &serviceRepo{db: db}
}

// *Command

func (ur *serviceRepo) CreateService(ctx context.Context, service *models.Service) (*models.Service, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ServiceRepo.CreateService")
	defer span.Finish()

	err := ur.db.WithContext(ctx).Create(&service).Error
	if err != nil {
		return nil, err
	}

	return service, nil
}

func (ur *serviceRepo) DeleteByID(ctx context.Context, ID uint32) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ServiceRepo.DeleteByID")
	defer span.Finish()

	service := new(models.Service)
	err := ur.db.WithContext(ctx).Delete(&service, ID).Error
	if err != nil {
		return err
	}

	return nil
}

// *Query

func (ur *serviceRepo) Find(ctx context.Context, filter map[string]interface{}, paginateQuery *utils.Pagination) ([]*models.Service, uint32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ServiceRepo.Find")
	defer span.Finish()

	var services []*models.Service
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
		Find(&services)
	if rows.Error != nil {
		return nil, 0, rows.Error
	}

	return services, uint32(count), nil
}
