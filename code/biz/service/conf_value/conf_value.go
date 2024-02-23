package conf_value

import (
	"context"

	"github.com/bezhai/multi-bot-task/biz/clients/redis_client"
)

func GetConfStringValue(ctx context.Context, key string) (string, error) {
	return redis_client.Redis.Get(ctx, key).Result()
}
