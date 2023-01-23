package redis

import (
	"time"

	"github.com/go-redis/redis/v9"

	"mail/config"
)

// Returns new redis client
func NewRedisClient(cfg *config.Config) *redis.Client {
	redisHost := cfg.Redis.Address

	if redisHost == "" {
		redisHost = ":6379"
	}

	client := redis.NewClient(
		&redis.Options{
			Addr:         redisHost,
			Password:     cfg.Redis.Password, // no password set
			DB:           cfg.Redis.DB,       // use default DB
			PoolSize:     cfg.Redis.PoolSize,
			PoolTimeout:  time.Duration(cfg.Redis.PoolTimeout) * time.Second,
			MinIdleConns: cfg.Redis.MinIdleConns,
		},
	)

	return client
}
