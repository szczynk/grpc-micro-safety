package main

import (
	"log"
	"os"
	"safety/config"
	"safety/internal/models"
	"safety/pkg/db"
	"safety/pkg/logger"
	"safety/pkg/utils"
)

// TODO(Bagus): migration
func main() {
	configPath := utils.GetConfigPath(os.Getenv("CONFIG_PATH"))
	cfg, err := config.GetConfig(configPath)
	if err != nil {
		log.Fatalf("Loading config: %v", err)
	}

	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

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
		&models.Office{},
		&models.Schedule{},
		&models.Attendance{},
		&models.Certificate{},
	)
	if err != nil {
		logger.Fatalf("gorm auto migrate: %s", err)
	}
	logger.Info("db migrated")
}
