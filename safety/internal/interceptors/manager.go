package interceptors

import (
	"context"

	"github.com/casbin/casbin-go-client/client"
	"google.golang.org/grpc"

	"safety/config"
	"safety/pkg/limiter"
	"safety/pkg/token"
)

type Interceptor interface {
	AuthUnary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error)
	AuthStream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error
	RateLimitUnary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error)
	RateLimitStream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error
}

// InterceptorManager
type interceptorManager struct {
	cfg        *config.Config
	tokenMaker *token.PasetoMaker
	casbin     *client.Enforcer
	limiter    *limiter.RateLimit
}

// InterceptorManager constructor
func NewInterceptorManager(
	cfg *config.Config,
	tokenMaker *token.PasetoMaker,
	casbin *client.Enforcer,
	limiter *limiter.RateLimit,
) Interceptor {
	return &interceptorManager{
		cfg:        cfg,
		tokenMaker: tokenMaker,
		casbin:     casbin,
		limiter:    limiter,
	}
}
