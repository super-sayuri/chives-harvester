package crypto

import (
	"encoding/json"
	"errors"
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

/**
  https://coinmarketcap.com/
*/

const cmcUrl = "https://pro-api.coinmarketcap.com/v2/cryptocurrency/quotes/latest"

type CmcFetcher struct {
	token    string
	currency string
	client   *http.Client
}

func (f *CmcFetcher) GetValue(items []*model.GoodsItem) ([]*model.MarketValue, error) {
	req, err := http.NewRequest("GET", cmcUrl, nil)
	if err != nil {
		return nil, err
	}

	ids := ""
	for i, item := range items {
		if i != 0 {
			ids = ids + ","
		}
		ids = ids + item.Alias["cmc"]
	}

	q := url.Values{}
	q.Add("convert", f.currency)
	q.Add("id", ids)

	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-CMC_PRO_API_KEY", f.token)

	req.URL.RawQuery = q.Encode()

	resp, err := f.client.Do(req)
	if err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	respObj := &cmcResp{}
	err = json.Unmarshal(respBody, respObj)
	if err != nil {
		return nil, err
	}
	if respObj.Status.ErrorCode != 0 {
		return nil, errors.New(fmt.Sprintf("%d:%s", respObj.Status.ErrorCode, respObj.Status.ErrorMsg))
	}

	marketvalues := make([]*model.MarketValue, 0)
	for _, item := range respObj.Data {
		quote := item.Quote[strings.ToUpper(f.currency)]
		marketValue := &model.MarketValue{
			ID:            item.Symbol,
			Name:          item.Name,
			Price:         decimal.NewFromFloat(quote.Price),
			Currency:      f.currency,
			ChangePercent: decimal.NewFromFloat(quote.PercentChange24h),
			High24h:       decimal.Decimal{},
			Low24h:        decimal.Decimal{},
		}
		marketValue.ChangeVal = marketValue.Price.Mul(marketValue.ChangePercent)
		marketvalues = append(marketvalues, marketValue)
	}
	return marketvalues, nil
}

func (f *CmcFetcher) GetName() string {
	return "coinmarketcap"
}

var _ Fetcher = (*CmcFetcher)(nil)

var _cmcFetcher *CmcFetcher

func NewCmcFetcher(conf *conf.CryptoConfig) Fetcher {
	_cmcFetcher = &CmcFetcher{
		token:    conf.Token,
		currency: conf.Currency,
		client:   client.GetDefaultHttpClient(),
	}
	return _cmcFetcher
}

type cmcResp struct {
	Status *cmcStatus            `json:"status"`
	Data   map[string]*cmcObject `json:"data"`
}

type cmcStatus struct {
	Timestamp   string `json:"timestamp,omitempty"`
	ErrorCode   int    `json:"error_code,omitempty"`
	ErrorMsg    string `json:"error_message,omitempty"`
	Elapsed     int    `json:"elapsed,omitempty"`
	CreditCount int    `json:"credit_count,omitempty"`
	NoticeCount int    `json:"notice_count,omitempty"`
}

type cmcObject struct {
	Id     int                  `json:"id,omitempty"`
	Name   string               `json:"name,omitempty"`
	Symbol string               `json:"symbol,omitempty"`
	Quote  map[string]*cmcQuote `json:"quote,omitempty"`
}

type cmcQuote struct {
	Price            float64 `json:"price,omitempty"`
	PercentChange24h float64 `json:"percent_change_24h,omitempty"`
}
