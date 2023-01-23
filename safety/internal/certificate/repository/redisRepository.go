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

// Certificate Redis Repository
type certificateRedisRepo struct {
	redisClient *redis.Client
}

// Certificate Redis repository constructor
func NewCertificateRedisRepository(redisClient *redis.Client) domain.CertificateRedisRepository {
	return &certificateRedisRepo{redisClient: redisClient}
}

// *Command

// Cache certificate with duration in seconds
func (r *certificateRedisRepo) CreateCertificate(ctx context.Context, key string, value string, dataBytes []byte, expire time.Duration) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CertificateRedisRepo.CreateCertificate")
	defer span.Finish()

	dataKey := r.createKey(key, value)

	return r.redisClient.Set(ctx, dataKey, dataBytes, expire).Err()
}

// Delete certificate by key
func (r *certificateRedisRepo) DeleteByID(ctx context.Context, key string, value string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CertificateRedisRepo.DeleteByID")
	defer span.Finish()

	return r.redisClient.Del(ctx, r.createKey(key, value)).Err()
}

// *Query

func (r *certificateRedisRepo) FindByID(ctx context.Context, key string, value string) ([]byte, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CertificateRedisRepo.FindByID")
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

func (r *certificateRedisRepo) createKey(key, value string) string {
	return fmt.Sprintf("%v: %v", key, value)
}
