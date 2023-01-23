package domain

import (
	"context"
	"safety/internal/models"
	"safety/pkg/utils"
	"time"

	"github.com/google/uuid"
)

// Certificate Repository
type CertificateRepository interface {
	CreateCertificate(ctx context.Context, certificate *models.Certificate) (*models.Certificate, error)
	UpdateByID(ctx context.Context, ID uint32, updates models.Certificate) (*models.Certificate, error)
	DeleteByID(ctx context.Context, ID uint32) error

	Find(ctx context.Context, filters map[string]interface{}, paginateQuery *utils.Pagination) ([]*models.Certificate, uint32, error)
	FindByID(ctx context.Context, ID uint32) (*models.Certificate, error)
	CountByUserID(ctx context.Context, userId uuid.UUID) (uint32, error)
}

// Certificate Redis Repository
type CertificateRedisRepository interface {
	CreateCertificate(ctx context.Context, key string, value string, dataBytes []byte, expire time.Duration) error
	DeleteByID(ctx context.Context, key string, value string) error

	FindByID(ctx context.Context, key string, value string) ([]byte, error)
}

// Certificate Usecase
type CertificateUseCase interface {
	CreateCertificate(ctx context.Context, certificate *models.Certificate) (*models.Certificate, error)
	UpdateByID(ctx context.Context, ID uint32, updates models.Certificate) (*models.Certificate, error)
	DeleteByID(ctx context.Context, ID uint32) error

	Find(ctx context.Context, filters map[string]string, paginateQuery *utils.Pagination, expire time.Duration) ([]*models.Certificate, uint32, error)
	FindByID(ctx context.Context, ID uint32, expire time.Duration) (*models.Certificate, error)
	CountByUserID(ctx context.Context, userId uuid.UUID, expire time.Duration) (uint32, error)
}
