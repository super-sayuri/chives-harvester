package db

import "sayuri_crypto_bot/conf"

const (
	DB_KEY_GROUPS        = "syr_groups_to_send"
	DB_KEY_CRYPTO_ITEMS  = "syr_crypto_items"
	DB_KEY_RECORD_FORMAT = "syr_avg_%s_records"

	REDIS_KEY_TG_COMMAND = "syt_tg_commands"
)

func Init(config *conf.Config) error {
	err := redisInit(config.Redis)
	if err != nil {
		return err
	}
	err = gormInit(config.Db)
	if err != nil {
		return err
	}
	return nil
}
