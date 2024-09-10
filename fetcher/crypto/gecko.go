package crypto

/**
  https://www.coingecko.com
*/
import (
	"sayuri_crypto_bot/conf"
	"sayuri_crypto_bot/model"
)

type GeckoFetcher struct {
	currency string
}

func (g *GeckoFetcher) GetValue(items []*model.GoodsItem) ([]*model.MarketValue, error) {
	return nil, nil
}

func (g *GeckoFetcher) GetName() string {
	return "gecko"
}

var _ Fetcher = (*GeckoFetcher)(nil)

var _geckoFetcher *GeckoFetcher

func NewGeckoFetcher(conf *conf.CryptoConfig) Fetcher {
	_geckoFetcher = &GeckoFetcher{
		currency: conf.Currency,
	}
	return _geckoFetcher
}
