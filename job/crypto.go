package job

import (
	"context"
	"sayuri_crypto_bot/conf"
	"sayuri_crypto_bot/db"
	"sayuri_crypto_bot/fetcher"
	cryptoFetcher "sayuri_crypto_bot/fetcher/crypto"
	"sayuri_crypto_bot/model"
	"sayuri_crypto_bot/sender"
	"sayuri_crypto_bot/template"
	"time"
)

func CryptoPrice(ctx context.Context) error {
	log := conf.GetLog(ctx)
	cryptoItems, err := db.GetCryptoItems(ctx)
	if err != nil {
		return err
	}
	markets, err := cryptoFetcher.GetCryptoFetcher().GetValue(cryptoItems)
	go saveMarkets(ctx, markets)
	if err != nil {
		return err
	}
	output := model.Output{
		Datetime: time.Now().Format("2006-01-02 15:04:05"),
		Items:    markets,
	}
	msg, err := template.GetString(template.Crypto, output)
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
			err = sender.TgSendSticker(i)
			if err != nil {
				log.Error("error when sending telegram stickers: ", err)
			}
		}(groupId)
	}
	return nil
}

func saveMarkets(ctx context.Context, markets []*model.MarketValue) {
	for _, market := range markets {
		go func(m *model.MarketValue) {
			fetcher.SaveItem(ctx, m)
			fetcher.DeleteMoreItems(ctx, m.ID)
		}(market)
	}
}
