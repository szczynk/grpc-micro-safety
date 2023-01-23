package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

func (s *userMessageProcessor) commitMessage(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	s.log.KafkaLogCommittedMessage(m.Topic, m.Partition, m.Offset)

	err := r.CommitMessages(ctx, m)
	if err != nil {
		s.log.WarnErrMsg("commitMessage", err)
	}
}

func (s *userMessageProcessor) commitErrMessage(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	s.log.KafkaLogCommittedMessage(m.Topic, m.Partition, m.Offset)

	err := r.CommitMessages(ctx, m)
	if err != nil {
		s.log.WarnErrMsg("commitMessage", err)
	}
}
