package db

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sayuri_crypto_bot/conf"
	"time"
)

var (
	_redis *redis.Client
)

func redisInit(config *conf.RedisConfig) error {

	_redis = redis.NewClient(&redis.Options{
		Addr:     config.Nodes,
		Username: config.Username,
		Password: config.Password,
	})
	_, err := _redis.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return nil
}

func GetRedisDb() *redis.Client {
	return _redis
}

func CheckUserAvailable(ctx context.Context, id int64) bool {
	key := fmt.Sprintf("user_blacklist_%d", id)
	res, err := _redis.Get(ctx, key).Result()
	if err != nil || len(res) == 0 {
		return true
	}
	return false
}

func CheckChatAvailable(ctx context.Context, id int64) bool {
	key := fmt.Sprintf("chat_blacklist_%d", id)
	res, err := _redis.Get(ctx, key).Result()
	if err != nil || len(res) == 0 {
		return true
	}
	return false
}

func AddUserPeriod(ctx context.Context, id int64) {
	key := fmt.Sprintf("user_blacklist_%d", id)
	_redis.Set(ctx, key, true, time.Second*time.Duration(conf.GetConfig().Tgbot.CallingGap))
}

func AddChatPeriod(ctx context.Context, id int64) {
	key := fmt.Sprintf("chat_blacklist_%d", id)
	_redis.Set(ctx, key, true, time.Second*time.Duration(conf.GetConfig().Tgbot.CallingGap))
}
