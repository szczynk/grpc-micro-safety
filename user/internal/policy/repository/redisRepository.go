package repository

import (
	"context"
	"fmt"
	"time"
	"user/internal/domain"
	"user/pkg/grpc_errors"

	"github.com/go-redis/redis/v9"
	"github.com/opentracing/opentracing-go"
)

// Policy Redis Repository
type policyRedisRepo struct {
	redisClient *redis.Client
}

// Policy Redis repository constructor
func NewPolicyRedisRepository(redisClient *redis.Client) domain.PolicyRedisRepository {
	return &policyRedisRepo{redisClient: redisClient}
}

// *Command

// Cache policy with duration in seconds
func (r *policyRedisRepo) CreatePolicy(ctx context.Context, key string, value string, dataBytes []byte, expire time.Duration) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PolicyRedisRepo.CreatePolicy")
	defer span.Finish()

	dataKey := r.createKey(key, value)

	return r.redisClient.Set(ctx, dataKey, dataBytes, expire).Err()
}

// Delete policy by key
func (r *policyRedisRepo) DeleteByID(ctx context.Context, key string, value string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PolicyRedisRepo.DeleteByID")
	defer span.Finish()

	return r.redisClient.Del(ctx, r.createKey(key, value)).Err()
}

// *Query

func (r *policyRedisRepo) FindByID(ctx context.Context, key string, value string) ([]byte, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PolicyRedisRepo.FindByID")
	defer span.Finish()

	dataBytes, err := r.redisClient.Get(ctx, r.createKey(key, value)).Bytes()
	if err != nil {
		if err != redis.Nil {
			return nil, grpc_errors.ErrNotFound
		}
		return nil, err
	}

	return dataBytes, nil
}

func (r *policyRedisRepo) createKey(key, value string) string {
	return fmt.Sprintf("%v: %v", key, value)
}
