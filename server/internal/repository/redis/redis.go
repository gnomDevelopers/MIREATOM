package redis

import (
	"fmt"
	"server/internal/config"

	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("localhost:%v", config.RedisPort),
		Password: config.RedisPassword,
		DB:       0,
	})

	return client
}
