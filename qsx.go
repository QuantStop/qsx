package qsx

import (
	"errors"
	"fmt"
	"github.com/quantstop/qsx/binance"
	"github.com/quantstop/qsx/coinbasepro"
	"github.com/quantstop/qsx/core"
	"github.com/quantstop/qsx/yfinance"
)

func NewExchange(name core.ExchangeName, auth *core.Auth) (core.Qsx, error) {

	found := false
	for _, x := range core.SupportedExchanges {
		if x == name {
			found = true
			break
		}
	}
	if !found {
		return nil, errors.New(fmt.Sprintf("qsx error: failed to create exchange, '%s' is not supported", name))
	}

	switch name {
	case core.CoinbasePro:
		c, err := coinbasepro.NewCoinbasePro(auth)
		if err != nil {
			return nil, err
		}
		return c, nil

	case core.Binance:
		b, err := binance.NewBinance(auth)
		if err != nil {
			return nil, err
		}
		return b, nil

	case core.YFinance:
		b, err := yfinance.NewYFinance(auth)
		if err != nil {
			return nil, err
		}
		return b, nil

	default:
		return nil, errors.New("qsx error: failed to create exchange, unexpected error")
	}

}
