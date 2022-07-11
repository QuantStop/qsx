package qsx

import (
	"errors"
	"fmt"
	"github.com/quantstop/qsx/exchange"
	"github.com/quantstop/qsx/vendors/binance"
	"github.com/quantstop/qsx/vendors/coinbasepro"
	"github.com/quantstop/qsx/vendors/yfinance"
)

// NewExchange creates an exchange connection and returns a struct that implements the IExchange interface
func NewExchange(name exchange.Name, config *exchange.Config) (exchange.IExchange, error) {

	found := false
	for _, x := range exchange.SupportedExchanges {
		if x == name {
			found = true
			break
		}
	}
	if !found {
		return nil, errors.New(fmt.Sprintf("qsx error: failed to create exchange, '%s' is not supported", name))
	}

	switch name {
	case exchange.CoinbasePro:
		c, err := coinbasepro.NewCoinbasePro(config)
		if err != nil {
			return nil, err
		}
		return c, nil

	case exchange.Binance:
		b, err := binance.NewBinance(config)
		if err != nil {
			return nil, err
		}
		return b, nil

	case exchange.YFinance:
		b, err := yfinance.NewYFinance(config)
		if err != nil {
			return nil, err
		}
		return b, nil

	default:
		return nil, errors.New("qsx error: failed to create exchange, unexpected error")
	}

}

// GetSupportedExchanges returns a list of all the supported exchanges
func GetSupportedExchanges() []exchange.Name {
	return exchange.SupportedExchanges
}
