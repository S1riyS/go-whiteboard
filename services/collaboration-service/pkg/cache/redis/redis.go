package redis

import (
	"context"
	"log/slog"
	"sync"

	"github.com/S1riyS/go-whiteboard/collaboration-service/internal/config"
	"github.com/S1riyS/go-whiteboard/collaboration-service/pkg/logger/slogext"
	"github.com/redis/go-redis/v9"
)

var (
	instance *redis.Client
	once     sync.Once
)

func MustNewClient(ctx context.Context, logger *slog.Logger, cfg config.RedisConfig) *redis.Client {
	once.Do(func() {
		const mark = "redis.MustNewClient"

		logger = logger.With(slog.String("mark", mark), slog.String("address", cfg.Address()))

		client := redis.NewClient(&redis.Options{
			Addr:     cfg.Address(),
			Password: cfg.Password,
			DB:       cfg.Database,
		})

		if err := client.Ping(ctx).Err(); err != nil {
			logger.Error("Failed to connect to redis", slogext.Err(err))
			panic(err)
		}

		instance = client
	})

	return instance
}
