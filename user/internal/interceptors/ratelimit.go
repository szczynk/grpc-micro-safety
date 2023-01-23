package interceptors

import (
	"context"
	"path"
	"user/pkg/grpc_errors"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

func (im *interceptorManager) RateLimitUnary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	method := path.Base(info.FullMethod)

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, grpc_errors.ErrNoCtxMetaData
	}

	var clientIP string
	if mdClientIPs := md.Get("x-forwarded-for"); len(mdClientIPs) > 0 {
		clientIP = mdClientIPs[0]
	}

	if mdClientIPs := md.Get("x-real-ip"); len(mdClientIPs) > 0 {
		clientIP = mdClientIPs[0]
	}

	if p, ok := peer.FromContext(ctx); ok {
		clientIP = p.Addr.String()
	}

	if len(clientIP) == 0 {
		clientIP = "unknown"
	}

	limiterContext, err := im.limiter.Limiter().Get(ctx, clientIP)
	if err != nil {
		return nil, err
	}

	if limiterContext.Reached {
		return nil, status.Errorf(codes.ResourceExhausted, "The method of %s's limit has been exceeded. Please try again later.", method)
	}

	newCtx := metadata.NewIncomingContext(ctx, md)

	reply, err := handler(newCtx, req)
	return reply, err
}

func (im *interceptorManager) RateLimitStream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	ctx := stream.Context()

	method := path.Base(info.FullMethod)

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return grpc_errors.ErrNoCtxMetaData
	}

	var clientIP string
	if mdClientIPs := md.Get("x-forwarded-for"); len(mdClientIPs) > 0 {
		clientIP = mdClientIPs[0]
	}

	if mdClientIPs := md.Get("x-real-ip"); len(mdClientIPs) > 0 {
		clientIP = mdClientIPs[0]
	}

	if p, ok := peer.FromContext(ctx); ok {
		clientIP = p.Addr.String()
	}

	if len(clientIP) == 0 {
		clientIP = "unknown"
	}

	limiterContext, err := im.limiter.Limiter().Get(ctx, clientIP)
	if err != nil {
		return err
	}

	if limiterContext.Reached {
		return status.Errorf(codes.ResourceExhausted, "The method of %s's limit has been exceeded. Please try again later.", method)
	}

	newCtx := metadata.NewIncomingContext(ctx, md)

	wrapped := grpc_middleware.WrapServerStream(stream)
	wrapped.WrappedContext = newCtx

	err = handler(srv, wrapped)
	return err
}
