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

// Workspace Redis Repository
type workspaceRedisRepo struct {
	redisClient *redis.Client
}

// Workspace Redis repository constructor
func NewWorkspaceRedisRepository(redisClient *redis.Client) domain.WorkspaceRedisRepository {
	return &workspaceRedisRepo{redisClient: redisClient}
}

// *Command

// Cache workspace with duration in seconds
func (r *workspaceRedisRepo) CreateWorkspace(ctx context.Context, key string, value string, dataBytes []byte, expire time.Duration) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "WorkspaceRedisRepo.CreateWorkspace")
	defer span.Finish()

	dataKey := r.createKey(key, value)

	return r.redisClient.Set(ctx, dataKey, dataBytes, expire).Err()
}

// Delete workspace by key
func (r *workspaceRedisRepo) DeleteByID(ctx context.Context, key string, value string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "WorkspaceRedisRepo.DeleteByID")
	defer span.Finish()

	return r.redisClient.Del(ctx, r.createKey(key, value)).Err()
}

// *Query

func (r *workspaceRedisRepo) FindByID(ctx context.Context, key string, value string) ([]byte, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "WorkspaceRedisRepo.FindByID")
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

func (r *workspaceRedisRepo) createKey(key, value string) string {
	return fmt.Sprintf("%v: %v", key, value)
}
