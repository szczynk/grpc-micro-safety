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

// Attendance repository
type attendanceRepo struct {
	db *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) domain.AttendanceRepository {
	return &attendanceRepo{db: db}
}

// *Command

func (ur *attendanceRepo) CreateAttendance(ctx context.Context, attendance *models.Attendance) (*models.Attendance, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "AttendanceRepo.CreateAttendance")
	defer span.Finish()

	err := ur.db.WithContext(ctx).Create(&attendance).Error
	if err != nil {
		return nil, err
	}

	return attendance, nil
}

func (ur *attendanceRepo) UpdateByID(ctx context.Context, ID uint32, updates models.Attendance) (*models.Attendance, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "AttendanceRepo.UpdateByID")
	defer span.Finish()

	attendance := new(models.Attendance)
	err := ur.db.WithContext(ctx).Model(&attendance).Where("id = ?", ID).Updates(updates).Error
	if err != nil {
		return nil, err
	}

	return attendance, nil
}

func (ur *attendanceRepo) DeleteByID(ctx context.Context, ID uint32) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "AttendanceRepo.DeleteByID")
	defer span.Finish()

	attendance := new(models.Attendance)
	err := ur.db.WithContext(ctx).Delete(&attendance, ID).Error
	if err != nil {
		return err
	}

	return nil
}

// *Query

func (ur *attendanceRepo) Find(ctx context.Context, filter map[string]interface{}, paginateQuery *utils.Pagination) ([]*models.Attendance, uint32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "AttendanceRepo.Find")
	defer span.Finish()

	var attendances []*models.Attendance
	var count int64

	result := ur.db.WithContext(ctx).Table("attendances").
		Select("attendances.*, schedules.office_id, schedules.date AS schedule_date, offices.name AS office_name, users.username AS user_username, users.avatar AS user_avatar").
		Joins("JOIN schedules ON attendances.schedule_id = schedules.id").
		Joins("JOIN offices ON schedules.office_id = offices.id").
		Joins("JOIN users ON attendances.user_id = users.id")
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
		Scan(&attendances)
	if row.Error != nil {
		return nil, 0, row.Error
	}

	return attendances, uint32(count), nil
}

func (ur *attendanceRepo) FindByID(ctx context.Context, ID uint32) (*models.Attendance, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "AttendanceRepo.FindByID")
	defer span.Finish()

	attendance := new(models.Attendance)
	err := ur.db.WithContext(ctx).Table("attendances").
		Select("attendances.*, schedules.office_id, schedules.date AS schedule_date, offices.name AS office_name, users.username AS user_username, users.avatar AS user_avatar").
		Joins("JOIN schedules ON attendances.schedule_id = schedules.id").
		Joins("JOIN offices ON schedules.office_id = offices.id").
		Joins("JOIN users ON attendances.user_id = users.id").
		First(&attendance, "id = ?", ID).Error
	if err != nil {
		return nil, err
	}

	return attendance, nil
}

func (ur *attendanceRepo) FindChecks(ctx context.Context, filter map[string]interface{}, paginateQuery *utils.Pagination) ([]*models.Attendance, uint32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "AttendanceRepo.Find")
	defer span.Finish()

	var attendances []*models.Attendance
	var count int64

	result := ur.db.WithContext(ctx).Table("attendances").
		Select("attendances.*, schedules.office_id, schedules.date AS schedule_date, offices.name AS office_name, users.username AS user_username, users.avatar AS user_avatar").
		Joins("JOIN schedules ON attendances.schedule_id = schedules.id").
		Joins("JOIN offices ON schedules.office_id = offices.id").
		Joins("JOIN users ON attendances.user_id = users.id").
		Where("attendances.check_in IS NOT null AND attendances.check_out IS NOT null")
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
		Scan(&attendances)
	if row.Error != nil {
		return nil, 0, row.Error
	}

	return attendances, uint32(count), nil
}

func (ur *attendanceRepo) FindCheckByID(ctx context.Context, ID uint32) (*models.Attendance, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "AttendanceRepo.FindByID")
	defer span.Finish()

	attendance := new(models.Attendance)
	err := ur.db.WithContext(ctx).Table("attendances").
		Select("attendances.*, schedules.office_id, schedules.date AS schedule_date, offices.name AS office_name, users.username AS user_username, users.avatar AS user_avatar").
		Joins("JOIN schedules ON attendances.schedule_id = schedules.id").
		Joins("JOIN offices ON schedules.office_id = offices.id").
		Joins("JOIN users ON attendances.user_id = users.id").
		Where("attendances.check_in IS NOT null AND attendances.check_out IS NOT null").
		First(&attendance, "id = ?", ID).Error
	if err != nil {
		return nil, err
	}

	return attendance, nil
}
