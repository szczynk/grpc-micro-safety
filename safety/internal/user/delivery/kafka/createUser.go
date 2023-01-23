package kafka

import (
	"context"

	"safety/internal/models"
	"safety/pb"

	"github.com/avast/retry-go"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

func (s *userMessageProcessor) processCreateUser(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	// when unmarshaled it will be this
	// msg := &pb.UserKafka{
	// 	Id:       newUser.ID.String(),
	// 	Email:    newUser.Email,
	// 	Username: newUser.Username,
	// 	Password: newUser.Password,
	// 	Role:     newUser.Role,
	// 	Avatar:   newUser.Avatar,
	// 	Verified: newUser.Verified,
	// }
	var msg = new(pb.UserKafka)
	err := proto.Unmarshal(m.Value, msg)
	if err != nil {
		s.log.WarnErrMsg("proto.Unmarshal", err)
		s.commitErrMessage(ctx, r, m)
		return
	}

	err = msg.ValidateAll()
	if err != nil {
		s.log.WarnErrMsg("validate", err)
		s.commitErrMessage(ctx, r, m)
		return
	}

	userId := msg.GetId()
	parsedUserId, err := uuid.Parse(userId)
	if err != nil {
		s.log.WarnErrMsg("uuid.Parse", err)
		s.commitErrMessage(ctx, r, m)
		return
	}

	err = retry.Do(func() error {
		newUser := &models.User{
			ID:       parsedUserId,
			Username: msg.GetUsername(),
			Email:    msg.GetEmail(),
			Password: msg.GetPassword(),
			Role:     msg.GetRole(),
			Avatar:   msg.GetAvatar(),
			Verified: msg.GetVerified(),
		}
		_, err := s.userUC.CreateUser(ctx, newUser)
		if err != nil {
			return err
		}

		return nil
	}, append(retryOptions, retry.Context(ctx))...)

	if err != nil {
		s.log.WarnErrMsg("processCreateUser", err)
		return
	}

	s.commitMessage(ctx, r, m)
}
