package job

import (
	"github.com/sirupsen/logrus"
	"sayuri_crypto_bot/db"
	"sayuri_crypto_bot/fetcher"
	"sayuri_crypto_bot/model"
	"sayuri_crypto_bot/sender"
	"sayuri_crypto_bot/util"
	"time"
)

func CryptoPrice(id string) error {
	log := logrus.WithField("jobId", id)
	cryotpItems, err := db.GetCryptoItems()
	if err != nil {
		return err
	}
	markets, err := fetcher.GeckoGetUsdValue(cryotpItems)
	if err != nil {
		return err
	}
	output := model.Output{
		Datetime: time.Now().Format("2006-01-02 15:04:05"),
		Items:    markets,
	}
	msg, err := util.TemplateGetString(util.CRYPTO_TEMPLATE, output)
	if err != nil {
		return err
	}
	groupIds, err := db.GetGroupIds()
	if err != nil {
		return err
	}
	for _, groupId := range groupIds {
		go func(i int64) {
			err := sender.TgSendData(i, msg)
			if err != nil {
				log.Error("error when sending telegram message: ", err)
			}
		}(groupId)
	}
	return nil
}
