package yfinance

import (
	"context"
	"github.com/quantstop/qsx/exchange"
	"github.com/quantstop/qsx/exchange/orderbook"
	"golang.org/x/time/rate"
	"net/http"
	"sync"
	"time"
)

type YFinance struct {
	exchange.Exchange
}

func NewYFinance(config *exchange.Config) (exchange.IExchange, error) {

	rl := rate.NewLimiter(rate.Every(time.Second), 10) // 10 requests per second

	api := exchange.New(
		&http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       0,
		},
		exchange.Options{
			ApiURL:  "",
			Verbose: false,
		},
		rl,
	)

	return &YFinance{
		exchange.Exchange{
			Name: exchange.YFinance,
			Auth: config.Auth,
			API:  api,
		},
	}, nil
}

func (y *YFinance) GetHistoricalCandles(ctx context.Context, productID string, granularity string) ([]exchange.Candle, error) {
	return nil, nil
}

func (y *YFinance) WatchFeed(shutdown chan struct{}, wg *sync.WaitGroup, product string, feed interface{}) (*orderbook.Orderbook, error) {
	return nil, nil
}

func (y *YFinance) ListProducts(ctx context.Context) ([]exchange.Product, error) {
	return nil, nil
}
