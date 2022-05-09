package binance

import (
	"context"
	"github.com/quantstop/qsx/core"
	"github.com/quantstop/qsx/core/orderbook"
	"sync"
)

func (b *Binance) GetHistoricalCandles(ctx context.Context, productID string, granularity string) ([]core.Candle, error) {
	return nil, nil
}

func (b *Binance) WatchFeed(shutdown chan struct{}, wg *sync.WaitGroup, product string, feed interface{}) (*orderbook.Orderbook, error) {
	return nil, nil
}

func (b *Binance) ListProducts(ctx context.Context) ([]core.Product, error) {
	return nil, nil
}
