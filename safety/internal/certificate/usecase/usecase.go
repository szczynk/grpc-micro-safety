package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"safety/internal/domain"
	"safety/internal/models"
	"safety/pkg/logger"
	"safety/pkg/utils"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
)

// Certificate UseCase
type certificateUseCase struct {
	logger               logger.Logger
	certificateRepo      domain.CertificateRepository
	certificateRedisRepo domain.CertificateRedisRepository
}

// New Certificate UseCase
func NewCertificateUseCase(logger logger.Logger, certificateRepo domain.CertificateRepository, certificateRedisRepo domain.CertificateRedisRepository) domain.CertificateUseCase {
	return &certificateUseCase{logger: logger, certificateRepo: certificateRepo, certificateRedisRepo: certificateRedisRepo}
}

// *Command

// create new certificate
func (u *certificateUseCase) CreateCertificate(ctx context.Context, certificate *models.Certificate) (*models.Certificate, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CertificateUseCase.CreateCertificate")
	defer span.Finish()

	newWorkspace, err := u.certificateRepo.CreateCertificate(ctx, certificate)
	if err != nil {
		u.logger.Errorf("certificateRepo.CreateCertificate: %v", err)
		return nil, fmt.Errorf("certificateRepo.CreateCertificate: %v", err)
	}

	return newWorkspace, nil
}

func (u *certificateUseCase) UpdateByID(ctx context.Context, ID uint32, updates models.Certificate) (*models.Certificate, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CertificateUseCase.UpdateByID")
	defer span.Finish()

	err := u.certificateRedisRepo.DeleteByID(ctx, "certificate_id:", strconv.FormatUint(uint64(ID), 10))
	if err != nil {
		u.logger.Errorf("certificateRedisRepo.DeleteByID: %v", err)
	}

	updatedCertificate, err := u.certificateRepo.UpdateByID(ctx, ID, updates)
	if err != nil {
		u.logger.Errorf("certificateRepo.UpdateByID: %v", err)
		return nil, fmt.Errorf("certificateRepo.UpdateByID: %v", err)
	}

	return updatedCertificate, err
}

func (u *certificateUseCase) DeleteByID(ctx context.Context, ID uint32) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CertificateUseCase.DeleteByID")
	defer span.Finish()

	err := u.certificateRedisRepo.DeleteByID(ctx, "certificate_id:", strconv.FormatUint(uint64(ID), 10))
	if err != nil {
		u.logger.Errorf("certificateRedisRepo.DeleteByID: %v", err)
	}

	err = u.certificateRepo.DeleteByID(ctx, ID)
	if err != nil {
		u.logger.Errorf("certificateRepo.DeleteByID: %v", err)
		return fmt.Errorf("certificateRepo.DeleteByID: %v", err)
	}

	return nil
}

// *Query

func (u *certificateUseCase) Find(ctx context.Context, filters map[string]string, paginateQuery *utils.Pagination, expire time.Duration) ([]*models.Certificate, uint32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CertificateUseCase.Find")
	defer span.Finish()

	keys := make([]string, 0, len(filters))
	for k := range filters {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var filterKey string
	parsedFilters := make(map[string]interface{}, len(filters))
	for _, k := range keys {
		if len(filters[k]) > 0 && filters[k] != "0" {
			filterKey += fmt.Sprintf("%v_%v-", k, filters[k])

			if k != "limit" && k != "page" && k != "sort" {
				parsedFilters[k] = filters[k]
			}
		}
	}
	filterKey = strings.TrimSuffix(filterKey, "-")

	var cachedCertificates models.CertificatesPaginate
	cachedByte, er := u.certificateRedisRepo.FindByID(ctx, "certificate_list:", filterKey)
	if er != nil {
		foundCertificateList, totalCount, err := u.certificateRepo.Find(ctx, parsedFilters, paginateQuery)
		if err != nil {
			u.logger.Errorf("certificateRepo.Find: %v", err)
			return nil, 0, fmt.Errorf("certificateRepo.Find: %v", err)
		}

		foundCertificates := models.CertificatesPaginate{
			CertificateList: foundCertificateList,
			TotalCount:      totalCount,
		}

		foundCertificateByte, err := json.Marshal(foundCertificates)
		if err != nil {
			u.logger.Errorf("json.Marshal: %v", err)
			return nil, 0, err
		}

		err = u.certificateRedisRepo.CreateCertificate(ctx, "certificate_list:", filterKey, foundCertificateByte, expire)
		if err != nil {
			u.logger.Errorf("certificateRedisRepo.CreateCertificate", err)
		}

		return foundCertificateList, totalCount, nil
	}

	err := json.Unmarshal(cachedByte, &cachedCertificates)
	if err != nil {
		u.logger.Errorf("json.Unmarshal: %v", err)
		return nil, 0, err
	}

	return cachedCertificates.CertificateList, cachedCertificates.TotalCount, nil
}

func (u *certificateUseCase) FindByID(ctx context.Context, ID uint32, expire time.Duration) (*models.Certificate, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CertificateUseCase.FindByID")
	defer span.Finish()

	cachedCertificate := new(models.Certificate)
	cachedByte, er := u.certificateRedisRepo.FindByID(ctx, "certificate_id:", strconv.FormatUint(uint64(ID), 10))
	if er != nil {
		foundCertificate, err := u.certificateRepo.FindByID(ctx, ID)
		if err != nil {
			u.logger.Errorf("certificateRepo.FindByID: %v", err)
			return nil, fmt.Errorf("certificateRepo.FindByID: %v", err)
		}

		foundCertificateByte, err := json.Marshal(foundCertificate)
		if err != nil {
			u.logger.Errorf("json.Marshal: %v", err)
			return nil, err
		}

		err = u.certificateRedisRepo.CreateCertificate(ctx, "certificate_id:", strconv.FormatUint(uint64(ID), 10), foundCertificateByte, expire)
		if err != nil {
			u.logger.Errorf("certificateRedisRepo.CreateCertificate: %v", err)
		}

		return foundCertificate, nil
	}

	err := json.Unmarshal(cachedByte, &cachedCertificate)
	if err != nil {
		u.logger.Errorf("json.Unmarshal: %v", err)
		return nil, err
	}

	return cachedCertificate, nil
}

func (u *certificateUseCase) CountByUserID(ctx context.Context, userId uuid.UUID, expire time.Duration) (uint32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CertificateUseCase.CountByUserID")
	defer span.Finish()

	var cachedCertificates models.CertificatesPaginate
	cachedByte, er := u.certificateRedisRepo.FindByID(ctx, "certificate_user_id:", userId.String())
	if er != nil {
		totalCount, err := u.certificateRepo.CountByUserID(ctx, userId)
		if err != nil {
			u.logger.Errorf("certificateRepo.Find: %v", err)
			return 0, fmt.Errorf("certificateRepo.Find: %v", err)
		}

		foundCertificates := models.CertificatesPaginate{
			TotalCount: totalCount,
		}

		foundCertificateByte, err := json.Marshal(foundCertificates)
		if err != nil {
			u.logger.Errorf("json.Marshal: %v", err)
			return 0, err
		}

		err = u.certificateRedisRepo.CreateCertificate(ctx, "certificate_user_id:", userId.String(), foundCertificateByte, expire)
		if err != nil {
			u.logger.Errorf("certificateRedisRepo.CreateCertificate", err)
		}

		return totalCount, nil
	}

	err := json.Unmarshal(cachedByte, &cachedCertificates)
	if err != nil {
		u.logger.Errorf("json.Unmarshal: %v", err)
		return 0, err
	}

	return cachedCertificates.TotalCount, nil
}
