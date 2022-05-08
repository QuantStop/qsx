package binance

import (
	"context"
	"github.com/quantstop/qsx/qsx"
	"sync"
)

func (b *Binance) GetHistoricalCandles(ctx context.Context, productID string) ([]qsx.Candle, error) {
	return nil, nil
}

func (b *Binance) WatchFeed(shutdown chan struct{}, wg *sync.WaitGroup, product string, feed interface{}) error {
	return nil
}
