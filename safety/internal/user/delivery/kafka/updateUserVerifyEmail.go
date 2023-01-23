package kafka

import (
	"context"
	"safety/pb"

	"github.com/avast/retry-go"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

func (s *userMessageProcessor) processUpdateUserVerifyEmail(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	// when unmarshaled it will be this
	// msg := &pb.UserKafka{
	// 	Email:      updatedUser.Email,
	// 	Verified:   updatedUser.Verified,
	// 	VerifiedAt: timestamppb.New(updatedUser.VerifiedAt),
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

	err = retry.Do(func() error {
		_, err := s.userUC.UpdateByEmail(ctx, msg.Email, map[string]interface{}{
			"verification_code": "",
			"verified":          msg.GetVerified(),
			"verified_at":       msg.GetVerifiedAt().AsTime(),
		})
		if err != nil {
			return err
		}

		return nil
	}, append(retryOptions, retry.Context(ctx))...)

	if err != nil {
		s.log.WarnErrMsg("processUpdateUserVerifyEmail", err)
		return
	}

	s.commitMessage(ctx, r, m)
}
