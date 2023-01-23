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

// Role Redis Repository
type roleRedisRepo struct {
	redisClient *redis.Client
}

// Role Redis repository constructor
func NewRoleRedisRepository(redisClient *redis.Client) domain.RoleRedisRepository {
	return &roleRedisRepo{redisClient: redisClient}
}

// *Command

// Cache role with duration in seconds
func (r *roleRedisRepo) CreateRole(ctx context.Context, key string, value string, dataBytes []byte, expire time.Duration) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RoleRedisRepo.CreateRole")
	defer span.Finish()

	dataKey := r.createKey(key, value)

	return r.redisClient.Set(ctx, dataKey, dataBytes, expire).Err()
}

// Delete role by key
func (r *roleRedisRepo) DeleteByID(ctx context.Context, key string, value string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RoleRedisRepo.DeleteByID")
	defer span.Finish()

	return r.redisClient.Del(ctx, r.createKey(key, value)).Err()
}

// *Query

func (r *roleRedisRepo) FindByID(ctx context.Context, key string, value string) ([]byte, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RoleRedisRepo.FindByID")
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

func (r *roleRedisRepo) createKey(key, value string) string {
	return fmt.Sprintf("%v: %v", key, value)
}
