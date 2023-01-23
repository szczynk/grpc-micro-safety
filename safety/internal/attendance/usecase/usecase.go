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

// Attendance UseCase
type attendanceUseCase struct {
	logger              logger.Logger
	attendanceRepo      domain.AttendanceRepository
	attendanceRedisRepo domain.AttendanceRedisRepository
}

// New Attendance UseCase
func NewAttendanceUseCase(logger logger.Logger, attendanceRepo domain.AttendanceRepository, attendanceRedisRepo domain.AttendanceRedisRepository) domain.AttendanceUseCase {
	return &attendanceUseCase{logger: logger, attendanceRepo: attendanceRepo, attendanceRedisRepo: attendanceRedisRepo}
}

// *Command

// create new attendance
func (u *attendanceUseCase) CreateAttendance(ctx context.Context, attendance *models.Attendance) (*models.Attendance, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "AttendanceUseCase.CreateAttendance")
	defer span.Finish()

	newWorkspace, err := u.attendanceRepo.CreateAttendance(ctx, attendance)
	if err != nil {
		u.logger.Errorf("attendanceRepo.CreateAttendance: %v", err)
		return nil, fmt.Errorf("attendanceRepo.CreateAttendance: %v", err)
	}

	return newWorkspace, nil
}

func (u *attendanceUseCase) UpdateByID(ctx context.Context, ID uint32, updates models.Attendance) (*models.Attendance, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "AttendanceUseCase.UpdateByID")
	defer span.Finish()

	err := u.attendanceRedisRepo.DeleteByID(ctx, "attendance_id:", strconv.FormatUint(uint64(ID), 10))
	if err != nil {
		u.logger.Errorf("attendanceRedisRepo.DeleteByID: %v", err)
	}

	updatedAttendance, err := u.attendanceRepo.UpdateByID(ctx, ID, updates)
	if err != nil {
		u.logger.Errorf("attendanceRepo.UpdateByID: %v", err)
		return nil, fmt.Errorf("attendanceRepo.UpdateByID: %v", err)
	}

	return updatedAttendance, err
}

func (u *attendanceUseCase) DeleteByID(ctx context.Context, ID uint32) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "AttendanceUseCase.DeleteByID")
	defer span.Finish()

	err := u.attendanceRedisRepo.DeleteByID(ctx, "attendance_id:", strconv.FormatUint(uint64(ID), 10))
	if err != nil {
		u.logger.Errorf("attendanceRedisRepo.DeleteByID: %v", err)
	}

	err = u.attendanceRepo.DeleteByID(ctx, ID)
	if err != nil {
		u.logger.Errorf("attendanceRepo.DeleteByID: %v", err)
		return fmt.Errorf("attendanceRepo.DeleteByID: %v", err)
	}

	return nil
}

// *Query

func (u *attendanceUseCase) Find(ctx context.Context, filters map[string]string, paginateQuery *utils.Pagination, expire time.Duration) ([]*models.Attendance, uint32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "AttendanceUseCase.Find")
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

	var cachedAttendances models.AttendancesPaginate
	cachedByte, er := u.attendanceRedisRepo.FindByID(ctx, "attendance_list:", filterKey)
	if er != nil {
		foundAttendanceList, totalCount, err := u.attendanceRepo.Find(ctx, parsedFilters, paginateQuery)
		if err != nil {
			u.logger.Errorf("attendanceRepo.Find: %v", err)
			return nil, 0, fmt.Errorf("attendanceRepo.Find: %v", err)
		}

		foundAttendances := models.AttendancesPaginate{
			AttendanceList: foundAttendanceList,
			TotalCount:     totalCount,
		}

		foundAttendanceByte, err := json.Marshal(foundAttendances)
		if err != nil {
			u.logger.Errorf("json.Marshal: %v", err)
			return nil, 0, err
		}

		err = u.attendanceRedisRepo.CreateAttendance(ctx, "attendance_list:", filterKey, foundAttendanceByte, expire)
		if err != nil {
			u.logger.Errorf("attendanceRedisRepo.CreateAttendance", err)
		}

		return foundAttendanceList, totalCount, nil
	}

	err := json.Unmarshal(cachedByte, &cachedAttendances)
	if err != nil {
		u.logger.Errorf("json.Unmarshal: %v", err)
		return nil, 0, err
	}

	return cachedAttendances.AttendanceList, cachedAttendances.TotalCount, nil
}

func (u *attendanceUseCase) FindByID(ctx context.Context, ID uint32, expire time.Duration) (*models.Attendance, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "AttendanceUseCase.FindByID")
	defer span.Finish()

	cachedAttendance := new(models.Attendance)
	cachedByte, er := u.attendanceRedisRepo.FindByID(ctx, "attendance_id:", strconv.FormatUint(uint64(ID), 10))
	if er != nil {
		foundAttendance, err := u.attendanceRepo.FindByID(ctx, ID)
		if err != nil {
			u.logger.Errorf("attendanceRepo.FindByID: %v", err)
			return nil, fmt.Errorf("attendanceRepo.FindByID: %v", err)
		}

		foundAttendanceByte, err := json.Marshal(foundAttendance)
		if err != nil {
			u.logger.Errorf("json.Marshal: %v", err)
			return nil, err
		}

		err = u.attendanceRedisRepo.CreateAttendance(ctx, "attendance_id:", strconv.FormatUint(uint64(ID), 10), foundAttendanceByte, expire)
		if err != nil {
			u.logger.Errorf("attendanceRedisRepo.CreateAttendance: %v", err)
		}

		return foundAttendance, nil
	}

	err := json.Unmarshal(cachedByte, &cachedAttendance)
	if err != nil {
		u.logger.Errorf("json.Unmarshal: %v", err)
		return nil, err
	}

	return cachedAttendance, nil
}
