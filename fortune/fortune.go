package fortune

import (
	"sayuri_crypto_bot/conf"
	"sayuri_crypto_bot/fortune/tarot"
)

func Init(config *conf.Config) error {
	err := tarot.Init(config)
	if err != nil {
		return err
	}
	return nil
}
