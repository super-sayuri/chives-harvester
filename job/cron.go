package job

import (
	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"sayuri_crypto_bot/conf"
)

func CronInit() {
	c := cron.New()

	c.AddFunc(conf.GetConfig().Cron.Crypto, WrapFunc(CryptoPrice, "Crypto Market Price"))

	c.Start()
}

func WrapFunc(f func(id string) error, name string) func() {
	return func() {
		id := uuid.NewString()
		log.Info("Start Job '", name, "'. ID: #", id)
		defer log.Info("End Job #", id)
		err := f(id)
		if err != nil {
			log.Error("Job #", id, " throw an error: ", err)
		}
	}
}
