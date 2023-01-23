package main

import (
	"auth/config"
	"auth/internal/server"
	casbinPkg "auth/pkg/casbin"
	"auth/pkg/db"
	"auth/pkg/errors"
	"auth/pkg/healthz"
	jaegerPkg "auth/pkg/jaeger"
	kafkaPkg "auth/pkg/kafka"
	"auth/pkg/limiter"
	loggerPkg "auth/pkg/logger"
	redisPkg "auth/pkg/redis"
	"auth/pkg/token"
	"auth/pkg/utils"
	"context"
	"fmt"
	"log"
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

	tokenMaker, err := token.NewPasetoMaker(cfg.Server.JwtSecretKey)
	if err != nil {
		logger.Fatalf("cannot create token maker: %v", err)
	}
	logger.Info("Token maker connected")

	kafkaProducer := kafkaPkg.NewProducer(logger, cfg)
	defer kafkaProducer.Close()

	healthClient, err := healthz.NewHealthzClient(ctx, cfg)
	if err != nil {
		logger.Fatalf("cannot connect to healthz client: %v", err)
	}
	logger.Info("Health client gateway connected")

	casbinClient, err := casbinPkg.NewCasbinClient(ctx, cfg)
	if err != nil {
		logger.Fatalf("cannot connect to casbin client: %v", err)
	}
	logger.Info("Casbin client connected")

	limiter, err := limiter.NewRateLimit(cfg)
	if err != nil {
		logger.Fatalf("cannot connect to limiter: %v", err)
	}
	defer limiter.Store().Close()
	logger.Info("Limiter connected")

	authServer := server.NewAuthServer(
		ctx, logger, cfg, db,
		redisClient,
		// metrics,
		tokenMaker, kafkaProducer,
		healthClient, casbinClient,
		limiter,
	)

	g.Go(func() error {
		return authServer.RunMetrics(cancel)
	})

	g.Go(func() error {
		return authServer.RunGateway()
	})

	g.Go(func() error {
		return authServer.Run()
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
