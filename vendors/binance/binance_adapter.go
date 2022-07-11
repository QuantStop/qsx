package binance

import (
	"context"
	"github.com/quantstop/qsx/exchange"
	"github.com/quantstop/qsx/exchange/orderbook"
	"sync"
)

func (b *Binance) GetHistoricalCandles(ctx context.Context, productID string, granularity string) ([]exchange.Candle, error) {
	return nil, nil
}

func (b *Binance) WatchFeed(shutdown chan struct{}, wg *sync.WaitGroup, product string, feed interface{}) (*orderbook.Orderbook, error) {
	return nil, nil
}

func (b *Binance) ListProducts(ctx context.Context) ([]exchange.Product, error) {
	return nil, nil
}
