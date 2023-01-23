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

// Method Redis Repository
type methodRedisRepo struct {
	redisClient *redis.Client
}

// Method Redis repository constructor
func NewMethodRedisRepository(redisClient *redis.Client) domain.MethodRedisRepository {
	return &methodRedisRepo{redisClient: redisClient}
}

// *Command

// Cache method with duration in seconds
func (r *methodRedisRepo) CreateMethod(ctx context.Context, key string, value string, dataBytes []byte, expire time.Duration) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "MethodRedisRepo.CreateMethod")
	defer span.Finish()

	dataKey := r.createKey(key, value)

	return r.redisClient.Set(ctx, dataKey, dataBytes, expire).Err()
}

// Delete method by key
func (r *methodRedisRepo) DeleteByID(ctx context.Context, key string, value string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "MethodRedisRepo.DeleteByID")
	defer span.Finish()

	return r.redisClient.Del(ctx, r.createKey(key, value)).Err()
}

// *Query

func (r *methodRedisRepo) FindByID(ctx context.Context, key string, value string) ([]byte, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "MethodRedisRepo.FindByID")
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

func (r *methodRedisRepo) createKey(key, value string) string {
	return fmt.Sprintf("%v: %v", key, value)
}
