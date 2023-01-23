package server

import (
	"auth/config"
	"auth/internal/interceptors"
	sessionRepository "auth/internal/session/repository"
	sessionUseCase "auth/internal/session/usecase"
	authGRPCService "auth/internal/user/delivery/grpc/service"
	userRepository "auth/internal/user/repository"
	userUseCase "auth/internal/user/usecase"
	"auth/pb"
	kafkaClient "auth/pkg/kafka"
	"auth/pkg/limiter"
	"auth/pkg/logger"
	"auth/pkg/token"
	"context"
	"net"
	"time"

	"github.com/casbin/casbin-go-client/client"
	"github.com/go-redis/redis/v9"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

// GRPC Auth Server
type Server struct {
	ctx         context.Context
	logger      logger.Logger
	cfg         *config.Config
	db          *gorm.DB
	redisClient *redis.Client
	// metrics       metric.Metrics
	tokenMaker    *token.PasetoMaker
	kafkaProducer kafkaClient.Producer
	healthClient  grpc_health_v1.HealthClient
	casbinClient  *client.Enforcer
	limiter       *limiter.RateLimit
}

// Server constructor
func NewAuthServer(
	ctx context.Context,
	logger logger.Logger,
	cfg *config.Config,
	db *gorm.DB,
	redisClient *redis.Client,
	// metrics metric.Metrics,
	tokenMaker *token.PasetoMaker,
	kafkaProducer kafkaClient.Producer,
	healthClient grpc_health_v1.HealthClient,
	casbinClient *client.Enforcer,
	limiter *limiter.RateLimit,
) *Server {
	return &Server{
		ctx:         ctx,
		logger:      logger,
		cfg:         cfg,
		db:          db,
		redisClient: redisClient,
		// metrics:       metrics,
		tokenMaker:    tokenMaker,
		kafkaProducer: kafkaProducer,
		healthClient:  healthClient,
		casbinClient:  casbinClient,
		limiter:       limiter,
	}
}

// Run service
func (s *Server) Run() error {
	errCh := make(chan error)

	im := interceptors.NewInterceptorManager(
		s.cfg, s.limiter,
	)

	userRepo := userRepository.NewUserRepository(s.db)
	userRedisRepo := userRepository.NewUserRedisRepository(s.redisClient)
	sessionRepo := sessionRepository.NewSessionRepository(s.redisClient)
	userUC := userUseCase.NewUserUseCase(s.logger, userRepo, userRedisRepo)
	sessionUC := sessionUseCase.NewSessionUseCase(s.logger, sessionRepo)

	listener, err := net.Listen("tcp", s.cfg.Server.Port)
	if err != nil {
		return err
	}
	defer listener.Close()

	zapLoggerOpts := []grpc_zap.Option{
		grpc_zap.WithDecider(
			func(fullMethodName string, err error) bool {
				if err == nil && fullMethodName == "/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo" {
					return false
				}
				return true
			},
		),
	}

	server := grpc.NewServer(
		grpc.KeepaliveParams(
			keepalive.ServerParameters{
				MaxConnectionIdle: s.cfg.Server.MaxConnectionIdle,
				Timeout:           s.cfg.Server.Timeout,
				MaxConnectionAge:  s.cfg.Server.MaxConnectionAge,
				Time:              s.cfg.Server.Timeout,
			},
		),
		//* https://github.com/grpc-ecosystem/go-grpc-middleware
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				grpc_ctxtags.StreamServerInterceptor(),
				grpc_prometheus.StreamServerInterceptor,
				grpc_zap.StreamServerInterceptor(
					s.logger.ZapLogger(), zapLoggerOpts...,
				),
				im.RateLimitStream,
				grpc_validator.StreamServerInterceptor(),
				grpc_recovery.StreamServerInterceptor(),
			)),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_ctxtags.UnaryServerInterceptor(),
				grpc_prometheus.UnaryServerInterceptor,
				grpc_zap.UnaryServerInterceptor(
					s.logger.ZapLogger(), zapLoggerOpts...,
				),
				im.RateLimitUnary,
				grpc_validator.UnaryServerInterceptor(),
				grpc_recovery.UnaryServerInterceptor(),
			),
		),
	)

	if s.cfg.Server.Mode != "production" {
		reflection.Register(server)
	}

	// server
	authGRPCService := authGRPCService.NewAuthGRPCService(
		s.logger, s.cfg, userUC, sessionUC,
		s.tokenMaker, s.kafkaProducer, s.casbinClient,
	)

	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	pb.RegisterAuthServiceServer(server, authGRPCService)

	grpc_prometheus.Register(server)

	go func() {
		s.logger.Infof("Server is listening on port: %v", s.cfg.Server.Port)
		// s.logger.Infof("app info: %+v", server.GetServiceInfo())

		errCh <- server.Serve(listener)
	}()

	select {
	case <-s.ctx.Done():
		c := make(chan bool)
		go func() {
			defer close(c)
			server.GracefulStop()
			listener.Close()
		}()
		select {
		case <-c:
		case <-time.After(5 * time.Second):
		}
		s.logger.Info("Server Exited Properly")
		return nil
	case err := <-errCh:
		return err
	}
}
