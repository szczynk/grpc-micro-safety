package repository

import (
	"auth/internal/domain"
	"auth/pkg/grpc_errors"
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/opentracing/opentracing-go"
)

// User Redis Repository
type userRedisRepo struct {
	redisClient *redis.Client
}

// User Redis repository constructor
func NewUserRedisRepository(redisClient *redis.Client) domain.UserRedisRepository {
	return &userRedisRepo{redisClient: redisClient}
}

// *Command

// Cache user with duration in seconds
func (r *userRedisRepo) CreateUser(ctx context.Context, key string, value string, dataBytes []byte, expire time.Duration) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserRedisRepo.CreateUser")
	defer span.Finish()

	dataKey := r.createKey(key, value)

	return r.redisClient.Set(ctx, dataKey, dataBytes, expire).Err()
}

// Delete user by key
func (r *userRedisRepo) DeleteByID(ctx context.Context, key string, value string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserRedisRepo.DeleteByID")
	defer span.Finish()

	return r.redisClient.Del(ctx, r.createKey(key, value)).Err()
}

// *Query

func (r *userRedisRepo) FindByID(ctx context.Context, key string, value string) ([]byte, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserRedisRepo.FindByID")
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

func (r *userRedisRepo) createKey(key, value string) string {
	return fmt.Sprintf("%v: %v", key, value)
}
