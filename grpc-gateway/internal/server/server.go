package server

import (
	"context"
	"gateway/config"
	"gateway/pkg/limiter"
	"gateway/pkg/logger"
	"gateway/pkg/token"

	"github.com/casbin/casbin-go-client/client"
	"github.com/minio/minio-go/v7"
	"github.com/rs/cors"
)

// GRPC Gateway Server
type Server struct {
	ctx    context.Context
	logger logger.Logger
	cfg    *config.Config
	// metrics       metric.Metrics
	tokenMaker *token.PasetoMaker
	minio      *minio.Client
	casbin     *client.Enforcer
	limiter    *limiter.RateLimit
	cors       *cors.Cors
}

// Server constructor
func NewGatewayServer(
	ctx context.Context,
	logger logger.Logger,
	cfg *config.Config,
	// metrics metric.Metrics,
	tokenMaker *token.PasetoMaker,
	minio *minio.Client,
	casbin *client.Enforcer,
	limiter *limiter.RateLimit,
	cors *cors.Cors,
) *Server {
	return &Server{
		ctx:    ctx,
		logger: logger,
		cfg:    cfg,
		// metrics:       metrics,
		tokenMaker: tokenMaker,
		minio:      minio,
		casbin:     casbin,
		limiter:    limiter,
		cors:       cors,
	}
}
