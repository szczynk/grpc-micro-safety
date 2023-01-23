package repository

import (
	"context"
	"fmt"
	"safety/internal/domain"
	"safety/pkg/grpc_errors"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/opentracing/opentracing-go"
)

// Office Redis Repository
type officeRedisRepo struct {
	redisClient *redis.Client
}

// Office Redis repository constructor
func NewOfficeRedisRepository(redisClient *redis.Client) domain.OfficeRedisRepository {
	return &officeRedisRepo{redisClient: redisClient}
}

// *Command

// Cache office with duration in seconds
func (r *officeRedisRepo) CreateOffice(ctx context.Context, key string, value string, dataBytes []byte, expire time.Duration) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OfficeRedisRepo.CreateOffice")
	defer span.Finish()

	dataKey := r.createKey(key, value)

	return r.redisClient.Set(ctx, dataKey, dataBytes, expire).Err()
}

// Delete office by key
func (r *officeRedisRepo) DeleteByID(ctx context.Context, key string, value string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OfficeRedisRepo.DeleteByID")
	defer span.Finish()

	return r.redisClient.Del(ctx, r.createKey(key, value)).Err()
}

// *Query

func (r *officeRedisRepo) FindByID(ctx context.Context, key string, value string) ([]byte, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "OfficeRedisRepo.FindByID")
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

func (r *officeRedisRepo) createKey(key, value string) string {
	return fmt.Sprintf("%v: %v", key, value)
}
