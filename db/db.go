package db

import "sayuri_crypto_bot/conf"

const (
	DB_KEY_GROUPS       = "syr_groups_to_send"
	DB_KEY_CRYPTO_ITEMS = "syr_crypto_items"
)

func Init(config *conf.Config) error {
	err := redisInit(config.Redis)
	if err != nil {
		return err
	}
	return nil
}
