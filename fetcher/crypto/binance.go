package crypto

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"io"
	"net/http"
	"net/url"
	"sayuri_crypto_bot/client"
	"sayuri_crypto_bot/conf"
	"sayuri_crypto_bot/model"
	"strings"
)

/*
https://binance-docs.github.io/apidocs
*/
const binanceUrl = "https://api-gcp.binance.com/api/v3/ticker/24hr"

type BinanceFetcher struct {
	currency string
	client   *http.Client
}

type BinanceRespObj struct {
	Symbol             string `json:"symbol,omitempty"`
	PriceChange        string `json:"priceChange,omitempty"`
	PriceChangePercent string `json:"priceChangePercent,omitempty"`
	LastPrice          string `json:"lastPrice,omitempty"`
	HighPrice          string `json:"highPrice,omitempty"`
	LowPrice           string `json:"lowPrice,omitempty"`
}

type BinanceResp []*BinanceRespObj

func (f *BinanceFetcher) GetValue(items []*model.GoodsItem) ([]*model.MarketValue, error) {
	req, err := http.NewRequest("GET", binanceUrl, nil)
	if err != nil {
		return nil, err
	}

	symbols := make([]string, len(items))
	for i, item := range items {
		symbols[i] = fmt.Sprintf("\"%s%s\"", item.Id, f.currency)
	}

	req.Header.Set("Accept", "application/json")

	q := url.Values{}
	q.Add("symbols", "["+strings.Join(symbols, ",")+"]")
	req.URL.RawQuery = q.Encode()

	resp, err := f.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("binance fetcher returned status code %d", resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	binanceResp := &BinanceResp{}
	err = json.Unmarshal(respBody, binanceResp)
	if err != nil {
		return nil, err
	}

	marketValues := make([]*model.MarketValue, len(items))
	for i, binanceObj := range *binanceResp {
		price, _ := decimal.NewFromString(binanceObj.LastPrice)
		changeValue, _ := decimal.NewFromString(binanceObj.PriceChange)
		changePercent, _ := decimal.NewFromString(binanceObj.PriceChangePercent)
		High24h, _ := decimal.NewFromString(binanceObj.HighPrice)
		Low24h, _ := decimal.NewFromString(binanceObj.LowPrice)
		marketValues[i] = &model.MarketValue{
			ID:            items[i].Id,
			Name:          items[i].Id,
			Price:         price,
			Currency:      f.currency,
			ChangeVal:     changeValue,
			ChangePercent: changePercent,
			High24h:       High24h,
			Low24h:        Low24h,
		}
	}
	return marketValues, nil
}

func (f *BinanceFetcher) GetName() string {
	return "Binance"
}

var _ Fetcher = (*BinanceFetcher)(nil)

var _binanceFetcher *BinanceFetcher

func NewBinanceFetcher(conf *conf.CryptoConfig) Fetcher {
	_binanceFetcher = &BinanceFetcher{
		currency: strings.ToUpper(conf.Currency),
		client:   client.GetDefaultHttpClient(),
	}
	return _binanceFetcher
}
