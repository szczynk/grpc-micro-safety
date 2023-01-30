package interceptors

import (
	"context"
	"fmt"
	"path"
	"strings"
	"user/pkg/grpc_errors"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (im *interceptorManager) AuthUnary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	service := path.Dir(info.FullMethod)[1:]
	method := path.Base(info.FullMethod)

	var isExpire bool
	switch method {
	case "RefreshAccessToken":
		isExpire = false
	case "Logout":
		isExpire = false
	case "Check":
		return handler(ctx, req)
	case "ServerReflectionInfo":
		return handler(ctx, req)
	default:
		isExpire = true
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, grpc_errors.ErrNoCtxMetaData
	}

	accessToken, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	payload, err := im.tokenMaker.VerifyToken(accessToken, isExpire)
	if err != nil {
		return nil, fmt.Errorf("invalid access token: %s", err)
	}

	userId := payload.Get("user_id")
	role := payload.Get("role")
	md.Append("user_id", userId)
	md.Append("role", role)

	allowed, err := im.casbin.Enforce(ctx, role, service, method)
	if err != nil {
		return nil, fmt.Errorf("casbin.Enforce: %s", err)
	}
	if !allowed {
		return nil, status.Errorf(codes.Unauthenticated, grpc_errors.ErrMethodNotAllowed.Error())
	}

	newCtx := metadata.NewIncomingContext(ctx, md)

	reply, err := handler(newCtx, req)
	return reply, err
}

func (im *interceptorManager) AuthStream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	ctx := stream.Context()

	service := path.Dir(info.FullMethod)[1:]
	method := path.Base(info.FullMethod)

	var isExpire bool
	switch method {
	case "RefreshAccessToken":
		isExpire = false
	case "Logout":
		isExpire = false
	case "Watch":
		wrapped := grpc_middleware.WrapServerStream(stream)
		wrapped.WrappedContext = ctx
		return handler(srv, wrapped)
	case "ServerReflectionInfo":
		wrapped := grpc_middleware.WrapServerStream(stream)
		wrapped.WrappedContext = ctx
		return handler(srv, wrapped)
	default:
		isExpire = true
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return grpc_errors.ErrNoCtxMetaData
	}

	values := md.Get("authorization")
	if len(values) == 0 {
		return grpc_errors.ErrNoAccessToken
	}

	authHeader := values[0]
	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return grpc_errors.ErrInvalidAuthHeader
	}

	authHeaderType := strings.ToLower(fields[0])
	if authHeaderType != "bearer" {
		return grpc_errors.ErrInvalidAuthHeader
	}

	accessToken := fields[1]

	payload, err := im.tokenMaker.VerifyToken(accessToken, isExpire)
	if err != nil {
		return fmt.Errorf("invalid access token: %s", err)
	}

	userId := payload.Get("user_id")
	role := payload.Get("role")
	md.Append("user_id", userId)
	md.Append("role", role)

	allowed, err := im.casbin.Enforce(ctx, role, service, method)
	if err != nil {
		return fmt.Errorf("casbin.Enforce: %s", err)
	}
	if !allowed {
		return grpc_errors.ErrMethodNotAllowed
	}

	newCtx := metadata.NewIncomingContext(ctx, md)

	wrapped := grpc_middleware.WrapServerStream(stream)
	wrapped.WrappedContext = newCtx

	err = handler(srv, wrapped)
	return err
}
