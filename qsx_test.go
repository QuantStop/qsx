package qsx

import (
	"context"
	"fmt"
	cbp "github.com/quantstop/qsx/coinbasepro"
	"github.com/quantstop/qsx/core"
	"golang.org/x/sync/errgroup"
	"sync"
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
		exchange, err := NewExchange(x, config)
		if err != nil {
			t.Error(err)
		}
		t.Logf("Exchange Name: %v", exchange.GetName())
	}

}

func TestCoinbaseCandles(t *testing.T) {
	config := &core.Config{
		Auth:    core.NewAuth(key, pass, secret),
		Sandbox: true,
	}
	coinbasepro, err := NewExchange("coinbasepro", config)
	if err != nil {
		t.Error(err)
	}
	candles, err := coinbasepro.GetHistoricalCandles(context.TODO(), "BTC-USD", "1m")
	if err != nil {
		t.Error(err)
	}
	for _, candle := range candles {
		t.Logf("Candle Time: %v | Open: %v | High: %v | Low: %v | Close: %v | Volume: %v", candle.Time, candle.Open, candle.High, candle.Low, candle.Close, candle.Volume)
	}
}

func TestCoinbaseListProducts(t *testing.T) {
	config := &core.Config{
		Auth:    core.NewAuth(key, pass, secret),
		Sandbox: true,
	}
	coinbasepro, err := NewExchange("coinbasepro", config)
	if err != nil {
		t.Error(err)
	}
	products, err := coinbasepro.ListProducts(context.TODO())
	if err != nil {
		t.Error(err)
	}
	for _, product := range products {
		t.Logf("Product: %v", product.ID)
	}
}

func TestCoinbaseFeed(t *testing.T) {
	config := &core.Config{
		Auth:    core.NewAuth(key, pass, secret),
		Sandbox: true,
	}
	coinbasepro, err := NewExchange("coinbasepro", config)
	if err != nil {
		t.Error(err)
	}
	feed := cbp.NewFeed()

	ctx := context.TODO()
	wg, ctx := errgroup.WithContext(ctx)

	s := &sync.WaitGroup{}
	shutdown := make(chan struct{})

	// start api client feed

	book, err := coinbasepro.WatchFeed(shutdown, s, "BTC-USD", feed)
	if err != nil {
		t.Error(err)
	}

	// Loop on Heartbeat channel
	wg.Go(func() error {
		for message := range feed.Heartbeat {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				out := fmt.Sprintf("%s | %s | %s | %v | %v", message.Type, message.Time.String(), message.ProductId, message.Sequence, message.LastTradeId)
				fmt.Println(out)

				bk := fmt.Sprintf("Best Bid: %v | Best Ask: %v ", book.GetBestBid(), book.GetBestOffer())
				fmt.Println(bk)
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
