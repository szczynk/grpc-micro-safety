package server

import (
	"context"
	"mail/config"
	kafkaConsumer "mail/internal/user/delivery/kafka"
	userRepository "mail/internal/user/repository"
	userUseCase "mail/internal/user/usecase"
	kafkaClient "mail/pkg/kafka"
	"mail/pkg/logger"
	"mail/pkg/mailer"
	"time"

	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
)

// GRPC Auth Server
type Server struct {
	ctx         context.Context
	logger      logger.Logger
	cfg         *config.Config
	mailer      *mailer.Mailer
	db          *gorm.DB
	redisClient *redis.Client
	// metrics     metric.Metrics
	kafkaProducer kafkaClient.Producer
}

// Server constructor
func NewMailServer(
	ctx context.Context,
	logger logger.Logger,
	cfg *config.Config,
	mailer *mailer.Mailer,
	db *gorm.DB,
	redisClient *redis.Client,
	// metrics metric.Metrics,
	kafkaProducer kafkaClient.Producer,
) *Server {
	return &Server{
		ctx:         ctx,
		logger:      logger,
		cfg:         cfg,
		mailer:      mailer,
		db:          db,
		redisClient: redisClient,
		// metrics:     metrics,
		kafkaProducer: kafkaProducer,
	}
}

// Run service
func (s *Server) Run() error {
	errCh := make(chan error)

	ctx, cancel := context.WithCancel(s.ctx)

	userRepo := userRepository.NewUserRepository(s.db)
	userRedisRepo := userRepository.NewUserRedisRepository(s.redisClient)
	userUC := userUseCase.NewUserUseCase(s.logger, userRepo, userRedisRepo)

	userMessageProcessor := kafkaConsumer.NewUserMessageProcessor(s.logger, s.cfg, s.mailer, userUC, s.kafkaProducer)

	userConsumerGroup := kafkaClient.NewConsumerGroup(s.logger, s.cfg.Kafka.Brokers, s.cfg.Kafka.GroupID)

	topics := []string{"mail_verify_email", "mail_reset_password"}
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
			cancel()
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
