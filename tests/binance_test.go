package tests

import (
	"context"
	"github.com/quantstop/qsx"
	"github.com/quantstop/qsx/exchange"
	"testing"
)

var binanceKey = ""
var binancePass = ""
var binanceSecret = ""

var BinanceClient exchange.IExchange

func TestBinanceClient(t *testing.T) {
	config := &exchange.Config{
		Auth:    exchange.NewAuth(binanceKey, binancePass, binanceSecret),
		Sandbox: true,
	}
	BinanceClient, err = qsx.NewExchange("binance", config)
	if err != nil {
		t.Error(err)
	}
}

func TestBinanceCandles(t *testing.T) {
	TestBinanceClient(t)
	candles, err := BinanceClient.GetHistoricalCandles(context.TODO(), "BTC-USD", "1m")
	if err != nil {
		t.Error(err)
	}
	for _, candle := range candles {
		t.Logf("Candle Time: %v | Open: %v | High: %v | Low: %v | Close: %v | Volume: %v", candle.Time, candle.Open, candle.High, candle.Low, candle.Close, candle.Volume)
	}
}
