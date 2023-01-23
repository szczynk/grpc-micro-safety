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

// Attendance Redis Repository
type attendanceRedisRepo struct {
	redisClient *redis.Client
}

// Attendance Redis repository constructor
func NewAttendanceRedisRepository(redisClient *redis.Client) domain.AttendanceRedisRepository {
	return &attendanceRedisRepo{redisClient: redisClient}
}

// *Command

// Cache attendance with duration in seconds
func (r *attendanceRedisRepo) CreateAttendance(ctx context.Context, key string, value string, dataBytes []byte, expire time.Duration) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "AttendanceRedisRepo.CreateAttendance")
	defer span.Finish()

	dataKey := r.createKey(key, value)

	return r.redisClient.Set(ctx, dataKey, dataBytes, expire).Err()
}

// Delete attendance by key
func (r *attendanceRedisRepo) DeleteByID(ctx context.Context, key string, value string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "AttendanceRedisRepo.DeleteByID")
	defer span.Finish()

	return r.redisClient.Del(ctx, r.createKey(key, value)).Err()
}

// *Query

func (r *attendanceRedisRepo) FindByID(ctx context.Context, key string, value string) ([]byte, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "AttendanceRedisRepo.FindByID")
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

func (r *attendanceRedisRepo) createKey(key, value string) string {
	return fmt.Sprintf("%v: %v", key, value)
}
