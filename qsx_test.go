package qsx

import (
	"context"
	"github.com/quantstop/qsx/qsx"
	"testing"
)

var key = ""
var pass = ""
var secret = ""

func TestNewClient(t *testing.T) {
	auth := qsx.NewAuth(key, pass, secret)

	for _, x := range qsx.SupportedExchanges {
		exchange, err := NewExchange(x, auth)
		if err != nil {
			t.Error(err)
		}
		t.Logf("Exchange Name: %v", exchange.GetName())
	}

}

func TestCoinbaseCandles(t *testing.T) {
	auth := qsx.NewAuth(key, pass, secret)
	coinbasepro, err := NewExchange("coinbasepro", auth)
	if err != nil {
		t.Error(err)
	}
	candles, err := coinbasepro.GetHistoricalCandles(context.TODO(), "BTC-USD")
	if err != nil {
		t.Error(err)
	}
	for _, candle := range candles {
		t.Logf("Candle Time: %v | Open: %v | High: %v | Low: %v | Close: %v | Volume: %v", candle.Time, candle.Open, candle.High, candle.Low, candle.Close, candle.Volume)
	}
}
