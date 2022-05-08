package qsx

import (
	"context"
	"fmt"
	cbp "github.com/quantstop/qsx/coinbasepro"
	"github.com/quantstop/qsx/qsx"
	"golang.org/x/sync/errgroup"
	"sync"
	"testing"
)

var key = "e47990d20df5c3f6ad2c0e1e098707c6"
var pass = "xieaj2i7dyf"
var secret = "WKaRL8SJKbBhDjRCIwIkmWTRB/bwYjLhXz0qcr467xUStUYsrplUuBlkLRInraVu7H4UqfsC1fF1Ybu01EObsQ=="

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

func TestCoinbaseFeed(t *testing.T) {
	auth := qsx.NewAuth(key, pass, secret)
	coinbasepro, err := NewExchange("coinbasepro", auth)
	if err != nil {
		t.Error(err)
	}
	feed := cbp.NewFeed()

	ctx := context.TODO()
	wg, ctx := errgroup.WithContext(ctx)

	s := &sync.WaitGroup{}
	shutdown := make(chan struct{})

	// start api client feed
	wg.Go(func() error {
		return coinbasepro.WatchFeed(shutdown, s, "BTC-USD", feed)
	})

	// Loop on Heartbeat channel
	wg.Go(func() error {
		for message := range feed.Heartbeat {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				out := fmt.Sprintf("%s | %s | %s | %v | %v", message.Type, message.Time.String(), message.ProductId, message.Sequence, message.LastTradeId)
				fmt.Println(out)
			}
		}
		return nil
	})

	// Loop on L2Channel channel
	wg.Go(func() error {
		for message := range feed.Level2 {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				out := fmt.Sprintf("%s | %s | %v | %v", message.Type, message.Time.String(), message.ProductId, message.Changes)
				fmt.Println(out)
			}
		}
		return nil
	})

	// Loop on L2ChannelSnapshot channel
	wg.Go(func() error {
		for message := range feed.Level2Snap {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				out := fmt.Sprintf("%s | %s", message.Type, message.ProductId)
				fmt.Println(out)
			}
		}
		return nil
	})

	// Loop on Matches channel
	wg.Go(func() error {
		for message := range feed.Matches {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				out := fmt.Sprintf("%s | %s | %s | %v", message.Type, message.Time.String(), message.ProductId, message.Price)
				fmt.Println(out)
			}
		}
		return nil
	})

	_ = wg.Wait()

}
