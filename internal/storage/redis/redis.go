package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/stepan41k/MyRest/internal/config"
)

type RedisStorage struct {
	client *redis.Client
}

func New(ctx context.Context, cfg config.RedisDB) (*redis.Client, error) {
	const op = "storage.redis.New"

	db := redis.NewClient(&redis.Options{
		Addr: cfg.Address,
		Password: cfg.Password,
		DB: cfg.DB,
		Username: cfg.User,
		MaxRetries: cfg.MaxRetries,
		DialTimeout: cfg.DialTimeout,
		ReadTimeout: cfg.Timeout,
		WriteTimeout: cfg.Timeout,
	})

	if err := db.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return db, nil
}