package repository

import (
	"context"
	"fmt"
	"safety/internal/domain"
	"safety/internal/models"
	"safety/pkg/utils"

	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

// Office repository
type officeRepo struct {
	db *gorm.DB
}

func NewOfficeRepository(db *gorm.DB) domain.OfficeRepository {
	return &officeRepo{db: db}
}

// *Command

func (ur *officeRepo) CreateOffice(ctx context.Context, office *models.Office) (*models.Office, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OfficeRepo.CreateOffice")
	defer span.Finish()

	err := ur.db.WithContext(ctx).Create(&office).Error
	if err != nil {
		return nil, err
	}

	return office, nil
}

func (ur *officeRepo) UpdateByID(ctx context.Context, ID uint32, updates models.Office) (*models.Office, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OfficeRepo.UpdateByID")
	defer span.Finish()

	office := new(models.Office)
	err := ur.db.WithContext(ctx).Model(&office).Where("id = ?", ID).Updates(updates).Error
	if err != nil {
		return nil, err
	}

	return office, nil
}

func (ur *officeRepo) DeleteByID(ctx context.Context, ID uint32) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OfficeRepo.DeleteByID")
	defer span.Finish()

	office := new(models.Office)
	err := ur.db.WithContext(ctx).Delete(&office, ID).Error
	if err != nil {
		return err
	}

	return nil
}

// *Query

func (ur *officeRepo) Find(ctx context.Context, filter map[string]interface{}, paginateQuery *utils.Pagination) ([]*models.Office, uint32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OfficeRepo.Find")
	defer span.Finish()

	var offices []*models.Office
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
		Find(&offices)
	if row.Error != nil {
		return nil, 0, row.Error
	}

	return offices, uint32(count), nil
}

func (ur *officeRepo) FindByID(ctx context.Context, ID uint32) (*models.Office, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OfficeRepo.FindByID")
	defer span.Finish()

	office := new(models.Office)
	err := ur.db.WithContext(ctx).First(&office, "id = ?", ID).Error
	if err != nil {
		return nil, err
	}

	return office, nil
}
