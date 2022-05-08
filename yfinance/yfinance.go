package yfinance

import (
	"context"
	"github.com/quantstop/qsx/core"
	"golang.org/x/time/rate"
	"net/http"
	"sync"
	"time"
)

type YFinance struct {
	core.Exchange
}

func NewYFinance(config *core.Config) (core.Qsx, error) {

	rl := rate.NewLimiter(rate.Every(time.Second), 10) // 10 requests per second

	api := core.New(
		&http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       0,
		},
		core.Options{
			ApiURL:  "",
			Verbose: false,
		},
		rl,
	)

	return &YFinance{
		core.Exchange{
			Name: core.YFinance,
			Auth: config.Auth,
			API:  api,
		},
	}, nil
}

func (y *YFinance) GetHistoricalCandles(ctx context.Context, productID string, granularity string) ([]core.Candle, error) {
	return nil, nil
}

func (y *YFinance) WatchFeed(shutdown chan struct{}, wg *sync.WaitGroup, product string, feed interface{}) error {
	return nil
}

func (y *YFinance) ListProducts(ctx context.Context) ([]core.Product, error) {
	return nil, nil
}
