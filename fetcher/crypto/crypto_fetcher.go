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

var cryptoInitMap map[string]func(conf *conf.CryptoConfig) Fetcher

func initCryptoInitMap() {
	cryptoInitMap = make(map[string]func(conf *conf.CryptoConfig) Fetcher)
	cryptoInitMap["gecko"] = NewGeckoFetcher
	cryptoInitMap["cmc"] = NewCmcFetcher
	cryptoInitMap["binance"] = NewBinanceFetcher
}

func InitCryptoFetcher(conf *conf.CryptoConfig) error {
	initCryptoInitMap()
	f, ok := cryptoInitMap[conf.App]
	if !ok {
		return errors.New("no Crypto Fetcher found")
	}
	_cryptoFetcher = f(conf)
	return nil
}

func GetCryptoFetcher() Fetcher {
	return _cryptoFetcher
}
