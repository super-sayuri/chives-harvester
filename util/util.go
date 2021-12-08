package util

import "sayuri_crypto_bot/conf"

func Init(config *conf.Config) error {
	var err error
	err = templateInit(config)
	if err != nil {
		return err
	}
	return nil
}
