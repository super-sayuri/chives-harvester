package fetcher

import (
	"context"
	"sayuri_crypto_bot/conf"
	"sayuri_crypto_bot/db"
	"sayuri_crypto_bot/model"
)

const maxItemNo = 100

func SaveItem(ctx context.Context, item *model.MarketValue) {
	log := conf.GetLog(ctx)
	db.SaveItemRecords(ctx, item.ID, item.Price)
	log.Info("success save item: " + item.ID + " with price: " + item.Price.String())
}

func DeleteMoreItems(ctx context.Context, id string) {
	log := conf.GetLog(ctx)
	err := db.DeleteMoreRecords(ctx, id, maxItemNo)
	if err != nil {
		log.Error("error when deleting more items: " + err.Error())
	}
}
