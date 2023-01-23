package main

import (
	"context"
	"fmt"
	"log"
	"mail/config"
	"mail/internal/server"
	"mail/pkg/db"
	"mail/pkg/errors"
	jaegerPkg "mail/pkg/jaeger"
	kafkaPkg "mail/pkg/kafka"
	loggerPkg "mail/pkg/logger"
	"mail/pkg/mailer"
	redisPkg "mail/pkg/redis"
	"mail/pkg/utils"
	"os"

	"github.com/opentracing/opentracing-go"
	"golang.org/x/sync/errgroup"
)

func main() {
	configPath := utils.GetConfigPath(os.Getenv("CONFIG_PATH"))
	cfg, err := config.GetConfig(configPath)
	if err != nil {
		log.Fatalf("Loading config: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	logger := loggerPkg.NewAPILogger(cfg)
	logger.InitLogger()
	logger.Infof("Starting %s microservice", cfg.Server.Name)
	logger.Infof(
		"AppVersion: %s, LogLevel: %s, Mode: %s",
		cfg.Server.AppVersion,
		cfg.Logger.Level,
		cfg.Server.Mode,
	)
	logger.Infof("Success parsed config: %#v", cfg.Server.AppVersion)

	mailer, err := mailer.NewMailer(cfg)
	if err != nil {
		logger.Fatalf("gorm init: %s", err)
	}

	db, err := db.NewDBInit(cfg)
	if err != nil {
		logger.Fatalf("gorm init: %s", err)
	}
	defer func() {
		sqlDB, _ := db.DB()
		_ = sqlDB.Close()
	}()
	logger.Info("Gorm connected")

	redisClient := redisPkg.NewRedisClient(cfg)
	defer redisClient.Close()
	logger.Info("Redis connected")

	tracer, closer, err := jaegerPkg.InitJaeger(cfg)
	if err != nil {
		logger.Fatal("cannot create tracer", err)
	}
	logger.Info("Jaeger connected")

	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()
	logger.Info("Opentracing connected")

	// TODO(Bagus): until i learn to insert metrics, disable metrics
	// metrics, err := metric.CreateMetrics(cfg)
	// if err != nil {
	// 	logger.Errorf("CreateMetrics Error: %s", err)
	// }
	// logger.Info("Metrics connected")

	kafkaProducer := kafkaPkg.NewProducer(logger, cfg)
	defer kafkaProducer.Close()

	mailServer := server.NewMailServer(
		ctx, logger, cfg,
		mailer,
		db, redisClient,
		// metrics,
		kafkaProducer,
	)

	g.Go(func() error {
		return mailServer.RunMetrics(cancel)
	})

	g.Go(func() error {
		return mailServer.Run()
	})

	g.Go(func() error {
		if sig := errors.SignalHandler(ctx); sig != nil {
			cancel()
			logger.Info(fmt.Sprintf("Server shutdown by signal: %s", sig))
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		logger.Error(fmt.Sprintf("Server terminated: %s", err))
	}
}
