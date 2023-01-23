package domain

import (
	"context"
	"time"
	"user/pb"
)

// Policy Redis Repository
type PolicyRedisRepository interface {
	CreatePolicy(ctx context.Context, key string, value string, dataBytes []byte, expire time.Duration) error
	DeleteByID(ctx context.Context, key string, value string) error

	FindByID(ctx context.Context, key string, value string) ([]byte, error)
}

// Policy Usecase
type PolicyUseCase interface {
	CreatePolicy(ctx context.Context, policy *pb.Policy) (bool, error)
	DeletePolicy(ctx context.Context, policy *pb.Policy) (bool, error)

	Find(ctx context.Context, filters map[string]string, expire time.Duration) ([][]string, uint32, error)
}
