package crypto

/**
  https://www.coingecko.com
*/
import (
	"sayuri_crypto_bot/conf"
	"sayuri_crypto_bot/model"
	"sync"
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
	sync.OnceFunc(func() {
		_geckoFetcher = &GeckoFetcher{
			currency: conf.Currency,
		}
	})
	return _geckoFetcher
}
