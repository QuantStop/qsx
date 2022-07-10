package tdameritrade

import (
	"context"
	"github.com/quantstop/qsx/core"
	"github.com/quantstop/qsx/core/orderbook"
	"sync"
)

func (t *TDAmeritrade) GetHistoricalCandles(ctx context.Context, productID string, granularity string) ([]core.Candle, error) {
	return nil, nil
}

func (t *TDAmeritrade) WatchFeed(shutdown chan struct{}, wg *sync.WaitGroup, product string, feed interface{}) (*orderbook.Orderbook, error) {
	return nil, nil
}

func (t *TDAmeritrade) ListProducts(ctx context.Context) ([]core.Product, error) {
	return nil, nil
}
