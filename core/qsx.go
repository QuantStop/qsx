package core

import (
	"context"
	"github.com/quantstop/qsx/core/orderbook"
	"sync"
)

type Qsx interface {
	GetName() ExchangeName
	IsCrypto() bool

	ListProducts(ctx context.Context) ([]Product, error)
	GetHistoricalCandles(ctx context.Context, productID string, granularity string) ([]Candle, error)

	WatchFeed(shutdown chan struct{}, wg *sync.WaitGroup, product string, feed interface{}) (*orderbook.Orderbook, error)
}
