package qsx

import (
	"errors"
	"fmt"
	"github.com/quantstop/qsx/binance"
	"github.com/quantstop/qsx/coinbasepro"
	"github.com/quantstop/qsx/qsx"
	"github.com/quantstop/qsx/yfinance"
)

func NewExchange(name qsx.ExchangeName, auth *qsx.Auth) (qsx.Qsx, error) {

	found := false
	for _, x := range qsx.SupportedExchanges {
		if x == name {
			found = true
			break
		}
	}
	if !found {
		return nil, errors.New(fmt.Sprintf("qsx error: failed to create exchange, '%s' is not supported", name))
	}

	switch name {
	case qsx.CoinbasePro:
		c, err := coinbasepro.NewCoinbasePro(auth)
		if err != nil {
			return nil, err
		}
		return c, nil

	case qsx.Binance:
		b, err := binance.NewBinance(auth)
		if err != nil {
			return nil, err
		}
		return b, nil

	case qsx.YFinance:
		b, err := yfinance.NewYFinance(auth)
		if err != nil {
			return nil, err
		}
		return b, nil

	default:
		return nil, errors.New("qsx error: failed to create exchange, unexpected error")
	}

}
