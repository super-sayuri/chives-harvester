package model

import "github.com/shopspring/decimal"

type MarketValue struct {
	ID            string
	Name          string
	Price         decimal.Decimal
	Currency      string
	ChangeVal     decimal.Decimal
	ChangePercent decimal.Decimal
	High24h       decimal.Decimal
	Low24h        decimal.Decimal
}
