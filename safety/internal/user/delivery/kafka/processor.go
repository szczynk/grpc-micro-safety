package kafka

import (
	"context"
	"safety/config"
	"safety/internal/domain"
	"safety/pkg/logger"
	"time"

	"github.com/avast/retry-go"
	"github.com/segmentio/kafka-go"
)

const (
	PoolSize      = 30
	retryAttempts = 3
	retryDelay    = 300 * time.Millisecond
)

var (
	retryOptions = []retry.Option{retry.Attempts(retryAttempts), retry.Delay(retryDelay), retry.DelayType(retry.BackOffDelay)}
)

type userMessageProcessor struct {
	log    logger.Logger
	cfg    *config.Config
	userUC domain.UserUseCase
}

func NewUserMessageProcessor(
	log logger.Logger,
	cfg *config.Config,
	userUC domain.UserUseCase,
) *userMessageProcessor {
	return &userMessageProcessor{
		log:    log,
		cfg:    cfg,
		userUC: userUC,
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
		case "safety_create_user":
			s.processCreateUser(ctx, r, m)
		case "safety_update_user_reset_password":
			s.processUpdateUserResetPassword(ctx, r, m)
		case "safety_update_user_verify_email":
			s.processUpdateUserVerifyEmail(ctx, r, m)
		case "safety_update_user_by_id":
			s.processUpdateUserById(ctx, r, m)
		case "safety_delete_user_by_id":
			s.processDeleteUserById(ctx, r, m)
		}
	}
}
