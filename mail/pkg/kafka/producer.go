package kafka

import (
	"context"
	"mail/config"
	"mail/pkg/logger"

	"github.com/segmentio/kafka-go"
)

type Producer interface {
	PublishMessage(ctx context.Context, msgs ...kafka.Message) error
	Close() error
}

type producer struct {
	log     logger.Logger
	brokers []string
	w       *kafka.Writer
}

// NewProducer create new kafka producer
func NewProducer(log logger.Logger, cfg *config.Config) *producer {
	brokers := cfg.Kafka.Brokers
	return &producer{
		log:     log,
		brokers: brokers,
		w:       NewKafkaWriter(brokers, kafka.LoggerFunc(log.Errorf)),
	}
}

func (p *producer) PublishMessage(ctx context.Context, msgs ...kafka.Message) error {
	return p.w.WriteMessages(ctx, msgs...)
}

func (p *producer) Close() error {
	return p.w.Close()
}
