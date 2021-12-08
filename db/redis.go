package db

import (
	"context"
	"github.com/go-redis/redis/v8"
	"sayuri_crypto_bot/conf"
	"strings"
)

func Init(config *conf.Config) error {
	err := redisInit(config.Redis)
	if err != nil {
		return err
	}
	return nil
}

var (
	_redis *redis.ClusterClient
)

func redisInit(config *conf.RedisConfig) error {

	_redis = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    strings.Split(config.Nodes, ","),
		Username: config.Username,
		Password: config.Password,
	})
	_, err := _redis.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return nil
}

func GetRedisDb() *redis.ClusterClient {
	return _redis
}
