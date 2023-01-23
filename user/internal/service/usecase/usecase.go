package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"
	"user/internal/domain"
	"user/internal/models"
	"user/pkg/logger"
	"user/pkg/utils"

	"github.com/opentracing/opentracing-go"
)

// Service UseCase
type serviceUseCase struct {
	logger           logger.Logger
	serviceRepo      domain.ServiceRepository
	serviceRedisRepo domain.ServiceRedisRepository
}

// New Service UseCase
func NewServiceUseCase(logger logger.Logger, serviceRepo domain.ServiceRepository, serviceRedisRepo domain.ServiceRedisRepository) domain.ServiceUseCase {
	return &serviceUseCase{logger: logger, serviceRepo: serviceRepo, serviceRedisRepo: serviceRedisRepo}
}

// *Command

// create new service
func (u *serviceUseCase) CreateService(ctx context.Context, service *models.Service) (*models.Service, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ServiceUseCase.CreateService")
	defer span.Finish()

	newService, err := u.serviceRepo.CreateService(ctx, service)
	if err != nil {
		u.logger.Errorf("serviceRepo.CreateService: %v", err)
		return nil, fmt.Errorf("serviceRepo.CreateService: %v", err)
	}

	return newService, nil
}

func (u *serviceUseCase) DeleteByID(ctx context.Context, ID uint32) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ServiceUseCase.DeleteByID")
	defer span.Finish()

	err := u.serviceRepo.DeleteByID(ctx, ID)
	if err != nil {
		u.logger.Errorf("serviceRepo.DeleteByID: %v", err)
		return fmt.Errorf("serviceRepo.DeleteByID: %v", err)
	}

	return nil
}

// *Query

func (u *serviceUseCase) Find(ctx context.Context, filters map[string]string, paginateQuery *utils.Pagination, expire time.Duration) ([]*models.Service, uint32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ServiceUseCase.Find")
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

	var cachedServices models.ServicesPaginate
	cachedByte, er := u.serviceRedisRepo.FindByID(ctx, "service_list:", filterKey)
	if er != nil {
		foundServiceList, totalCount, err := u.serviceRepo.Find(ctx, parsedFilters, paginateQuery)
		if err != nil {
			u.logger.Errorf("serviceRepo.Find: %v", err)
			return nil, 0, fmt.Errorf("serviceRepo.Find: %v", err)
		}

		foundServices := models.ServicesPaginate{
			ServiceList: foundServiceList,
			TotalCount:  totalCount,
		}

		foundServiceByte, err := json.Marshal(foundServices)
		if err != nil {
			u.logger.Errorf("json.Marshal: %v", err)
			return nil, 0, err
		}

		err = u.serviceRedisRepo.CreateService(ctx, "service_list:", filterKey, foundServiceByte, expire)
		if err != nil {
			u.logger.Errorf("serviceRedisRepo.CreateService", err)
		}

		return foundServiceList, totalCount, nil
	}

	err := json.Unmarshal(cachedByte, &cachedServices)
	if err != nil {
		u.logger.Errorf("json.Unmarshal: %v", err)
		return nil, 0, err
	}

	return cachedServices.ServiceList, cachedServices.TotalCount, nil
}
