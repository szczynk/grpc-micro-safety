package healthz

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type HealthzClient struct {
}

func NewHealthzClient(ctx context.Context, endpoint string) (grpc_health_v1.HealthClient, error) {
	conn, err := grpc.DialContext(ctx, endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	c := grpc_health_v1.NewHealthClient(conn)
	return c, nil
}
