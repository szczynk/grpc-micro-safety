package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"user/internal/models"
	"user/pb"
	"user/pkg/grpc_errors"
	"user/pkg/utils"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// *Command

func (u *usersService) CreateUser(ctx context.Context, r *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "user.CreateUser")
	defer span.Finish()

	username := strings.TrimSpace(r.GetUsername())
	email := strings.ToLower(strings.TrimSpace(r.GetEmail()))
	role := strings.ToLower(r.GetRole())
	avatar := r.GetAvatar()
	verified := r.GetVerified()

	if len(avatar) == 0 {
		avatar = fmt.Sprintf("https://avatars.dicebear.com/api/open-peeps/%s.svg", username)
	}

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
		Verified: verified,
	}

	createdUser, err := u.userUC.CreateUser(ctx, newUser)
	if err != nil {
		u.logger.Errorf("userUC.CreateUser: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "CreateUser: %v", err)
	}

	_, err = u.casbin.AddRoleForUser(ctx, createdUser.ID.String(), createdUser.Role)
	if err != nil {
		u.logger.Errorf("casbin.AddRoleForUser: %v", err)
	}

	msg := &pb.UserKafka{
		Id:       newUser.ID.String(),
		Email:    newUser.Email,
		Username: newUser.Username,
		Password: newUser.Password,
		Role:     newUser.Role,
		Avatar:   newUser.Avatar,
		Verified: newUser.Verified,
	}
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
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

	var successMsg string
	if !newUser.Verified {
		msg2 := &pb.UserKafka{
			Id:       newUser.ID.String(),
			Username: newUser.Username,
			Email:    newUser.Email,
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

		successMsg = fmt.Sprintf("We sent an email with a verification code to %s.", newUser.Email)
	} else {
		successMsg = ""
	}

	err = u.SendHeader(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{User: u.UserModelToProto(createdUser), Message: successMsg}, nil
}

func (u *usersService) UpdateUserById(ctx context.Context, r *pb.UpdateUserByIdRequest) (*pb.UpdateUserByIdResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "user.UpdateUserByID")
	defer span.Finish()

	parsedUserID, err := uuid.Parse(r.GetId())
	if err != nil {
		u.logger.Errorf("ConvertUserID: %v", err)
		return nil, status.Errorf(codes.Internal, "ConvertUserID: %v", err)
	}

	foundUser, err := u.userUC.FindByID(ctx, parsedUserID, u.cfg.Cache.Expire)
	if err != nil {
		u.logger.Errorf("userUC.FindByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "userUC.FindByID: %v", err)
	}

	email := strings.ToLower(strings.TrimSpace(r.GetEmail()))
	role := strings.ToLower(r.GetRole())
	username := strings.TrimSpace(r.GetUsername())
	password := r.GetPassword()
	avatar := r.GetAvatar()
	verified := r.GetVerified()

	if len(avatar) == 0 {
		avatar = fmt.Sprintf("https://avatars.dicebear.com/api/open-peeps/%s.svg", username)
	}

	if len(role) > 0 {
		// err = u.casbin.LoadPolicy(ctx)
		// if err != nil {
		// 	u.logger.Errorf("casbin.LoadPolicy: %v", err)
		// 	return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "casbin.LoadPolicy: %v", err)
		// }

		_, err = u.casbin.DeleteRoleForUser(ctx, foundUser.ID.String(), foundUser.Role)
		if err != nil {
			u.logger.Errorf("casbin.DeleteRoleForUser: %v", err)
			return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "casbin.DeleteRoleForUser: %v", err)
		}

		_, err = u.casbin.AddRoleForUser(ctx, foundUser.ID.String(), role)
		if err != nil {
			u.logger.Errorf("casbin.AddRoleForUser: %v", err)
			return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "casbin.AddRoleForUser: %v", err)
		}
	}

	var hashedPassword string
	if len(password) > 0 {
		hashedPassword, err = utils.HashPassword(password)
		if err != nil {
			return nil, err
		}
	}

	updatedUser, err := u.userUC.UpdateByID(ctx, parsedUserID, models.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
		Role:     role,
		Avatar:   avatar,
		Verified: verified,
	})
	if err != nil {
		u.logger.Errorf("userUC.UpdateByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "userUC.UpdateByID: %v", err)
	}

	msg := &pb.UserKafka{
		Id:       updatedUser.ID.String(),
		Email:    updatedUser.Email,
		Username: updatedUser.Username,
		Password: updatedUser.Password,
		Role:     updatedUser.Role,
		Avatar:   updatedUser.Avatar,
		Verified: updatedUser.Verified,
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

	err = u.SendHeader(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateUserByIdResponse{User: u.UserModelToProto(updatedUser)}, nil
}

func (u *usersService) DeleteUserById(ctx context.Context, r *pb.DeleteUserByIdRequest) (*pb.DeleteUserByIdResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "user.DeleteUserByID")
	defer span.Finish()

	parsedUserID, err := uuid.Parse(r.GetId())
	if err != nil {
		u.logger.Errorf("ConvertUserID: %v", err)
		return nil, status.Errorf(codes.Internal, "ConvertUserID: %v", err)
	}

	_, err = u.userUC.FindByID(ctx, parsedUserID, u.cfg.Cache.Expire)
	if err != nil {
		u.logger.Errorf("userUC.FindByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "userUC.FindByID: %v", err)
	}

	// err = u.casbin.LoadPolicy(ctx)
	// if err != nil {
	// 	u.logger.Errorf("casbin.LoadPolicy: %v", err)
	// 	return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "casbin.LoadPolicy: %v", err)
	// }

	_, err = u.casbin.DeleteUser(ctx, parsedUserID.String())
	if err != nil {
		u.logger.Errorf("casbin.DeleteUser: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "casbin.DeleteUser: %v", err)
	}

	err = u.userUC.DeleteByID(ctx, parsedUserID)
	if err != nil {
		u.logger.Errorf("userUC.DeleteByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "userUC.DeleteByID: %v", err)
	}

	msg := &pb.UserKafka{
		Id: parsedUserID.String(),
	}
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "proto.Marshal: %v", err)
	}

	message := kafka.Message{
		Topic: "safety_delete_user_by_id",
		Value: msgBytes,
	}
	err = u.kafkaProducer.PublishMessage(ctx, message)
	if err != nil {
		u.logger.Errorf("SendDeleteUserKafkaByIdFailed: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "SendDeleteUserKafkaByIdFailed: %v", err)
	}

	err = u.SendHeader(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteUserByIdResponse{Message: "User deleted successfully"}, nil
}

// *Query

func (u *usersService) FindUsers(ctx context.Context, r *pb.FindUsersRequest) (*pb.FindUsersResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "user.FindUsers")
	defer span.Finish()

	// filter
	username := strings.TrimSpace(r.GetUsername())
	email := strings.ToLower(strings.TrimSpace(r.GetEmail()))
	role := strings.ToLower(r.GetRole())
	verified := strings.ToLower(r.GetVerified())

	paginateQuery := utils.NewPaginationQuery(r.GetLimit(), r.GetPage(), r.GetSort())

	filters := map[string]string{
		"username": username,
		"email":    email,
		"role":     role,
		"verified": verified,

		"limit": strconv.FormatUint(uint64(r.GetLimit()), 10),
		"page":  strconv.FormatUint(uint64(r.GetPage()), 10),
		"sort":  r.GetSort(),
	}

	userList, totalCount, err := u.userUC.Find(ctx, filters, paginateQuery, u.cfg.Cache.Expire)
	if err != nil {
		u.logger.Errorf("userUC.Find: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "userUC.Find: %v", err)
	}

	parsedUserList := make([]*pb.User, 0, totalCount)
	for _, user := range userList {
		parsedUserList = append(parsedUserList, u.UserModelToProto(user))
	}

	err = u.SendHeader(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.FindUsersResponse{
		TotalCount: totalCount,
		TotalPages: paginateQuery.GetTotalPages(totalCount),
		Page:       paginateQuery.GetPage(),
		Limit:      paginateQuery.GetLimit(),
		HasMore:    paginateQuery.GetHasMore(totalCount),
		Users:      parsedUserList,
	}, nil
}

func (u *usersService) FindUserById(ctx context.Context, r *pb.FindUserByIdRequest) (*pb.FindUserByIdResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "user.FindUserByID")
	defer span.Finish()

	parsedUserID, err := uuid.Parse(r.GetId())
	if err != nil {
		u.logger.Errorf("ConvertUserID: %v", err)
		return nil, status.Errorf(codes.Internal, "ConvertUserID: %v", err)
	}

	user, err := u.userUC.FindByID(ctx, parsedUserID, u.cfg.Cache.Expire)
	if err != nil {
		u.logger.Errorf("userUC.FindByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "userUC.FindByID: %v", err)
	}

	err = u.SendHeader(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.FindUserByIdResponse{User: u.UserModelToProto(user)}, nil
}

// TODO(Bagus): implement OTP

// TODO(Bagus): implement gomail gmail

// TODO(Bagus): convert Gorm to pure SQL or other Perfomance ORM

// TODO(Bagus): implement Docker

// TODO(Bagus): implement Docker Compose and convert it to Kompose
