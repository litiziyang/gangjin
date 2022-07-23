package redis

import (
	"comm/tool"
	"github.com/go-redis/redis/v8"
)

func GetRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     tool.GetEnvDefault("REDIS_HOST", "redis:6379"),
		Password: tool.GetEnvDefault("REDIS_PASSWORD", ""),
		DB:       2,
	})
}
