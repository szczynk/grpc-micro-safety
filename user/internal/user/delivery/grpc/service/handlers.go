package service

import (
	"context"
	"fmt"
	"strings"
	"user/config"
	"user/internal/domain"
	"user/internal/models"
	"user/pb"
	"user/pkg/grpc_errors"
	kafkaClient "user/pkg/kafka"
	"user/pkg/logger"
	"user/pkg/token"

	"github.com/casbin/casbin-go-client/client"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type usersService struct {
	pb.UnimplementedUserServiceServer
	logger        logger.Logger
	cfg           *config.Config
	userUC        domain.UserUseCase
	sessionUC     domain.SessionUseCase
	tokenMaker    *token.PasetoMaker
	kafkaProducer kafkaClient.Producer
	casbin        *client.Enforcer
}

// User service constructor
func NewUserGRPCService(
	logger logger.Logger,
	cfg *config.Config,
	userUC domain.UserUseCase,
	sessionUC domain.SessionUseCase,
	tokenMaker *token.PasetoMaker,
	kafkaProducer kafkaClient.Producer,
	casbin *client.Enforcer,
) *usersService {
	return &usersService{
		logger:        logger,
		cfg:           cfg,
		userUC:        userUC,
		sessionUC:     sessionUC,
		tokenMaker:    tokenMaker,
		kafkaProducer: kafkaProducer,
		casbin:        casbin,
	}
}

// *Command

func (u *usersService) RefreshAccessToken(ctx context.Context, r *pb.RefreshAccessTokenRequest) (*pb.RefreshAccessTokenResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "user.RefreshAccessToken")
	defer span.Finish()

	//* to renew access token, u need to post refresh token
	refreshToken := r.GetRefreshToken()
	refreshPayload, err := u.tokenMaker.VerifyToken(refreshToken, true)
	if err != nil {
		u.logger.Errorf("VerifyRefreshToken: %v", err)
		return nil, status.Errorf(codes.Internal, "VerifyRefreshToken: %v", err)
	}

	parsedSessionID, err := uuid.Parse(refreshPayload.Get("session_id"))
	if err != nil {
		u.logger.Errorf("ConvertSessionID: %v", err)
		return nil, status.Errorf(codes.Internal, "ConvertSessionID: %v", err)
	}

	parsedUserID, err := uuid.Parse(refreshPayload.Get("user_id"))
	if err != nil {
		u.logger.Errorf("ConvertUserID: %v", err)
		return nil, status.Errorf(codes.Internal, "ConvertUserID: %v", err)
	}

	session, err := u.sessionUC.GetSessionByID(ctx, parsedSessionID)
	if err != nil {
		u.logger.Errorf("sessionUC.GetSessionByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "sessionUC.GetSessionByID: %v", err)
	}

	if session.RefreshToken != refreshToken {
		u.logger.Errorf("SessionIncorrectRefreshToken: %v", err)
		return nil, status.Errorf(codes.Unauthenticated, "SessionIncorrectRefreshToken: %v", err)
	}

	if session.IsBlocked {
		u.logger.Errorf("SessionBlocked: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "SessionBlocked: %v", err)
	}

	if session.UserID != parsedUserID {
		u.logger.Errorf("SessionIncorrectUserID: %v", err)
		return nil, status.Errorf(codes.Unauthenticated, "SessionIncorrectUserID: %v", err)
	}

	accessToken, accessPayload, err := u.tokenMaker.CreateToken(parsedUserID, refreshPayload.Get("role"), u.cfg.AccessToken.Expire)
	if err != nil {
		u.logger.Errorf("CreateAccessToken: %v", err)
		return nil, status.Errorf(codes.Internal, "CreateAccessToken: %v", err)
	}

	res := &pb.RefreshAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: timestamppb.New(accessPayload.Expiration),
	}
	return res, nil
}

func (u *usersService) UpdateMe(ctx context.Context, r *pb.UpdateMeRequest) (*pb.UpdateMeResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "user.UpdateMe")
	defer span.Finish()

	userId, _, err := u.ExtractUserId(ctx)
	if err != nil {
		u.logger.Errorf("ExtractUserId: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "ExtractUserId: %v", err)
	}

	parsedUserID, err := uuid.Parse(userId)
	if err != nil {
		u.logger.Errorf("ConvertUserID: %v", err)
		return nil, status.Errorf(codes.Internal, "ConvertUserID: %v", err)
	}

	username := strings.TrimSpace(r.GetUsername())
	avatar := r.GetAvatar()

	user, err := u.userUC.UpdateByID(ctx, parsedUserID, models.User{
		Username: username,
		Avatar:   avatar,
	})
	if err != nil {
		u.logger.Errorf("userUC.UpdateByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "userUC.UpdateByID: %v", err)
	}

	msgUserKafka := &pb.UserKafka{
		Id:       user.ID.String(),
		Username: user.Username,
		Avatar:   user.Avatar,
	}
	msgBytes, err := proto.Marshal(msgUserKafka)
	if err != nil {
		u.logger.Errorf("proto.Marshal: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "proto.Marshal: %v", err)
	}

	message := kafka.Message{
		Topic: "safety_update_user_by_id",
		Value: msgBytes,
	}
	err = u.kafkaProducer.PublishMessage(ctx, message)
	if err != nil {
		u.logger.Errorf("SendUpdateUserKafkaByIdFailed: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "SendUpdateUserKafkaByIdFailed: %v", err)
	}

	return &pb.UpdateMeResponse{User: u.UserModelToProto(user)}, nil
}

func (u *usersService) ChangeEmail(ctx context.Context, r *pb.ChangeEmailRequest) (*pb.ChangeEmailResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "user.ChangeEmail")
	defer span.Finish()

	userId, _, err := u.ExtractUserId(ctx)
	if err != nil {
		u.logger.Errorf("ExtractUserId: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "ExtractUserId: %v", err)
	}

	parsedUserID, err := uuid.Parse(userId)
	if err != nil {
		u.logger.Errorf("ConvertUserID: %v", err)
		return nil, status.Errorf(codes.Internal, "ConvertUserID: %v", err)
	}

	email := strings.ToLower(strings.TrimSpace(r.GetEmail()))
	verified := false

	user, err := u.userUC.UpdateByID(ctx, parsedUserID, models.User{
		Email:    email,
		Verified: verified,
	})
	if err != nil {
		u.logger.Errorf("userUC.UpdateByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "userUC.UpdateByID: %v", err)
	}

	msg := &pb.UserKafka{
		Id:       user.ID.String(),
		Email:    user.Email,
		Username: user.Username,
		Verified: user.Verified,
	}
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "proto.Marshal: %v", err)
	}

	message := kafka.Message{
		Topic: "safety_update_user_by_id",
		Value: msgBytes,
	}
	err = u.kafkaProducer.PublishMessage(ctx, message)
	if err != nil {
		u.logger.Errorf("SendUpdateUserKafkaByIdFailed: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "SendUpdateUserKafkaByIdFailed: %v", err)
	}

	msg2 := &pb.UserKafka{
		Id:       user.ID.String(),
		Username: user.Username,
		Email:    user.Email,
	}
	msgBytes2, err := proto.Marshal(msg2)
	if err != nil {
		u.logger.Errorf("proto.Marshal: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "proto.Marshal: %v", err)
	}

	message2 := kafka.Message{
		Topic: "mail_verify_email",
		Value: msgBytes2,
	}
	err = u.kafkaProducer.PublishMessage(ctx, message2)
	if err != nil {
		u.logger.Errorf("SendVerifyEmailFailed: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "SendVerifyEmailFailed: %v", err)
	}

	successMsg := fmt.Sprintf("We sent an email with a verification code to %s.", user.Email)
	return &pb.ChangeEmailResponse{User: u.UserModelToProto(user), Message: successMsg}, nil
}

func (u *usersService) Logout(ctx context.Context, r *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "user.Logout")
	defer span.Finish()

	//* to logout, u need to post refresh token
	refreshToken := r.GetRefreshToken()
	refreshPayload, err := u.tokenMaker.VerifyToken(refreshToken, true)
	if err != nil {
		u.logger.Errorf("VerifyRefreshToken: %v", err)
		return nil, status.Errorf(codes.Internal, "VerifyRefreshToken: %v", err)
	}

	parsedSessionID, err := uuid.Parse(refreshPayload.Get("session_id"))
	if err != nil {
		u.logger.Errorf("ConvertSessionID: %v", err)
		return nil, status.Errorf(codes.Internal, "ConvertSessionID: %v", err)
	}

	err = u.sessionUC.DeleteByID(ctx, parsedSessionID)
	if err != nil {
		u.logger.Errorf("sessionUC.DeleteByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "sessionUC.GetSessionByID: %v", err)
	}

	return &pb.LogoutResponse{}, nil
}

// *Query

func (u *usersService) GetMe(ctx context.Context, r *pb.GetMeRequest) (*pb.GetMeResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "user.GetMe")
	defer span.Finish()

	userId, _, err := u.ExtractUserId(ctx)
	if err != nil {
		u.logger.Errorf("ExtractUserId: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "ExtractUserId: %v", err)
	}

	parsedUserID, err := uuid.Parse(userId)
	if err != nil {
		u.logger.Errorf("ConvertUserID: %v", err)
		return nil, status.Errorf(codes.Internal, "ConvertUserID: %v", err)
	}

	user, err := u.userUC.FindByID(ctx, parsedUserID, u.cfg.Cache.Expire)
	if err != nil {
		u.logger.Errorf("userUC.FindByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "userUC.FindByID: %v", err)
	}

	return &pb.GetMeResponse{User: u.UserModelToProto(user)}, nil
}

// * Utils

func (u *usersService) UserModelToProto(user *models.User) *pb.User {
	return &pb.User{
		Id:        user.ID.String(),
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		Avatar:    user.Avatar,
		Verified:  user.Verified,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}
