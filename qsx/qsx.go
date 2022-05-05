package qsx

import "context"

type Qsx interface {
	GetName() ExchangeName
	IsCrypto() bool

	GetHistoricalCandles(ctx context.Context, productID string) ([]Candle, error)
}
