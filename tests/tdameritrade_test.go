package tests

import (
	"github.com/quantstop/qsx"
	"github.com/quantstop/qsx/core"
	"testing"
)

var tdaKey = ""
var tdaPass = ""
var tdaSecret = ""

var TDAClient core.Qsx
var tdaErr error

func TestTDAClient(t *testing.T) {
	config := &core.Config{
		Auth:    core.NewAuth(tdaKey, tdaPass, tdaSecret),
		Sandbox: true,
	}
	TDAClient, tdaErr = qsx.NewExchange("tdameritrade", config)
	if tdaErr != nil {
		t.Error(tdaErr)
	}
}
