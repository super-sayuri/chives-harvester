package job

import (
	"context"
	"sayuri_crypto_bot/conf"
	"sayuri_crypto_bot/db"
	"sayuri_crypto_bot/fetcher"
	"sayuri_crypto_bot/model"
	"sayuri_crypto_bot/sender"
	"sayuri_crypto_bot/template"
	"time"
)

func CryptoPrice(ctx context.Context) error {
	log := conf.GetLog(ctx)
	cryotpItems, err := db.GetCryptoItems(ctx)
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
	msg, err := template.TemplateGetString(template.TEMPLATE_CRYPTO, output)
	if err != nil {
		return err
	}
	groupIds, err := db.GetGroupIds(ctx)
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
