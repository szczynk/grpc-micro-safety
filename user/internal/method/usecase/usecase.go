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

// Method UseCase
type methodUseCase struct {
	logger          logger.Logger
	methodRepo      domain.MethodRepository
	methodRedisRepo domain.MethodRedisRepository
}

// New Method UseCase
func NewMethodUseCase(logger logger.Logger, methodRepo domain.MethodRepository, methodRedisRepo domain.MethodRedisRepository) domain.MethodUseCase {
	return &methodUseCase{logger: logger, methodRepo: methodRepo, methodRedisRepo: methodRedisRepo}
}

// *Command

// create new method
func (u *methodUseCase) CreateMethod(ctx context.Context, method *models.Method) (*models.Method, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "MethodUseCase.CreateMethod")
	defer span.Finish()

	newMethod, err := u.methodRepo.CreateMethod(ctx, method)
	if err != nil {
		u.logger.Errorf("methodRepo.CreateMethod: %v", err)
		return nil, fmt.Errorf("methodRepo.CreateMethod: %v", err)
	}

	return newMethod, nil
}

func (u *methodUseCase) DeleteByID(ctx context.Context, ID uint32) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "MethodUseCase.DeleteByID")
	defer span.Finish()

	err := u.methodRepo.DeleteByID(ctx, ID)
	if err != nil {
		u.logger.Errorf("methodRepo.DeleteByID: %v", err)
		return fmt.Errorf("methodRepo.DeleteByID: %v", err)
	}

	return nil
}

// *Query

func (u *methodUseCase) Find(ctx context.Context, filters map[string]string, paginateQuery *utils.Pagination, expire time.Duration) ([]*models.Method, uint32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "MethodUseCase.Find")
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

	var cachedMethods models.MethodsPaginate
	cachedByte, er := u.methodRedisRepo.FindByID(ctx, "method_list:", filterKey)
	if er != nil {
		foundMethodList, totalCount, err := u.methodRepo.Find(ctx, parsedFilters, paginateQuery)
		if err != nil {
			u.logger.Errorf("methodRepo.Find: %v", err)
			return nil, 0, fmt.Errorf("methodRepo.Find: %v", err)
		}

		foundMethods := models.MethodsPaginate{
			MethodList: foundMethodList,
			TotalCount: totalCount,
		}

		foundMethodByte, err := json.Marshal(foundMethods)
		if err != nil {
			u.logger.Errorf("json.Marshal: %v", err)
			return nil, 0, err
		}

		err = u.methodRedisRepo.CreateMethod(ctx, "method_list:", filterKey, foundMethodByte, expire)
		if err != nil {
			u.logger.Errorf("methodRedisRepo.CreateMethod", err)
		}

		return foundMethodList, totalCount, nil
	}

	err := json.Unmarshal(cachedByte, &cachedMethods)
	if err != nil {
		u.logger.Errorf("json.Unmarshal: %v", err)
		return nil, 0, err
	}

	return cachedMethods.MethodList, cachedMethods.TotalCount, nil
}
