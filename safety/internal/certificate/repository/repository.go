package repository

import (
	"context"
	"fmt"
	"safety/internal/domain"
	"safety/internal/models"
	"safety/pkg/utils"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Certificate repository
type certificateRepo struct {
	db *gorm.DB
}

func NewCertificateRepository(db *gorm.DB) domain.CertificateRepository {
	return &certificateRepo{db: db}
}

// *Command

func (ur *certificateRepo) CreateCertificate(ctx context.Context, certificate *models.Certificate) (*models.Certificate, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CertificateRepo.CreateCertificate")
	defer span.Finish()

	err := ur.db.WithContext(ctx).Create(&certificate).Error
	if err != nil {
		return nil, err
	}

	return certificate, nil
}

func (ur *certificateRepo) UpdateByID(ctx context.Context, ID uint32, updates models.Certificate) (*models.Certificate, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CertificateRepo.UpdateByID")
	defer span.Finish()

	certificate := new(models.Certificate)
	err := ur.db.WithContext(ctx).Model(&certificate).Clauses(clause.Returning{}).
		Where("id = ?", ID).Updates(updates).Error
	if err != nil {
		return nil, err
	}

	return certificate, nil
}

func (ur *certificateRepo) DeleteByID(ctx context.Context, ID uint32) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CertificateRepo.DeleteByID")
	defer span.Finish()

	certificate := new(models.Certificate)
	err := ur.db.WithContext(ctx).Delete(&certificate, ID).Error
	if err != nil {
		return err
	}

	return nil
}

// *Query

func (ur *certificateRepo) Find(ctx context.Context, filter map[string]interface{}, paginateQuery *utils.Pagination) ([]*models.Certificate, uint32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CertificateRepo.Find")
	defer span.Finish()

	var certificates []*models.Certificate
	var count int64

	result := ur.db.WithContext(ctx).Table("certificates").
		Select("certificates.*, users.username AS user_username, users.avatar AS user_avatar").
		Joins("JOIN users ON certificates.user_id = users.id")
	for k, v := range filter {
		switch k {
		case "admin_username":
			v1 := fmt.Sprintf("%%%v%%", v)
			result = result.Where(k+" LIKE ?", v1)
		default:
			result = result.Where(k+" = ?", v)
		}
	}

	result.Count(&count)

	row := result.Offset(int(paginateQuery.GetOffset())).
		Limit(int(paginateQuery.GetLimit())).
		Order(paginateQuery.GetSort()).
		Scan(&certificates)
	if row.Error != nil {
		return nil, 0, row.Error
	}

	return certificates, uint32(count), nil
}

func (ur *certificateRepo) FindByID(ctx context.Context, ID uint32) (*models.Certificate, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CertificateRepo.FindByID")
	defer span.Finish()

	certificate := new(models.Certificate)
	err := ur.db.WithContext(ctx).Table("certificates").
		First(&certificate, "id = ?", ID).Error
	if err != nil {
		return nil, err
	}

	return certificate, nil
}

func (ur *certificateRepo) CountByUserID(ctx context.Context, userId uuid.UUID) (uint32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CertificateRepo.CountByUserID")
	defer span.Finish()

	var count int64

	result := ur.db.WithContext(ctx).Table("certificates").
		Where("status = 'approved'").
		Where("user_id = ?", userId).
		Count(&count)

	if result.Error != nil {
		return 0, result.Error
	}

	return uint32(count), nil
}
