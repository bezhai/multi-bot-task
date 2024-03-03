package redis

import (
	"context"

	"github.com/redis/go-redis/v9"

	"github.com/bezhai/multi-bot-task/utils/env_utils"
)

var Redis *redis.Client

func Init() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     env_utils.Value("REDIS_SERVER_IP") + ":6379",
		Password: env_utils.Value("REDIS_PASSWORD"),
		DB:       0, // use default DB
	})
}

func MustGet(ctx context.Context, key string) string {
	result, err := Redis.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return result
}

func MustGetSlice(ctx context.Context, key string) []string {
	result, err := Redis.SMembers(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return result
}
