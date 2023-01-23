package main

import (
	"context"
	"fmt"
	"gateway/config"
	"gateway/internal/server"
	casbinPkg "gateway/pkg/casbin"
	"gateway/pkg/cors"
	"gateway/pkg/errors"
	jaegerPkg "gateway/pkg/jaeger"
	"gateway/pkg/limiter"
	loggerPkg "gateway/pkg/logger"
	minioPkg "gateway/pkg/minio"
	"gateway/pkg/token"
	"gateway/pkg/utils"
	"log"
	"os"

	"github.com/opentracing/opentracing-go"
	"golang.org/x/sync/errgroup"
)

func main() {
	configPath := utils.GetConfigPath(os.Getenv("CONFIG_PATH"))
	fmt.Println(configPath)
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

	minioClient, err := minioPkg.NewMinioClient(ctx, cfg)
	if err != nil {
		logger.Fatalf("Minio init: %s", err)
	}
	logger.Info("Minio connected")

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

	rscors := cors.NewCors(cfg)

	gatewayServer := server.NewGatewayServer(
		ctx, logger, cfg,
		// metrics,
		tokenMaker,
		minioClient,
		casbinClient,
		limiter,
		rscors,
	)

	g.Go(func() error {
		return gatewayServer.RunMetrics(cancel)
	})

	g.Go(func() error {
		return gatewayServer.RunGateway()
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
