package db

import (
	"context"
	"github.com/go-redis/redis/v8"
	"sayuri_crypto_bot/conf"
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
