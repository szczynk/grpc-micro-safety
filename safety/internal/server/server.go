package server

import (
	"context"
	"net"
	"safety/config"
	attendanceGRPCService "safety/internal/attendance/delivery/grpc/service"
	attendanceRepository "safety/internal/attendance/repository"
	attendanceUseCase "safety/internal/attendance/usecase"
	certificateGRPCService "safety/internal/certificate/delivery/grpc/service"
	certificateRepository "safety/internal/certificate/repository"
	certificateUseCase "safety/internal/certificate/usecase"
	checkGRPCService "safety/internal/check/delivery/grpc/service"
	"safety/internal/interceptors"
	officeGRPCService "safety/internal/office/delivery/grpc/service"
	officeRepository "safety/internal/office/repository"
	officeUseCase "safety/internal/office/usecase"
	scheduleGRPCService "safety/internal/schedule/delivery/grpc/service"
	scheduleRepository "safety/internal/schedule/repository"
	scheduleUseCase "safety/internal/schedule/usecase"
	kafkaConsumer "safety/internal/user/delivery/kafka"
	userRepository "safety/internal/user/repository"
	userUseCase "safety/internal/user/usecase"
	workspaceGRPCService "safety/internal/workspace/delivery/grpc/service"
	workspaceRepository "safety/internal/workspace/repository"
	workspaceUseCase "safety/internal/workspace/usecase"
	"safety/pb"
	kafkaClient "safety/pkg/kafka"
	"safety/pkg/limiter"
	"safety/pkg/logger"
	"safety/pkg/token"
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
	tokenMaker   *token.PasetoMaker
	healthClient grpc_health_v1.HealthClient
	casbinClient *client.Enforcer
	limiter      *limiter.RateLimit
}

// Server constructor
func NewSafetyServer(
	ctx context.Context,
	logger logger.Logger,
	cfg *config.Config,
	db *gorm.DB,
	redisClient *redis.Client,
	// metrics metric.Metrics,
	tokenMaker *token.PasetoMaker,
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
		tokenMaker:   tokenMaker,
		healthClient: healthClient,
		casbinClient: casbinClient,
		limiter:      limiter,
	}
}

// Run service
func (s *Server) Run() error {
	errCh := make(chan error)

	im := interceptors.NewInterceptorManager(
		s.cfg, s.tokenMaker, s.casbinClient, s.limiter,
	)

	userRepo := userRepository.NewUserRepository(s.db)
	userUC := userUseCase.NewUserUseCase(s.logger, userRepo)

	officeRepo := officeRepository.NewOfficeRepository(s.db)
	officeRedisRepo := officeRepository.NewOfficeRedisRepository(s.redisClient)
	officeUC := officeUseCase.NewOfficeUseCase(s.logger, officeRepo, officeRedisRepo)

	workspaceRepo := workspaceRepository.NewWorkspaceRepository(s.db)
	workspaceRedisRepo := workspaceRepository.NewWorkspaceRedisRepository(s.redisClient)
	workspaceUC := workspaceUseCase.NewWorkspaceUseCase(s.logger, workspaceRepo, workspaceRedisRepo)

	scheduleRepo := scheduleRepository.NewScheduleRepository(s.db)
	scheduleRedisRepo := scheduleRepository.NewScheduleRedisRepository(s.redisClient)
	scheduleUC := scheduleUseCase.NewScheduleUseCase(s.logger, scheduleRepo, scheduleRedisRepo)

	attendanceRepo := attendanceRepository.NewAttendanceRepository(s.db)
	attendanceRedisRepo := attendanceRepository.NewAttendanceRedisRepository(s.redisClient)
	attendanceUC := attendanceUseCase.NewAttendanceUseCase(s.logger, attendanceRepo, attendanceRedisRepo)

	certificateRepo := certificateRepository.NewCertificateRepository(s.db)
	certificateRedisRepo := certificateRepository.NewCertificateRedisRepository(s.redisClient)
	certificateUC := certificateUseCase.NewCertificateUseCase(s.logger, certificateRepo, certificateRedisRepo)

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
				im.AuthStream,
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
				im.AuthUnary,
				grpc_recovery.UnaryServerInterceptor(),
			),
		),
	)

	if s.cfg.Server.Mode != "production" {
		reflection.Register(server)
	}

	// service
	officeGRPCService := officeGRPCService.NewOfficeGRPCService(
		s.logger, s.cfg, officeUC,
	)
	workspaceGRPCService := workspaceGRPCService.NewWorkspaceGRPCService(
		s.logger, s.cfg, workspaceUC,
	)
	scheduleGRPCService := scheduleGRPCService.NewScheduleGRPCService(
		s.logger, s.cfg, scheduleUC,
	)
	attendanceGRPCService := attendanceGRPCService.NewAttendanceGRPCService(
		s.logger, s.cfg, attendanceUC, certificateUC, scheduleUC,
	)
	certificateGRPCService := certificateGRPCService.NewCertificateGRPCService(
		s.logger, s.cfg, certificateUC,
	)
	checkGRPCService := checkGRPCService.NewCheckGRPCService(
		s.logger, s.cfg, attendanceUC,
	)

	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	pb.RegisterOfficeServiceServer(server, officeGRPCService)
	pb.RegisterWorkspaceServiceServer(server, workspaceGRPCService)
	pb.RegisterScheduleServiceServer(server, scheduleGRPCService)
	pb.RegisterAttendanceServiceServer(server, attendanceGRPCService)
	pb.RegisterCertificateServiceServer(server, certificateGRPCService)
	pb.RegisterCheckServiceServer(server, checkGRPCService)

	grpc_prometheus.Register(server)

	go func() {
		s.logger.Infof("Server is listening on port: %v", s.cfg.Server.Port)
		// s.logger.Infof("app info: %#v", server.GetServiceInfo())

		errCh <- server.Serve(listener)
	}()

	ctx, cancel := context.WithCancel(s.ctx)

	userMessageProcessor := kafkaConsumer.NewUserMessageProcessor(s.logger, s.cfg, userUC)

	userConsumerGroup := kafkaClient.NewConsumerGroup(s.logger, s.cfg.Kafka.Brokers, s.cfg.Kafka.GroupID)

	topics := []string{
		"safety_create_user",
		"safety_update_user_by_id",
		"safety_update_user_reset_password",
		"safety_update_user_verify_email",
		"safety_delete_user_by_id",
	}
	go func() {
		err := userConsumerGroup.ConsumeTopicWithErrGroup(
			ctx,
			topics,
			kafkaConsumer.PoolSize,
			userMessageProcessor.ProcessMessagesErrGroup,
		)
		if err != nil {
			s.logger.Errorf("(userConsumerGroup ConsumeTopicWithErrGroup) err: %v", err)
			cancel()
			return
		}
		errCh <- err
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
