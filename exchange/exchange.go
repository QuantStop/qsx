package exchange

import (
	"context"
	"github.com/quantstop/qsx/exchange/orderbook"
	"sync"
)

// IExchange is the interface which all supported exchanges must implement
type IExchange interface {

	// GetName returns the exchanges unique name
	GetName() Name

	// SupportsCrypto returns true if the exchange supports trading cryptocurrencies
	SupportsCrypto() bool

	// ListProducts returns and array of Product's
	// A Product is a market trading pair such as BTC-USD or stock such as AAPL
	ListProducts(ctx context.Context) ([]Product, error)

	// GetHistoricalCandles returns and array of Candle's in the normal format for candlestick data
	GetHistoricalCandles(ctx context.Context, productID string, granularity string) ([]Candle, error)

	// WatchFeed is the function to start watching a websocket feed for a specific product
	WatchFeed(shutdown chan struct{}, wg *sync.WaitGroup, product string, feed interface{}) (*orderbook.Orderbook, error)
}

// Exchange is the base type that all supported exchanges must include
// It implements common methods of the IExchange interface
type Exchange struct {

	// Name is the unique name of the exchange/marketplace
	Name Name

	// Crypto is a boolean to check if the exchange has cryptocurrency support
	Crypto bool

	// Auth is for authentication information (keys, tokens etc.)
	Auth *Auth

	// API is http(s) client connection and routes to the vendors' service
	API *Client

	// Websocket is the ws(s) client connection to the vendors' service
	Websocket Websocket
}

// GetName implements the IExchange interface, and returns the Exchange's Name
func (base *Exchange) GetName() Name {
	return base.Name
}

// SupportsCrypto implements the IExchange interface, and returns if the exchange has cryptocurrency support
func (base *Exchange) SupportsCrypto() bool {
	return base.Crypto
}
