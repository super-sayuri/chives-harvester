package db

import "sayuri_crypto_bot/conf"

const (
	DB_KEY_GROUPS        = "syr_groups_to_send"
	DB_KEY_CRYPTO_ITEMS  = "syr_crypto_items"
	DB_KEY_RECORD_FORMAT = "syr_avg_%s_records"

	DB_KEY_API_CONFIG = "syr_api_conf"
)

func Init(config *conf.Config) error {
	err := redisInit(config.Redis)
	if err != nil {
		return err
	}
	return nil
}
