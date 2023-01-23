package casbin

import (
	"context"
	"fmt"
	"gateway/config"

	"github.com/casbin/casbin-go-client/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CasbinClient struct {
}

func NewCasbinClient(ctx context.Context, cfg *config.Config) (*client.Enforcer, error) {
	dataSourceName := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=casbin sslmode=%s TimeZone=Asia/Jakarta",
		cfg.Casbin.PostgresHost,
		cfg.Casbin.PostgresPort,
		cfg.Casbin.PostgresUser,
		cfg.Casbin.PostgresPassword,
		cfg.Casbin.PostgresSSLMode,
	)

	casbinConfig := client.Config{
		DriverName:    "postgres",
		ConnectString: dataSourceName,
		ModelText:     "",
		DbSpecified:   true,
	}

	casbinURL := cfg.Casbin.URL

	if casbinURL == "" {
		casbinURL = "localhost:50051"
	}

	cc, err := client.NewClient(ctx, casbinURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	e, err := cc.NewEnforcer(ctx, casbinConfig)
	if err != nil {
		return nil, err
	}

	return e, nil
}
