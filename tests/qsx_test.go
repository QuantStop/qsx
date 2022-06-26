package tests

import (
	"github.com/quantstop/qsx"
	"github.com/quantstop/qsx/core"
	"testing"
)

var key = ""
var pass = ""
var secret = ""

func TestNewClient(t *testing.T) {
	config := &core.Config{
		Auth:    core.NewAuth(key, pass, secret),
		Sandbox: true,
	}

	for _, x := range core.SupportedExchanges {
		exchange, err := qsx.NewExchange(x, config)
		if err != nil {
			t.Error(err)
		}
		t.Logf("Exchange Name: %v", exchange.GetName())
	}

}
