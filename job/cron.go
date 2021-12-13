package job

import (
	"context"
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

func WrapFunc(f func(context.Context) error, name string) func() {
	return func() {
		id := uuid.NewString()
		log.Info("Start Job '", name, "'. ID: #", id)
		defer log.Info("End Job #", id)
		ctx := context.WithValue(context.Background(), conf.LOG_KEY_JOB_ID, id)
		err := f(ctx)
		if err != nil {
			log.Error("Job #", id, " throw an error: ", err)
		}
	}
}
