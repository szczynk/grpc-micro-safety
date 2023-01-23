package main

import (
	"context"
	"user/internal/models"
	"user/pkg/logger"
	"user/pkg/utils"

	"github.com/casbin/casbin-go-client/client"
	"gorm.io/gorm"
)

func UserServiceMigration(ctx context.Context, logger logger.Logger, db *gorm.DB, casbin *client.Enforcer) {
	hashedPassword, err := utils.HashPassword("admin123")
	if err != nil {
		logger.Errorf("utils.HashPassword: %v", err)
	}

	defaultAdmin := models.User{
		Username: "admin",
		Email:    "admin@winow.com",
		Password: hashedPassword,
		Role:     "admin",
		Avatar:   "https://avatars.dicebear.com/api/open-peeps/admin.svg",
		Verified: true,
	}

	u := db.WithContext(ctx).
		Where(models.User{Email: "admin@winow.com"}).
		FirstOrCreate(&defaultAdmin)
	if u.Error != nil {
		logger.Warnf("Default admin already existed err: %v", u.Error)
	}

	if u.RowsAffected == 1 {
		_, err = casbin.AddRoleForUser(ctx, defaultAdmin.ID.String(), defaultAdmin.Role)
		if err != nil {
			logger.Errorf("casbin.AddRoleForUser: %v", err)
		}
		logger.Info("Default admin created")
	} else {
		logger.Info("Default admin defined")
	}

	roles := []models.Role{
		{Role: "user"},
		{Role: "admin"},
	}
	r := db.WithContext(ctx).Create(&roles)
	if r.Error != nil {
		logger.Warnf("Default role already existed err: %v", r.Error)
	}
	logger.Info("Default role defined")

	services := []models.Service{
		{Service: "pb.UserService"},
		{Service: "pb.RoleService"},
		{Service: "pb.ServiceService"},
		{Service: "pb.MethodService"},
		{Service: "pb.PolicyService"},
	}
	s := db.WithContext(ctx).Create(&services)
	if s.Error != nil {
		logger.Warnf("Default service already existed err: %v", s.Error)
	}
	logger.Info("Default service defined")

	methods := []models.Method{
		{Method: "RefreshAccessToken"},
		{Method: "UpdateMe"},
		{Method: "ChangeEmail"},
		{Method: "Logout"},

		{Method: "GetMe"},

		{Method: "CreateUser"},
		{Method: "UpdateUserById"},
		{Method: "DeleteUserById"},

		{Method: "FindUsers"},
		{Method: "FindUserById"},

		{Method: "CreateRole"},
		{Method: "DeleteRoleById"},

		{Method: "FindRoles"},

		{Method: "CreateService"},
		{Method: "DeleteServiceById"},

		{Method: "FindServices"},

		{Method: "CreateMethod"},
		{Method: "DeleteMethodById"},

		{Method: "FindMethods"},

		{Method: "CreatePolicy"},
		{Method: "DeletePolicy"},

		{Method: "FindPolicies"},
	}
	m := db.WithContext(ctx).Create(&methods)
	if m.Error != nil {
		logger.Warnf("Default method already existed err: %v", m.Error)
	}
	logger.Info("Default method defined")

	adminPolicies := [][]string{
		{"admin", "pb.UserService", "RefreshAccessToken"},
		{"admin", "pb.UserService", "UpdateMe"},
		{"admin", "pb.UserService", "ChangeEmail"},
		{"admin", "pb.UserService", "Logout"},

		{"admin", "pb.UserService", "GetMe"},

		{"admin", "pb.UserService", "CreateUser"},
		{"admin", "pb.UserService", "UpdateUserById"},
		{"admin", "pb.UserService", "DeleteUserById"},

		{"admin", "pb.UserService", "FindUsers"},
		{"admin", "pb.UserService", "FindUserById"},

		{"admin", "pb.RoleService", "CreateRole"},
		{"admin", "pb.RoleService", "DeleteRoleById"},

		{"admin", "pb.RoleService", "FindRoles"},

		{"admin", "pb.ServiceService", "CreateService"},
		{"admin", "pb.ServiceService", "DeleteServiceById"},

		{"admin", "pb.ServiceService", "FindServices"},

		{"admin", "pb.MethodService", "CreateMethod"},
		{"admin", "pb.MethodService", "DeleteMethodById"},

		{"admin", "pb.MethodService", "FindMethods"},

		{"admin", "pb.PolicyService", "CreatePolicy"},
		{"admin", "pb.PolicyService", "DeletePolicy"},

		{"admin", "pb.PolicyService", "FindPolicies"},
	}

	userPolicies := [][]string{
		{"user", "pb.UserService", "RefreshAccessToken"},
		{"user", "pb.UserService", "UpdateMe"},
		{"user", "pb.UserService", "ChangeEmail"},
		{"user", "pb.UserService", "Logout"},

		{"user", "pb.UserService", "GetMe"},
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

	logger.Info("User Service Migration Done")
}
