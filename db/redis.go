package db

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/shopspring/decimal"
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

func SaveItemRecords(ctx context.Context, name string, value decimal.Decimal) {
	key := fmt.Sprintf(DB_KEY_RECORD_FORMAT, name)
	_redis.LPush(ctx, key, value.String())
}

func DeleteMoreRecords(ctx context.Context, id string, threhold int) error {
	key := fmt.Sprintf(DB_KEY_RECORD_FORMAT, id)
	records, err := _redis.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return err
	}
	l := len(records)
	for l > threhold {
		_, err = _redis.RPop(ctx, key).Result()
		if err != nil {
			return err
		}
		l--
	}
	return nil
}

func SaveTgbotCommands(ctx context.Context, cs string) error {
	return _redis.Set(ctx, REDIS_KEY_TG_COMMAND, cs,
		time.Minute*time.Duration(conf.GetConfig().Tgbot.CallingGap)).Err()
}

func GetTgbotCommands(ctx context.Context) (string, error) {
	return _redis.Get(ctx, REDIS_KEY_TG_COMMAND).Result()
}
