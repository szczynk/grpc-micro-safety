package main

import (
	"context"
	"user/internal/models"
	"user/pkg/logger"

	"github.com/casbin/casbin-go-client/client"
	"gorm.io/gorm"
)

func SafetyServiceMigration(ctx context.Context, logger logger.Logger, db *gorm.DB, casbin *client.Enforcer) {
	services := []models.Service{
		{Service: "pb.AttendanceService"},
		{Service: "pb.CertificateService"},
		{Service: "pb.CheckService"},
		{Service: "pb.OfficeService"},
		{Service: "pb.ScheduleService"},
		{Service: "pb.WorkspaceService"},

		{Service: "pb.ImageService"},
	}
	s := db.WithContext(ctx).Create(&services)
	if s.Error != nil {
		logger.Warnf("Default service already existed err: %v", s.Error)
	}
	logger.Info("Default service defined")

	methods := []models.Method{
		{Method: "CreateAttendance"},
		{Method: "UpdateAttendanceById"},
		{Method: "DeleteAttendanceById"},

		{Method: "FindAttendances"},
		{Method: "FindAttendanceById"},

		{Method: "CreateCertificate"},
		{Method: "UpdateCertificateById"},
		{Method: "DeleteCertificateById"},

		{Method: "FindCertificates"},
		{Method: "FindCertificateById"},

		{Method: "CheckIn"},
		{Method: "CheckOut"},

		{Method: "FindChecks"},
		{Method: "FindCheckById"},

		{Method: "CreateOffice"},
		{Method: "UpdateOfficeById"},
		{Method: "DeleteOfficeById"},

		{Method: "FindOffices"},
		{Method: "FindOfficeById"},

		{Method: "CreateSchedule"},
		{Method: "UpdateScheduleById"},
		{Method: "DeleteScheduleById"},

		{Method: "FindSchedules"},
		{Method: "FindScheduleById"},

		{Method: "CreateWorkspace"},
		{Method: "DeleteWorkspaceById"},

		{Method: "FindWorkspaces"},

		{Method: "UploadImage"},
		{Method: "GetImageId"},
	}
	m := db.WithContext(ctx).Create(&methods)
	if m.Error != nil {
		logger.Warnf("Default method already existed err: %v", m.Error)
	}
	logger.Info("Default method defined")

	adminPolicies := [][]string{
		{"admin", "pb.AttendanceService", "CreateAttendance"},
		{"admin", "pb.AttendanceService", "UpdateAttendanceById"},
		{"admin", "pb.AttendanceService", "DeleteAttendanceById"},

		{"admin", "pb.AttendanceService", "FindAttendances"},
		{"admin", "pb.AttendanceService", "FindAttendanceById"},

		{"admin", "pb.CertificateService", "CreateCertificate"},
		{"admin", "pb.CertificateService", "UpdateCertificateById"},
		{"admin", "pb.CertificateService", "DeleteCertificateById"},

		{"admin", "pb.CertificateService", "FindCertificates"},
		{"admin", "pb.CertificateService", "FindCertificateById"},

		{"admin", "pb.CheckService", "CheckIn"},
		{"admin", "pb.CheckService", "CheckOut"},

		{"admin", "pb.CheckService", "FindChecks"},
		{"admin", "pb.CheckService", "FindCheckById"},

		{"admin", "pb.OfficeService", "CreateOffice"},
		{"admin", "pb.OfficeService", "UpdateOfficeById"},
		{"admin", "pb.OfficeService", "DeleteOfficeById"},

		{"admin", "pb.OfficeService", "FindOffices"},
		{"admin", "pb.OfficeService", "FindOfficeById"},

		{"admin", "pb.ScheduleService", "CreateSchedule"},
		{"admin", "pb.ScheduleService", "UpdateScheduleById"},
		{"admin", "pb.ScheduleService", "DeleteScheduleById"},

		{"admin", "pb.ScheduleService", "FindSchedules"},
		{"admin", "pb.ScheduleService", "FindScheduleById"},

		{"admin", "pb.WorkspaceService", "CreateWorkspace"},
		{"admin", "pb.WorkspaceService", "DeleteWorkspaceById"},

		{"admin", "pb.WorkspaceService", "FindWorkspaces"},

		{"admin", "pb.ImageService", "UploadImage"},
		{"admin", "pb.ImageService", "GetImageId"},
	}

	userPolicies := [][]string{
		{"user", "pb.AttendanceService", "CreateAttendance"},
		// {"user", "pb.AttendanceService", "UpdateAttendanceById"},
		// {"user", "pb.AttendanceService", "DeleteAttendanceById"},

		{"user", "pb.AttendanceService", "FindAttendances"},
		{"user", "pb.AttendanceService", "FindAttendanceById"},

		{"user", "pb.CertificateService", "CreateCertificate"},
		{"user", "pb.CertificateService", "UpdateCertificateById"},
		// {"user", "pb.CertificateService", "DeleteCertificateById"},

		{"user", "pb.CertificateService", "FindCertificates"},
		{"user", "pb.CertificateService", "FindCertificateById"},

		{"user", "pb.CheckService", "CheckIn"},
		{"user", "pb.CheckService", "CheckOut"},

		{"user", "pb.CheckService", "FindChecks"},
		{"user", "pb.CheckService", "FindCheckById"},

		// {"user", "pb.OfficeService", "CreateOffice"},
		// {"user", "pb.OfficeService", "UpdateOfficeById"},
		// {"user", "pb.OfficeService", "DeleteOfficeById"},

		{"user", "pb.OfficeService", "FindOffices"},
		{"user", "pb.OfficeService", "FindOfficeById"},

		// {"user", "pb.ScheduleService", "CreateSchedule"},
		// {"user", "pb.ScheduleService", "UpdateScheduleById"},
		// {"user", "pb.ScheduleService", "DeleteScheduleById"},

		{"user", "pb.ScheduleService", "FindSchedules"},
		{"user", "pb.ScheduleService", "FindScheduleById"},

		// {"user", "pb.WorkspaceService", "CreateWorkspace"},
		// {"user", "pb.WorkspaceService", "DeleteWorkspaceById"},

		{"user", "pb.WorkspaceService", "FindWorkspaces"},

		{"user", "pb.ImageService", "UploadImage"},
		{"user", "pb.ImageService", "GetImageId"},
	}

	policies := append(adminPolicies, userPolicies...)

	for _, policy := range policies {
		ok, err := casbin.HasPolicy(ctx, policy)
		if err != nil {
			logger.Fatalf("casbin has policy err: %v", err)
		}
		if !ok {
			ok1, err := casbin.AddPolicy(ctx, policy)
			if err != nil {
				logger.Fatalf("casbin add policy err: %v", err)
			}
			if !ok1 {
				logger.Warnf("this %v policy already existed", policy)
			}
		} else {
			logger.Warnf("this %v policy already existed", policy)
		}
	}
	logger.Info("Casbin Policy defined")

	logger.Info("Safety Service Migration Done")
}
