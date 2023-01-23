package kafka

import (
	"context"
	"mail/pb"

	"mail/internal/models"
	mailerPkg "mail/pkg/mailer"
	"mail/pkg/utils"

	"github.com/avast/retry-go"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/thanhpk/randstr"
	"google.golang.org/protobuf/proto"
)

func (s *userMessageProcessor) processVerifyEmail(ctx context.Context, r *kafka.Reader, m kafka.Message) {
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

	code := randstr.String(64)
	encodedCode := utils.EncodeBase64(code)

	err = retry.Do(func() error {
		_, err := s.userUC.UpdateByID(ctx, parsedUserId, models.User{
			VerificationCode: code,
		})
		if err != nil {
			return err
		}

		msgUserKafka := &pb.UserKafka{
			Id:               userId,
			VerificationCode: code,
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
			URL:      "http://localhost:3000/verify-email/" + encodedCode,
			Email:    msg.GetEmail(),
			Username: msg.GetUsername(),
			Subject:  "Email Verification",
		}
		err = s.mailer.SendEmail(&emailData, "verificationCode.html")
		if err != nil {
			return err
		}

		return nil
	}, append(retryOptions, retry.Context(ctx))...)

	if err != nil {
		s.log.WarnErrMsg("processVerifyEmail", err)
		return
	}

	s.commitMessage(ctx, r, m)
}
