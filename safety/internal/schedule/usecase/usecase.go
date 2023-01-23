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

// Schedule UseCase
type scheduleUseCase struct {
	logger            logger.Logger
	scheduleRepo      domain.ScheduleRepository
	scheduleRedisRepo domain.ScheduleRedisRepository
}

// New Schedule UseCase
func NewScheduleUseCase(logger logger.Logger, scheduleRepo domain.ScheduleRepository, scheduleRedisRepo domain.ScheduleRedisRepository) domain.ScheduleUseCase {
	return &scheduleUseCase{logger: logger, scheduleRepo: scheduleRepo, scheduleRedisRepo: scheduleRedisRepo}
}

// *Command

// create new schedule
func (u *scheduleUseCase) CreateSchedule(ctx context.Context, schedule *models.CreateSchedule) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ScheduleUseCase.CreateSchedule")
	defer span.Finish()

	err := u.scheduleRepo.CreateSchedule(ctx, schedule)
	if err != nil {
		u.logger.Errorf("scheduleRepo.CreateSchedule: %v", err)
		return fmt.Errorf("scheduleRepo.CreateSchedule: %v", err)
	}

	return nil
}

func (u *scheduleUseCase) UpdateByID(ctx context.Context, ID uint32, updates models.Schedule) (*models.Schedule, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ScheduleUseCase.UpdateByID")
	defer span.Finish()

	err := u.scheduleRedisRepo.DeleteByID(ctx, "schedule_id:", strconv.FormatUint(uint64(ID), 10))
	if err != nil {
		u.logger.Errorf("scheduleRedisRepo.DeleteByID: %v", err)
	}

	updatedSchedule, err := u.scheduleRepo.UpdateByID(ctx, ID, updates)
	if err != nil {
		u.logger.Errorf("scheduleRepo.UpdateByID: %v", err)
		return nil, fmt.Errorf("scheduleRepo.UpdateByID: %v", err)
	}

	return updatedSchedule, err
}

func (u *scheduleUseCase) DeleteByID(ctx context.Context, ID uint32) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ScheduleUseCase.DeleteByID")
	defer span.Finish()

	err := u.scheduleRedisRepo.DeleteByID(ctx, "schedule_id:", strconv.FormatUint(uint64(ID), 10))
	if err != nil {
		u.logger.Errorf("scheduleRedisRepo.DeleteByID: %v", err)
	}

	err = u.scheduleRepo.DeleteByID(ctx, ID)
	if err != nil {
		u.logger.Errorf("scheduleRepo.DeleteByID: %v", err)
		return fmt.Errorf("scheduleRepo.DeleteByID: %v", err)
	}

	return nil
}

// *Query

func (u *scheduleUseCase) Find(ctx context.Context, filters map[string]string, paginateQuery *utils.Pagination, expire time.Duration) ([]*models.Schedule, uint32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ScheduleUseCase.Find")
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

	var cachedSchedules models.SchedulesPaginate
	cachedByte, er := u.scheduleRedisRepo.FindByID(ctx, "schedule_list:", filterKey)
	if er != nil {
		foundScheduleList, totalCount, err := u.scheduleRepo.Find(ctx, parsedFilters, paginateQuery)
		if err != nil {
			u.logger.Errorf("scheduleRepo.Find: %v", err)
			return nil, 0, fmt.Errorf("scheduleRepo.Find: %v", err)
		}

		foundSchedules := models.SchedulesPaginate{
			ScheduleList: foundScheduleList,
			TotalCount:   totalCount,
		}

		foundScheduleByte, err := json.Marshal(foundSchedules)
		if err != nil {
			u.logger.Errorf("json.Marshal: %v", err)
			return nil, 0, err
		}

		err = u.scheduleRedisRepo.CreateSchedule(ctx, "schedule_list:", filterKey, foundScheduleByte, expire)
		if err != nil {
			u.logger.Errorf("scheduleRedisRepo.CreateSchedule", err)
		}

		return foundScheduleList, totalCount, nil
	}

	err := json.Unmarshal(cachedByte, &cachedSchedules)
	if err != nil {
		u.logger.Errorf("json.Unmarshal: %v", err)
		return nil, 0, err
	}

	return cachedSchedules.ScheduleList, cachedSchedules.TotalCount, nil
}

func (u *scheduleUseCase) FindByID(ctx context.Context, ID uint32, expire time.Duration) (*models.Schedule, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ScheduleUseCase.FindByID")
	defer span.Finish()

	cachedSchedule := new(models.Schedule)
	cachedByte, er := u.scheduleRedisRepo.FindByID(ctx, "schedule_id:", strconv.FormatUint(uint64(ID), 10))
	if er != nil {
		foundSchedule, err := u.scheduleRepo.FindByID(ctx, ID)
		if err != nil {
			u.logger.Errorf("scheduleRepo.FindByID: %v", err)
			return nil, fmt.Errorf("scheduleRepo.FindByID: %v", err)
		}

		foundScheduleByte, err := json.Marshal(foundSchedule)
		if err != nil {
			u.logger.Errorf("json.Marshal: %v", err)
			return nil, err
		}

		err = u.scheduleRedisRepo.CreateSchedule(ctx, "schedule_id:", strconv.FormatUint(uint64(ID), 10), foundScheduleByte, expire)
		if err != nil {
			u.logger.Errorf("scheduleRedisRepo.CreateSchedule: %v", err)
		}

		return foundSchedule, nil
	}

	err := json.Unmarshal(cachedByte, &cachedSchedule)
	if err != nil {
		u.logger.Errorf("json.Unmarshal: %v", err)
		return nil, err
	}

	return cachedSchedule, nil
}
