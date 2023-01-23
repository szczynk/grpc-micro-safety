package repository

import (
	"context"
	"safety/internal/domain"
	"safety/internal/models"
	"safety/pkg/utils"
	"time"

	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

// Schedule repository
type scheduleRepo struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) domain.ScheduleRepository {
	return &scheduleRepo{db: db}
}

// *Command

func (ur *scheduleRepo) CreateSchedule(ctx context.Context, schedule *models.CreateSchedule) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ScheduleRepo.CreateSchedule")
	defer span.Finish()

	officeId, totalCapacity := schedule.OfficeID, schedule.TotalCapacity
	month := time.Month(schedule.Month)
	year := int(schedule.Year)

	day := 1
	gmt, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return err
	}

	start := time.Date(year, month, day, 0, 0, 0, 0, gmt)
	end := start.AddDate(0, 1, 0)
	diffInDays := int(end.Sub(start).Hours() / 24)

	scheduleList := make([]models.Schedule, 0, diffInDays)

	for start != end {
		scheduleList = append(
			scheduleList,
			models.Schedule{
				TotalCapacity: totalCapacity,
				OfficeID:      officeId,
				Date:          start,
			},
		)
		start = start.AddDate(0, 0, 1)
	}

	err = ur.db.WithContext(ctx).Create(&scheduleList).Error
	if err != nil {
		return err
	}

	return nil
}

func (ur *scheduleRepo) UpdateByID(ctx context.Context, ID uint32, updates models.Schedule) (*models.Schedule, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ScheduleRepo.UpdateByID")
	defer span.Finish()

	schedule := new(models.Schedule)
	err := ur.db.WithContext(ctx).Model(&schedule).Where("id = ?", ID).Updates(updates).Error
	if err != nil {
		return nil, err
	}

	return schedule, nil
}

func (ur *scheduleRepo) DeleteByID(ctx context.Context, ID uint32) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ScheduleRepo.DeleteByID")
	defer span.Finish()

	schedule := new(models.Schedule)
	err := ur.db.WithContext(ctx).Delete(&schedule, ID).Error
	if err != nil {
		return err
	}

	return nil
}

// *Query

func (ur *scheduleRepo) Find(ctx context.Context, filter map[string]interface{}, paginateQuery *utils.Pagination) ([]*models.Schedule, uint32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ScheduleRepo.Find")
	defer span.Finish()

	var schedules []*models.Schedule
	var count int64

	result := ur.db.WithContext(ctx).Table("schedules").
		Select("schedules.id,schedules.created_at, schedules.updated_at, schedules.deleted_at, schedules.office_id, schedules.total_capacity, schedules.capacity, schedules.date, offices.name AS office_name").
		Joins("JOIN offices ON schedules.office_id = offices.id")
	for k, v := range filter {
		switch k {
		case "year":
			result = result.Where("EXTRACT("+k+" FROM date) = ?", v)
		case "month":
			result = result.Where("EXTRACT("+k+" FROM date) = ?", v)
		default:
			result = result.Where(k+" = ?", v)
		}
	}

	result.Count(&count)

	row := result.Offset(int(paginateQuery.GetOffset())).
		Limit(int(paginateQuery.GetLimit())).
		Order(paginateQuery.GetSort()).
		Scan(&schedules)
	if row.Error != nil {
		return nil, 0, row.Error
	}

	return schedules, uint32(count), nil
}

func (ur *scheduleRepo) FindByID(ctx context.Context, ID uint32) (*models.Schedule, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ScheduleRepo.FindByID")
	defer span.Finish()

	schedule := new(models.Schedule)
	err := ur.db.WithContext(ctx).Table("schedules").
		Select("schedules.id, schedules.created_at, schedules.updated_at, schedules.deleted_at, schedules.office_id, schedules.total_capacity, schedules.capacity, schedules.date, offices.name AS office_name").
		Joins("JOIN offices ON schedules.office_id = offices.id").
		First(&schedule, "schedules.id = ?", ID).Error
	if err != nil {
		return nil, err
	}

	return schedule, nil
}
