package kafka

import (
	"context"
	"mail/config"

	"github.com/segmentio/kafka-go"
)

// NewKafkaConn create new kafka connection
func NewKafkaConn(ctx context.Context, cfg *config.Config) (*kafka.Conn, error) {
	return kafka.DialContext(ctx, "tcp", cfg.Kafka.Brokers[0])
}
