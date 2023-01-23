package kafka

import (
	"context"
	"mail/pb"
	"time"

	"mail/internal/models"
	mailerPkg "mail/pkg/mailer"
	"mail/pkg/utils"

	"github.com/avast/retry-go"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/thanhpk/randstr"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *userMessageProcessor) processResetPassword(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	// when unmarshaled it will be this
	// msg := &pb.UserKafka{
	// 	Id:       user.ID.String(),
	// 	Username: user.Username,
	// 	Email:    user.Email,
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

	// Generate reset password token
	resetToken := randstr.String(64)
	encodedResetToken := utils.EncodeBase64(resetToken)

	err = retry.Do(func() error {
		resetTokenAt := time.Now().Add(time.Minute * 15)

		_, err := s.userUC.UpdateByID(ctx, parsedUserId, models.User{
			PasswordResetToken: resetToken,
			PasswordResetAt:    resetTokenAt,
		})
		if err != nil {
			return err
		}

		msgUserKafka := &pb.UserKafka{
			Id:                 userId,
			PasswordResetToken: resetToken,
			PasswordResetAt:    timestamppb.New(resetTokenAt),
		}
		msgBytes, err := proto.Marshal(msgUserKafka)
		if err != nil {
			return err
		}

		message := kafka.Message{
			Topic: "safety_update_user_by_id",
			Value: msgBytes,
		}
		err = s.kafkaProducer.PublishMessage(ctx, message)
		if err != nil {
			return err
		}

		emailData := mailerPkg.EmailData{
			URL:      "http://localhost:3000/reset-password/" + encodedResetToken,
			Email:    msg.GetEmail(),
			Username: msg.GetUsername(),
			Subject:  "Reset Password (Valid for 10 minutes)",
		}
		err = s.mailer.SendEmail(&emailData, "resetPassword.html")
		if err != nil {
			return err
		}

		return nil
	}, append(retryOptions, retry.Context(ctx))...)

	if err != nil {
		s.log.WarnErrMsg("processResetPassword", err)
		return
	}

	s.commitMessage(ctx, r, m)
}
