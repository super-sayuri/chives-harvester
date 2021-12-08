package fetcher

import (
	"errors"
	"github.com/shopspring/decimal"
	gecko "github.com/superoo7/go-gecko/v3"
	"sayuri_crypto_bot/model"
)

func GeckoGetUsdValue(items []*model.GoodsItem) ([]*model.MarketValue, error) {
	return GeckoGetValue(items, "usd")
}

func GeckoGetValue(items []*model.GoodsItem, currency string) ([]*model.MarketValue, error) {
	ids := make([]string, 0)
	for _, item := range items {
		geckoId := item.GetAlias("gecko")
		if len(geckoId) != 0 {
			ids = append(ids, geckoId)
		}
	}
	if len(ids) == 0 {
		return nil, errors.New("no valid in input items")
	}
	cg := gecko.NewClient(nil)
	markets, err := cg.CoinsMarket(currency, ids, "", len(ids), 0, false, nil)
	if err != nil {
		return nil, err
	}
	res := make([]*model.MarketValue, 0)
	for _, market := range *markets {
		res = append(res, &model.MarketValue{
			Name:          market.Name,
			Price:         decimal.NewFromFloat(market.CurrentPrice),
			Currency:      currency,
			ChangeVal:     decimal.NewFromFloat(market.PriceChange24h),
			ChangePercent: decimal.NewFromFloat(market.PriceChangePercentage24h),
			High24h:       decimal.NewFromFloat(market.High24),
			Low24h:        decimal.NewFromFloat(market.Low24),
		})
	}
	return res, nil
}
