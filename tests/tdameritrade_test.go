package tests

import (
	"github.com/quantstop/qsx"
	"github.com/quantstop/qsx/exchange"
	"testing"
)

var tdaKey = ""
var tdaPass = ""
var tdaSecret = ""

var TDAClient exchange.IExchange
var tdaErr error

func TestTDAClient(t *testing.T) {
	config := &exchange.Config{
		Auth:    exchange.NewAuth(tdaKey, tdaPass, tdaSecret),
		Sandbox: true,
	}
	TDAClient, tdaErr = qsx.NewExchange("tdameritrade", config)
	if tdaErr != nil {
		t.Error(tdaErr)
	}
}
