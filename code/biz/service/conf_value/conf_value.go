package conf_value

import (
	"context"

	"github.com/bezhai/multi-bot-task/biz/clients/redis_client"
	"github.com/bezhai/multi-bot-task/biz/utils/langx/slicex"
)

func GetConfStringValue(ctx context.Context, key string) (string, error) {
	return redis_client.Redis.Get(ctx, key).Result()
}

func SetConfStringValue(ctx context.Context, key string, value string) (string, error) {
	return redis_client.Redis.Set(ctx, key, value, 0).Result()
}

func GetConfMemberValue(ctx context.Context, key string) ([]string, error) {
	return redis_client.Redis.SMembers(ctx, key).Result()
}

func SetConfMemberValue(ctx context.Context, key string, value []string) error {
	res, err := redis_client.Redis.SMembers(ctx, key).Result()
	if err != nil {
		return err
	}
	addMembers := make([]interface{}, 0)
	removeMembers := make([]interface{}, 0)
	for _, s := range value {
		if !slicex.Contains(res, s) {
			addMembers = append(addMembers, s)
		}
	}
	for _, s := range res {
		if !slicex.Contains(value, s) {
			removeMembers = append(removeMembers, s)
		}
	}
	if len(addMembers) > 0 {
		err = redis_client.Redis.SAdd(ctx, key, addMembers...).Err()
		if err != nil {
			return err
		}
	}
	if len(removeMembers) > 0 {
		err = redis_client.Redis.SRem(ctx, key, removeMembers...).Err()
		if err != nil {
			return err
		}
	}
	return nil
}
