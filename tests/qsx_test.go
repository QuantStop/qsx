package tests

import (
	"github.com/quantstop/qsx"
	"github.com/quantstop/qsx/exchange"
	"testing"
)

var key = ""
var pass = ""
var secret = ""

func TestNewClient(t *testing.T) {
	config := &exchange.Config{
		Auth:    exchange.NewAuth(key, pass, secret),
		Sandbox: true,
	}

	for _, x := range exchange.SupportedExchanges {
		ex, err := qsx.NewExchange(x, config)
		if err != nil {
			t.Error(err)
		}
		t.Logf("Exchange Name: %v", ex.GetName())
	}
}

func TestSupportedExchanges(t *testing.T) {
	for i, x := range qsx.GetSupportedExchanges() {
		t.Logf("%v. %v\n", i+1, x)
	}
}
