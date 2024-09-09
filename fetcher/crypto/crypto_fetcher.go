package crypto

import (
	"errors"
	"sayuri_crypto_bot/conf"
	"sayuri_crypto_bot/model"
)

type Fetcher interface {
	GetValue(items []*model.GoodsItem) ([]*model.MarketValue, error)
	GetName() string
}

var _cryptoFetcher Fetcher

var cyrptoInitMap map[string]func(conf *conf.CryptoConfig) Fetcher

func initCryptoInitMap() {
	cyrptoInitMap["gecko"] = NewGeckoFetcher
	cyrptoInitMap["cmc"] = NewCmcFetcher
}

func InitCryptoFetcher(conf *conf.CryptoConfig) error {
	initCryptoInitMap()
	f, ok := cyrptoInitMap[conf.App]
	if !ok {
		return errors.New("no Crypto Fetcher found")
	}
	_cryptoFetcher = f(conf)
	return nil
}

func GetCryptoFetcher() Fetcher {
	return _cryptoFetcher
}
