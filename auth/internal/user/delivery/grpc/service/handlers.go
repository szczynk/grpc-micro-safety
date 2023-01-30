package service

import (
	"auth/config"
	"auth/internal/domain"
	"auth/internal/models"
	"auth/pb"
	"auth/pkg/grpc_errors"
	kafkaClient "auth/pkg/kafka"
	"auth/pkg/logger"
	"auth/pkg/token"
	"auth/pkg/utils"
	"context"
	"fmt"
	"strings"
	"time"

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
	pb.UnimplementedAuthServiceServer
	logger        logger.Logger
	cfg           *config.Config
	userUC        domain.UserUseCase
	sessionUC     domain.SessionUseCase
	tokenMaker    *token.PasetoMaker
	kafkaProducer kafkaClient.Producer
	casbin        *client.Enforcer
}

// Auth service constructor
func NewAuthGRPCService(
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

func (u *usersService) Register(ctx context.Context, r *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "user.Register")
	defer span.Finish()

	username := strings.TrimSpace(r.GetUsername())
	email := strings.ToLower(strings.TrimSpace(r.GetEmail()))
	role := "user"
	avatar := fmt.Sprintf("https://avatars.dicebear.com/api/open-peeps/%s.svg", username)

	hashedPassword, err := utils.HashPassword(r.GetPassword())
	if err != nil {
		return nil, err
	}

	newUser := &models.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
		Role:     role,
		Avatar:   avatar,
	}

	registeredUser, err := u.userUC.CreateUser(ctx, newUser)
	if err != nil {
		u.logger.Errorf("userUC.CreateUser: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "Register: %v", err)
	}

	_, err = u.casbin.AddRoleForUser(ctx, registeredUser.ID.String(), registeredUser.Role)
	if err != nil {
		u.logger.Errorf("casbin.AddRoleForUser: %v", err)
	}

	msg := &pb.UserKafka{
		Id:       registeredUser.ID.String(),
		Username: registeredUser.Username,
		Email:    registeredUser.Email,
		Password: registeredUser.Password,
		Role:     registeredUser.Role,
		Avatar:   registeredUser.Avatar,
	}
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		u.logger.Errorf("proto.Marshal: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "proto.Marshal: %v", err)
	}

	message := kafka.Message{
		Topic: "safety_create_user",
		Value: msgBytes,
	}
	err = u.kafkaProducer.PublishMessage(ctx, message)
	if err != nil {
		u.logger.Errorf("SendCreateUserKafkaFailed: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "SendCreateUserKafkaFailed: %v", err)
	}

	msg2 := &pb.UserKafka{
		Id:       registeredUser.ID.String(),
		Username: registeredUser.Username,
		Email:    registeredUser.Email,
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

	err = u.SendHeader(ctx)
	if err != nil {
		return nil, err
	}

	successMsg := fmt.Sprintf("We sent an email with a verification code to %s.", registeredUser.Email)
	return &pb.RegisterResponse{User: u.UserModelToProto(registeredUser), Message: successMsg}, nil
}

func (u *usersService) Login(ctx context.Context, r *pb.LoginRequest) (*pb.LoginResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "user.Login")
	defer span.Finish()

	email := strings.ToLower(strings.TrimSpace(r.GetEmail()))
	password := r.GetPassword()

	user, err := u.userUC.Login(ctx, email, password)
	if err != nil {
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "Login: %v", err)
	}

	accessToken, accessPayload, err := u.tokenMaker.CreateToken(user.ID, user.Role, u.cfg.AccessToken.Expire)
	if err != nil {
		u.logger.Errorf("CreateAccessToken: %v", err)
		return nil, status.Errorf(codes.Internal, "CreateAccessToken: %v", err)
	}

	refreshToken, refreshPayload, err := u.tokenMaker.CreateToken(user.ID, user.Role, u.cfg.RefreshToken.MaxAge)
	if err != nil {
		u.logger.Errorf("CreateRefreshToken: %v", err)
		return nil, status.Errorf(codes.Internal, "CreateRefreshToken: %v", err)
	}

	mtdt, err := u.ExtractMetadata(ctx)
	if err != nil {
		u.logger.Errorf("ExtractMetadata: %v", err)
		return nil, status.Errorf(codes.Internal, "ExtractMetadata: %v", err)
	}

	parsedSessionID, err := uuid.Parse(refreshPayload.Get("session_id"))
	if err != nil {
		u.logger.Errorf("ConvertSessionID: %v", err)
		return nil, status.Errorf(codes.Internal, "ConvertSessionID: %v", err)
	}

	newSession := &models.Session{
		ID:           parsedSessionID,
		UserID:       user.ID,
		RefreshToken: refreshToken,
		UserAgent:    mtdt.UserAgent,
		ClientIp:     mtdt.ClientIP,
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.Expiration,
	}
	err = u.sessionUC.CreateSession(ctx, u.cfg.RefreshToken.MaxAge, newSession)
	if err != nil {
		u.logger.Errorf("sessionUC.CreateSession: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "sessionUC.CreateSession: %v", err)
	}

	err = u.SendHeader(ctx)
	if err != nil {
		return nil, err
	}

	res := &pb.LoginResponse{
		User:                  u.UserModelToProto(user),
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.Expiration),
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.Expiration),
	}
	return res, nil
}

func (u *usersService) ForgotPassword(ctx context.Context, r *pb.ForgotPasswordRequest) (*pb.ForgotPasswordResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "user.ForgotPassword")
	defer span.Finish()

	email := strings.ToLower(strings.TrimSpace(r.GetEmail()))

	user, err := u.userUC.FindByEmail(ctx, email)
	if err != nil {
		u.logger.Errorf("userUC.FindByEmail: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "FindByEmail: %v", err)
	}

	if !user.Verified {
		u.logger.Error("User not verified")
		return &pb.ForgotPasswordResponse{Message: "Please verify your email"}, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "User not verified")
	}

	msg := &pb.UserKafka{
		Id:       user.ID.String(),
		Username: user.Username,
		Email:    user.Email,
	}
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		u.logger.Errorf("proto.Marshal: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "proto.Marshal: %v", err)
	}

	message := kafka.Message{
		Topic: "mail_reset_password",
		Value: msgBytes,
	}
	err = u.kafkaProducer.PublishMessage(ctx, message)
	if err != nil {
		u.logger.Errorf("SendResetPassTokenFailed: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "SendResetPassTokenFailed: %v", err)
	}

	err = u.SendHeader(ctx)
	if err != nil {
		return nil, err
	}

	successMsg := fmt.Sprintf("We sent an email with a password reset token to %s.", user.Email)
	return &pb.ForgotPasswordResponse{Message: successMsg}, nil
}

func (u *usersService) ResetPassword(ctx context.Context, r *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "user.ResetPassword")
	defer span.Finish()

	resetToken := r.GetResetToken()
	if len(resetToken) == 0 {
		err := grpc_errors.ErrNoResetPassToken
		u.logger.Error(err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), err.Error())
	}

	decodedResetToken, err := utils.DecodeBase64(resetToken)
	if err != nil {
		u.logger.Errorf("decodeResetPassToken: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "decodeResetPassToken: %v", err)
	}

	user, err := u.userUC.FindByResetPasswordToken(ctx, decodedResetToken)
	if err != nil {
		u.logger.Errorf("userUC.FindByResetPasswordToken: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "findByResetPasswordToken: %v", err)
	}

	if !user.Verified {
		u.logger.Error("User not verified")
		return &pb.ResetPasswordResponse{Message: "Please verify your email"}, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "User not verified")
	}

	hashedPassword, err := utils.HashPassword(r.GetPassword())
	if err != nil {
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "hashPassword: %v", err)
	}

	_, err = u.userUC.UpdateByEmail(ctx, user.Email, map[string]interface{}{
		"password":             hashedPassword,
		"password_reset_token": "",
	})
	if err != nil {
		u.logger.Errorf("userUC.UpdateByEmail: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "UpdatePassword: %v", err)
	}

	msg := &pb.UserKafka{
		Email:    user.Email,
		Password: hashedPassword,
	}
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		u.logger.Errorf("proto.Marshal: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "proto.Marshal: %v", err)
	}

	message := kafka.Message{
		Topic: "safety_update_user_reset_password",
		Value: msgBytes,
	}
	err = u.kafkaProducer.PublishMessage(ctx, message)
	if err != nil {
		u.logger.Errorf("SendUpdateUserKafkaResetPasswordFailed: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "SendUpdateUserKafkaResetPasswordFailed: %v", err)
	}

	err = u.SendHeader(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.ResetPasswordResponse{Message: "Reset password successful"}, nil
}

func (u *usersService) VerifyEmail(ctx context.Context, r *pb.VerifyEmailRequest) (*pb.VerifyEmailResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "user.VerifyEmail")
	defer span.Finish()

	code := r.GetCode()
	if len(code) == 0 {
		err := grpc_errors.ErrNoEmailVerifyCode
		u.logger.Error(err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), err.Error())
	}

	decodedCode, err := utils.DecodeBase64(code)
	if err != nil {
		u.logger.Errorf("decodeVerificationCode: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "decodeVerificationCode: %v", err)
	}

	user, err := u.userUC.FindByVerificationCode(ctx, decodedCode)
	if err != nil {
		u.logger.Errorf("userUC.FindByVerificationCode: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "FindByVerificationCode: %v", err)
	}

	if user.Verified {
		u.logger.Error("User already verified")
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "User already verified")
	}

	verified := true
	verifiedAt := time.Now()

	_, err = u.userUC.UpdateByEmail(ctx, user.Email, map[string]interface{}{
		"verification_code": "",
		"verified":          verified,
		"verified_at":       verifiedAt,
	})
	if err != nil {
		u.logger.Errorf("userUC.UpdateByEmail: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "UpdateEmailVerification: %v", err)
	}

	msg := &pb.UserKafka{
		Email:      user.Email,
		Verified:   verified,
		VerifiedAt: timestamppb.New(verifiedAt),
	}
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		u.logger.Errorf("proto.Marshal: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "proto.Marshal: %v", err)
	}

	message := kafka.Message{
		Topic: "safety_update_user_verify_email",
		Value: msgBytes,
	}
	err = u.kafkaProducer.PublishMessage(ctx, message)
	if err != nil {
		u.logger.Errorf("SendUpdateUserKafkaVerifyEmailFailed: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "SendUpdateUserKafkaVerifyEmailFailed: %v", err)
	}

	err = u.SendHeader(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.VerifyEmailResponse{Message: "Email verified successfully"}, nil
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
