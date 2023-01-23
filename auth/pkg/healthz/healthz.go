package healthz

import (
	"auth/config"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type HealthzClient struct {
}

func NewHealthzClient(ctx context.Context, cfg *config.Config) (grpc_health_v1.HealthClient, error) {
	endpoint := cfg.Server.Port

	conn, err := grpc.DialContext(ctx, endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	c := grpc_health_v1.NewHealthClient(conn)
	return c, nil
}
