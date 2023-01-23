package main

import (
	"context"
	"log"
	"os"
	"user/config"
	"user/internal/models"
	casbinPkg "user/pkg/casbin"
	"user/pkg/db"
	"user/pkg/logger"
	"user/pkg/utils"
)

func main() {
	configPath := utils.GetConfigPath(os.Getenv("CONFIG_PATH"))
	cfg, err := config.GetConfig(configPath)
	if err != nil {
		log.Fatalf("Loading config: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := logger.NewAPILogger(cfg)
	logger.InitLogger()
	logger.Infof(
		"AppVersion: %s, LogLevel: %s, Mode: %s",
		cfg.Server.AppVersion,
		cfg.Logger.Level,
		cfg.Server.Mode,
	)
	logger.Infof("Success parsed config: %#v", cfg.Server.AppVersion)

	db, err := db.NewDBInit(cfg)
	if err != nil {
		logger.Fatalf("gorm init: %s", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Errorf("gorm.DB.DB(): %v", err)
	}
	defer sqlDB.Close()

	err = db.Debug().AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Service{},
		&models.Method{},
	)
	if err != nil {
		logger.Fatalf("gorm auto migrate: %s", err)
	}
	logger.Info("db migrated")

	casbin, err := casbinPkg.NewCasbinClient(ctx, cfg)
	if err != nil {
		logger.Fatalf("cannot connect to casbin client: %v", err)
	}
	logger.Info("Casbin client connected")

	UserServiceMigration(ctx, logger, db, casbin)
	SafetyServiceMigration(ctx, logger, db, casbin)
}
