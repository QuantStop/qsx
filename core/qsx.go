package core

import (
	"context"
	"sync"
)

type Qsx interface {
	GetName() ExchangeName
	IsCrypto() bool

	GetHistoricalCandles(ctx context.Context, productID string) ([]Candle, error)

	WatchFeed(shutdown chan struct{}, wg *sync.WaitGroup, product string, feed interface{}) error
}
