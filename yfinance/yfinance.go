package yfinance

import (
	"context"
	"github.com/quantstop/qsx/qsx"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

type YFinance struct {
	qsx.Exchange
}

func NewYFinance(auth *qsx.Auth) (qsx.Qsx, error) {

	rl := rate.NewLimiter(rate.Every(time.Second), 10) // 10 requests per second

	api := qsx.New(
		&http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       0,
		},
		qsx.Options{
			ApiURL:  "",
			Verbose: false,
		},
		rl,
	)

	return &YFinance{
		qsx.Exchange{
			Name: qsx.YFinance,
			Auth: auth,
			API:  api,
		},
	}, nil
}

func (y *YFinance) GetHistoricalCandles(ctx context.Context, productID string) ([]qsx.Candle, error) {
	return nil, nil
}
