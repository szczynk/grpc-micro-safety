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

func (s *userMessageProcessor) processUpdateUserById(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	// when unmarshaled it will be this
	// msg := &pb.UserKafka{
	// 	Id:       updatedUser.ID.String(),
	// 	Email:    updatedUser.Email,
	// 	Username: updatedUser.Username,
	// 	Password: updatedUser.Password,
	// 	Role:     updatedUser.Role,
	// 	Avatar:   updatedUser.Avatar,
	// 	Verified: updatedUser.Verified,
	// 	VerificationCode: updatedUser.VerificationCode,
	// 	PasswordResetToken: updatedUser.PasswordResetToken,
	// 	PasswordResetAt: updatedUser.PasswordResetAt,
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
		_, err := s.userUC.UpdateByID(ctx, parsedUserId, models.User{
			Username:           msg.GetUsername(),
			Email:              msg.GetEmail(),
			Password:           msg.GetPassword(),
			Role:               msg.GetRole(),
			Avatar:             msg.GetAvatar(),
			Verified:           msg.GetVerified(),
			VerificationCode:   msg.GetVerificationCode(),
			PasswordResetToken: msg.GetPasswordResetToken(),
			PasswordResetAt:    msg.GetPasswordResetAt().AsTime(),
		})
		if err != nil {
			return err
		}

		return nil
	}, append(retryOptions, retry.Context(ctx))...)

	if err != nil {
		s.log.WarnErrMsg("processUpdateUserById", err)
		return
	}

	s.commitMessage(ctx, r, m)
}
