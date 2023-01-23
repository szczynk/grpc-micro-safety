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

	"github.com/opentracing/opentracing-go"
)

// Office UseCase
type officeUseCase struct {
	logger          logger.Logger
	officeRepo      domain.OfficeRepository
	officeRedisRepo domain.OfficeRedisRepository
}

// New Office UseCase
func NewOfficeUseCase(logger logger.Logger, officeRepo domain.OfficeRepository, officeRedisRepo domain.OfficeRedisRepository) domain.OfficeUseCase {
	return &officeUseCase{logger: logger, officeRepo: officeRepo, officeRedisRepo: officeRedisRepo}
}

// *Command

// create new office
func (u *officeUseCase) CreateOffice(ctx context.Context, office *models.Office) (*models.Office, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OfficeUseCase.CreateOffice")
	defer span.Finish()

	newOffice, err := u.officeRepo.CreateOffice(ctx, office)
	if err != nil {
		u.logger.Errorf("officeRepo.CreateOffice: %v", err)
		return nil, fmt.Errorf("officeRepo.CreateOffice: %v", err)
	}

	return newOffice, nil
}

func (u *officeUseCase) UpdateByID(ctx context.Context, ID uint32, updates models.Office) (*models.Office, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OfficeUseCase.UpdateByID")
	defer span.Finish()

	err := u.officeRedisRepo.DeleteByID(ctx, "office_id:", strconv.FormatUint(uint64(ID), 10))
	if err != nil {
		u.logger.Errorf("officeRedisRepo.DeleteByID: %v", err)
	}

	updatedOffice, err := u.officeRepo.UpdateByID(ctx, ID, updates)
	if err != nil {
		u.logger.Errorf("officeRepo.UpdateByID: %v", err)
		return nil, fmt.Errorf("officeRepo.UpdateByID: %v", err)
	}

	return updatedOffice, err
}

func (u *officeUseCase) DeleteByID(ctx context.Context, ID uint32) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OfficeUseCase.DeleteByID")
	defer span.Finish()

	err := u.officeRedisRepo.DeleteByID(ctx, "office_id:", strconv.FormatUint(uint64(ID), 10))
	if err != nil {
		u.logger.Errorf("officeRedisRepo.DeleteByID: %v", err)
	}

	err = u.officeRepo.DeleteByID(ctx, ID)
	if err != nil {
		u.logger.Errorf("officeRepo.DeleteByID: %v", err)
		return fmt.Errorf("officeRepo.DeleteByID: %v", err)
	}

	return nil
}

// *Query

func (u *officeUseCase) Find(ctx context.Context, filters map[string]string, paginateQuery *utils.Pagination, expire time.Duration) ([]*models.Office, uint32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OfficeUseCase.Find")
	defer span.Finish()

	keys := make([]string, 0, len(filters))
	for k := range filters {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var filterKey string
	parsedFilters := make(map[string]interface{}, len(filters))
	for _, k := range keys {
		if len(filters[k]) > 0 {
			filterKey += fmt.Sprintf("%v_%v-", k, filters[k])

			if k != "limit" && k != "page" && k != "sort" {
				parsedFilters[k] = filters[k]
			}
		}
	}
	filterKey = strings.TrimSuffix(filterKey, "-")

	var cachedOffices models.OfficesPaginate
	cachedByte, er := u.officeRedisRepo.FindByID(ctx, "office_list:", filterKey)
	if er != nil {
		foundOfficeList, totalCount, err := u.officeRepo.Find(ctx, parsedFilters, paginateQuery)
		if err != nil {
			u.logger.Errorf("officeRepo.Find: %v", err)
			return nil, 0, fmt.Errorf("officeRepo.Find: %v", err)
		}

		foundOffices := models.OfficesPaginate{
			OfficeList: foundOfficeList,
			TotalCount: totalCount,
		}

		foundOfficeByte, err := json.Marshal(foundOffices)
		if err != nil {
			u.logger.Errorf("json.Marshal: %v", err)
			return nil, 0, err
		}

		err = u.officeRedisRepo.CreateOffice(ctx, "office_list:", filterKey, foundOfficeByte, expire)
		if err != nil {
			u.logger.Errorf("officeRedisRepo.CreateOffice", err)
		}

		return foundOfficeList, totalCount, nil
	}

	err := json.Unmarshal(cachedByte, &cachedOffices)
	if err != nil {
		u.logger.Errorf("json.Unmarshal: %v", err)
		return nil, 0, err
	}

	return cachedOffices.OfficeList, cachedOffices.TotalCount, nil
}

func (u *officeUseCase) FindByID(ctx context.Context, ID uint32, expire time.Duration) (*models.Office, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OfficeUseCase.FindByID")
	defer span.Finish()

	cachedOffice := new(models.Office)
	cachedByte, er := u.officeRedisRepo.FindByID(ctx, "office_id:", strconv.FormatUint(uint64(ID), 10))
	if er != nil {
		foundOffice, err := u.officeRepo.FindByID(ctx, ID)
		if err != nil {
			u.logger.Errorf("officeRepo.FindByID: %v", err)
			return nil, fmt.Errorf("officeRepo.FindByID: %v", err)
		}

		foundOfficeByte, err := json.Marshal(foundOffice)
		if err != nil {
			u.logger.Errorf("json.Marshal: %v", err)
			return nil, err
		}

		err = u.officeRedisRepo.CreateOffice(ctx, "office_id:", strconv.FormatUint(uint64(ID), 10), foundOfficeByte, expire)
		if err != nil {
			u.logger.Errorf("officeRedisRepo.CreateOffice: %v", err)
		}

		return foundOffice, nil
	}

	err := json.Unmarshal(cachedByte, &cachedOffice)
	if err != nil {
		u.logger.Errorf("json.Unmarshal: %v", err)
		return nil, err
	}

	return cachedOffice, nil
}
