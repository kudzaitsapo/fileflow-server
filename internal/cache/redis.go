package cache

import (
	"github.com/kudzaitsapo/fileflow-server/internal/config"
	"github.com/redis/go-redis/v9"
)

func Initialise(cfg config.RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Host,
		Password: cfg.Password,
		DB:       cfg.Database,
	})

	return client
}