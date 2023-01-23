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

// Service Redis Repository
type serviceRedisRepo struct {
	redisClient *redis.Client
}

// Service Redis repository constructor
func NewServiceRedisRepository(redisClient *redis.Client) domain.ServiceRedisRepository {
	return &serviceRedisRepo{redisClient: redisClient}
}

// *Command

// Cache service with duration in seconds
func (r *serviceRedisRepo) CreateService(ctx context.Context, key string, value string, dataBytes []byte, expire time.Duration) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ServiceRedisRepo.CreateService")
	defer span.Finish()

	dataKey := r.createKey(key, value)

	return r.redisClient.Set(ctx, dataKey, dataBytes, expire).Err()
}

// Delete service by key
func (r *serviceRedisRepo) DeleteByID(ctx context.Context, key string, value string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ServiceRedisRepo.DeleteByID")
	defer span.Finish()

	return r.redisClient.Del(ctx, r.createKey(key, value)).Err()
}

// *Query

func (r *serviceRedisRepo) FindByID(ctx context.Context, key string, value string) ([]byte, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ServiceRedisRepo.FindByID")
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

func (r *serviceRedisRepo) createKey(key, value string) string {
	return fmt.Sprintf("%v: %v", key, value)
}
