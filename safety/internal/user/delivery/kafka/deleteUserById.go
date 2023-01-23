package kafka

import (
	"context"
	"safety/pb"

	"github.com/avast/retry-go"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

func (s *userMessageProcessor) processDeleteUserById(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	// when unmarshaled it will be this
	// msg := &pb.UserKafka{
	// 	Id: user.ID.String(),
	// }
	var msg = new(pb.UserKafka)
	err := proto.Unmarshal(m.Value, msg)
	if err != nil {
		s.log.WarnErrMsg("proto.Unmarshal", err)
		s.commitErrMessage(ctx, r, m)
		return
	}

	// err = msg.ValidateAll()
	// if err != nil {
	// 	s.log.WarnErrMsg("validate", err)
	// 	s.commitErrMessage(ctx, r, m)
	// 	return
	// }

	userId := msg.GetId()
	parsedUserId, err := uuid.Parse(userId)
	if err != nil {
		s.log.WarnErrMsg("uuid.Parse", err)
		s.commitErrMessage(ctx, r, m)
		return
	}

	err = retry.Do(func() error {
		err := s.userUC.DeleteByID(ctx, parsedUserId)
		if err != nil {
			return err
		}

		return nil
	}, append(retryOptions, retry.Context(ctx))...)

	if err != nil {
		s.log.WarnErrMsg("processDeleteUserById", err)
		return
	}

	s.commitMessage(ctx, r, m)
}
