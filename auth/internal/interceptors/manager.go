package interceptors

import (
	"context"

	"google.golang.org/grpc"

	"auth/config"
	"auth/pkg/limiter"
)

type Interceptor interface {
	RateLimitUnary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error)
	RateLimitStream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error
}

// InterceptorManager
type interceptorManager struct {
	cfg     *config.Config
	limiter *limiter.RateLimit
}

// InterceptorManager constructor
func NewInterceptorManager(
	cfg *config.Config,
	limiter *limiter.RateLimit,
) Interceptor {
	return &interceptorManager{
		cfg:     cfg,
		limiter: limiter,
	}
}
