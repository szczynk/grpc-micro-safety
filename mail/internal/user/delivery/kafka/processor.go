package kafka

import (
	"context"
	"mail/config"
	"mail/internal/domain"
	kafkaClient "mail/pkg/kafka"
	"mail/pkg/logger"
	"mail/pkg/mailer"
	"time"

	"github.com/avast/retry-go"
	"github.com/segmentio/kafka-go"
)

const (
	PoolSize      = 30
	retryAttempts = 3
	retryDelay    = 5 * time.Second
)

var (
	retryOptions = []retry.Option{retry.Attempts(retryAttempts), retry.Delay(retryDelay), retry.DelayType(retry.BackOffDelay)}
)

type userMessageProcessor struct {
	log           logger.Logger
	cfg           *config.Config
	mailer        *mailer.Mailer
	userUC        domain.UserUseCase
	kafkaProducer kafkaClient.Producer
}

func NewUserMessageProcessor(
	log logger.Logger,
	cfg *config.Config,
	mailer *mailer.Mailer,
	userUC domain.UserUseCase,
	kafkaProducer kafkaClient.Producer,
) *userMessageProcessor {
	return &userMessageProcessor{
		log:           log,
		cfg:           cfg,
		mailer:        mailer,
		userUC:        userUC,
		kafkaProducer: kafkaProducer,
	}
}

func (s *userMessageProcessor) ProcessMessagesErrGroup(ctx context.Context, r *kafka.Reader, workerID int) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		m, err := r.FetchMessage(ctx)
		if err != nil {
			s.log.Warnf("(ProcessMessages) workerID: %d, err: %v", workerID, err)
			continue
		}

		s.log.KafkaProcessMessage(m.Topic, m.Partition, m.Value, workerID, m.Offset, m.Time)

		switch m.Topic {
		case "mail_verify_email":
			s.processVerifyEmail(ctx, r, m)
		case "mail_reset_Password":
			s.processResetPassword(ctx, r, m)
		}
	}
}
