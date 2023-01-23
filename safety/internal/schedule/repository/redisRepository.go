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

// Schedule Redis Repository
type scheduleRedisRepo struct {
	redisClient *redis.Client
}

// Schedule Redis repository constructor
func NewScheduleRedisRepository(redisClient *redis.Client) domain.ScheduleRedisRepository {
	return &scheduleRedisRepo{redisClient: redisClient}
}

// *Command

// Cache schedule with duration in seconds
func (r *scheduleRedisRepo) CreateSchedule(ctx context.Context, key string, value string, dataBytes []byte, expire time.Duration) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ScheduleRedisRepo.CreateSchedule")
	defer span.Finish()

	dataKey := r.createKey(key, value)

	return r.redisClient.Set(ctx, dataKey, dataBytes, expire).Err()
}

// Delete schedule by key
func (r *scheduleRedisRepo) DeleteByID(ctx context.Context, key string, value string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ScheduleRedisRepo.DeleteByID")
	defer span.Finish()

	return r.redisClient.Del(ctx, r.createKey(key, value)).Err()
}

// *Query

func (r *scheduleRedisRepo) FindByID(ctx context.Context, key string, value string) ([]byte, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ScheduleRedisRepo.FindByID")
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

func (r *scheduleRedisRepo) createKey(key, value string) string {
	return fmt.Sprintf("%v: %v", key, value)
}
