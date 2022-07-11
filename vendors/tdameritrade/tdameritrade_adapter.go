package tdameritrade

import (
	"context"
	"github.com/quantstop/qsx/exchange"
	"github.com/quantstop/qsx/exchange/orderbook"
	"sync"
)

func (t *TDAmeritrade) GetHistoricalCandles(ctx context.Context, productID string, granularity string) ([]exchange.Candle, error) {
	return nil, nil
}

func (t *TDAmeritrade) WatchFeed(shutdown chan struct{}, wg *sync.WaitGroup, product string, feed interface{}) (*orderbook.Orderbook, error) {
	return nil, nil
}

func (t *TDAmeritrade) ListProducts(ctx context.Context) ([]exchange.Product, error) {
	return nil, nil
}
