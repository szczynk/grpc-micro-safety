package limiter

import (
	"auth/config"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/ulule/limiter/v3"
	redisStore "github.com/ulule/limiter/v3/drivers/store/redis"
)

type RateLimit struct {
	limiter    *limiter.Limiter
	redisStore *redis.Client
}

func NewRateLimit(cfg *config.Config) (*RateLimit, error) {
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

	// global limit per server is 50 reqs/second
	rate, err := limiter.NewRateFromFormatted(cfg.Limiter.Limit)
	if err != nil {
		return &RateLimit{redisStore: client}, err
	}

	store, err := redisStore.NewStore(client)
	if err != nil {
		return &RateLimit{redisStore: client}, err
	}

	instance := limiter.New(store, rate)
	return &RateLimit{limiter: instance, redisStore: client}, nil
}

func (r *RateLimit) Limiter() *limiter.Limiter {
	return r.limiter
}

func (r *RateLimit) Store() *redis.Client {
	return r.redisStore
}
