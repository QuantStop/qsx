package coinbasepro

import (
	"context"
	"fmt"
	"github.com/quantstop/qsx/core"
	"math"
	"time"
)

// Product
// Only a maximum of one of trading_disabled, cancel_only, post_only, limit_only can be true at once. If none are true,
// the product is trading normally.
// !! When limit_only is true, matching can occur if a limit order crosses the book.
// !! Product ID will not change once assigned to a Product but all other fields are subject to change.
type Product struct {
	ID string `json:"id"`

	// BaseCurrency is the base in the pair of currencies that comprise the Product
	BaseCurrency CurrencyName `json:"base_currency"`

	// QuoteCurrency
	QuoteCurrency CurrencyName `json:"quote_currency"`

	// BaseMinSize defines the minimum order size
	BaseMinSize string `json:"base_min_size"`

	// BaseMaxSize defines the maximum order size
	BaseMaxSize string `json:"base_max_size"`

	// QuoteIncrement
	QuoteIncrement string `json:"quote_increment"`

	// BaseIncrement specifies the minimum increment for the BaseCurrency
	BaseIncrement string `json:"base_increment"`

	// DisplayName
	DisplayName string `json:"display_name"`

	// MinMarketFunds defines the minimum funds allowed
	MinMarketFunds string `json:"min_market_funds"`

	// MaxMarketFunds defines the maximum funds allowed
	MaxMarketFunds string `json:"max_market_funds"`

	// MarginEnabled
	MarginEnabled bool `json:"margin_enabled"`

	// PostOnly indicates whether only maker orders can be placed. No orders will be matched when post_only mode is active.
	PostOnly bool `json:"post_only"`

	// LimitOnly indicates whether this product only accepts limit orders.
	LimitOnly bool `json:"limit_only"`

	// CancelOnly indicates whether this product only accepts cancel requests for orders.
	CancelOnly bool `json:"cancel_only"`

	// Status
	Status ProductStatus `json:"status"`

	// StatusMessage provides any extra information regarding the status, if available
	StatusMessage string `json:"status_message"`

	// AuctionMode
	AuctionMode bool `json:"auction_mode"`
}

// ProductID values could perhaps be dynamically validated from '/products' endpoint
type ProductID string

// ProductStatus has little documentation; all sandbox products have a status value of `online`
type ProductStatus string

// BookLevel represents the level of detail/aggregation in an OrderBook.
// BookLevelBest and BookLevelTop50 are aggregates.
// BookLevelFull requests the entire order book.
type BookLevel int

const (
	// BookLevelUndefined defaults to BookLevel_Best
	BookLevelUndefined BookLevel = 0
	// BookLevelBest requests only the best bid and ask and is aggregated.
	BookLevelBest BookLevel = 1
	// BookLevelTop50 requests the top 50 bids and asks and is aggregated.
	BookLevelTop50 BookLevel = 2
	// BookLevelFull is non-aggregated and returns the entire order book.
	BookLevelFull BookLevel = 3
)

func (p BookLevel) Params() []string {
	level := p
	if p == BookLevelUndefined {
		level = BookLevelBest
	}
	return []string{fmt.Sprintf("level=%d", level)}
}

type AggregatedOrderBook struct {
	Sequence int                   `json:"sequence"`
	Bids     []AggregatedBookEntry `json:"bids"`
	Asks     []AggregatedBookEntry `json:"asks"`
}

type AggregatedBookEntry struct {
	Price     string `json:"price"`
	Size      string `json:"size"`
	NumOrders int    `json:"num_orders"`
}

type OrderBook struct {
	Sequence int         `json:"sequence"`
	Bids     []BookEntry `json:"bids"`
	Asks     []BookEntry `json:"asks"`
}

type BookEntry struct {
	Price   string `json:"price"`
	Size    string `json:"size"`
	OrderID string `json:"order_id"`
}

// HistoricRateFilter holds filters historic rates for a product by date and sets the granularity of the response.
// If either one of the start or end fields are not provided then both fields will be ignored.
// If a custom time range is not declared then one ending now is selected.
// The granularity field must be one of the following values:
//  {60, 300, 900, 3600, 21600, 86400}.
// Otherwise, the request will be rejected. These values correspond to time slices representing:
// one minute, five minutes, fifteen minutes, one hour, six hours, and one day, respectively.
// If data points are readily available, the response may contain as many as 300 candles and some of candles
// may precede the start value. The maximum number of data points for a single request is 300 candles.
// If the start/end time and granularity results in more than 300 data points, the request will be rejected.
// To retrieve finer granularity data over a larger time range, multiple requests with new start/end ranges must be used.
type HistoricRateFilter struct {
	Granularity TimePeriod `json:"granularity"`
	End         Time       `json:"end"`
	Start       Time       `json:"start"`
}

func (h *HistoricRateFilter) Params() []string {
	params := []string{fmt.Sprintf("granularity=%d", h.Granularity)}
	if !h.End.Time().IsZero() {
		end := h.End.Time().Format(time.RFC3339Nano)
		params = append(params, fmt.Sprintf("end=%s", end))
	}
	if !h.Start.Time().IsZero() {
		start := h.Start.Time().Format(time.RFC3339Nano)
		params = append(params, fmt.Sprintf("start=%s", start))
	}
	return params
}

type TimePeriodParam time.Duration

func (t TimePeriodParam) Validate() error {
	return t.TimePeriod().Valid()
}

func (t TimePeriodParam) TimePeriod() TimePeriod {
	return TimePeriod(int(math.Round(time.Duration(t).Seconds())))
}

type TimePeriod int

const (
	TimePeriod1Minute   TimePeriod = 60
	TimePeriod5Minutes  TimePeriod = 300
	TimePeriod15Minutes TimePeriod = 900
	TimePeriod1Hour     TimePeriod = 3600
	TimePeriod6Hours    TimePeriod = 21600
	TimePeriod1Day      TimePeriod = 86400
)

func (t TimePeriod) Valid() error {
	switch t {
	case TimePeriod1Minute, TimePeriod5Minutes, TimePeriod15Minutes, TimePeriod1Hour, TimePeriod6Hours, TimePeriod1Day:
		return nil
	default:
		return fmt.Errorf("timeslice(%ds) is invalid", t)
	}
}

type HistoricRates struct {
	Candles []*Candle
}

// A Candle is a common representation of a historic rate.
type Candle struct {
	Close  float64 `json:"close"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Open   float64 `json:"open"`
	Time   Time    `json:"time"`
	Volume float64 `json:"volume"`
}

// ProductTicker holds snapshot information about the last trade (tick), best bid/ask and 24h volume.
type ProductTicker struct {
	Ask     float64 `json:"ask"`
	Bid     float64 `json:"bid"`
	Price   float64 `json:"price"`
	Size    float64 `json:"size"`
	TradeID int     `json:"trade_id"`
	Time    Time    `json:"time"`
	Volume  float64 `json:"volume"`
}

// ProductTrades represents the latest trades for a product
type ProductTrades struct {
	Trades []*ProductTrade `json:"trades"`
	Page   *Pagination     `json:"page"`
}

type ProductTrade struct {
	Price   float64 `json:"price,string"`
	Side    Side    `json:"side"`
	Size    float64 `json:"size,string"`
	Time    Time    `json:"time"`
	TradeID int     `json:"trade_id"`
}

// ProductStats represents a 24 hr stat for the product.
// Volume is in base currency units.
// Open, High, Low are in quote currency units.
type ProductStats struct {
	High        float64 `json:"high"`
	Last        float64 `json:"last"`
	Low         float64 `json:"low"`
	Open        float64 `json:"open"`
	Volume      float64 `json:"volume"`
	Volume30Day float64 `json:"volume_30day"`
}

// Market Data

// ListCoinbaseProducts retrieves the list Currency pairs available for trading. The list is not paginated.
func (c *CoinbasePro) ListCoinbaseProducts(ctx context.Context) ([]Product, error) {
	var products []Product
	return products, c.API.Get(ctx, fmt.Sprintf("/%s", coinbaseproProducts), &products)
}

// GetProduct retrieves the details of a single Currency pair.
func (c *CoinbasePro) GetProduct(ctx context.Context, productID ProductID) (Product, error) {
	var product Product
	path := fmt.Sprintf("/%s/%s", coinbaseproProducts, productID)
	return product, c.API.Get(ctx, path, &product)
}

// GetAggregatedOrderBook retrieves an aggregated, BookLevelBest (1) and BookLevelTop50 (2), representation of a Product
// OrderBook. Aggregated levels return only one Size for each active Price (as if there was only a single Order for that Size at the level).
func (c *CoinbasePro) GetAggregatedOrderBook(ctx context.Context, productID ProductID, level BookLevel) (AggregatedOrderBook, error) {
	var aggregatedBook AggregatedOrderBook
	path := fmt.Sprintf("/%s/%s/%s/%s", coinbaseproProducts, productID, coinbaseproOrderbook, core.Query(level.Params()))
	return aggregatedBook, c.API.Get(ctx, path, &aggregatedBook)
}

// GetOrderBook retrieves the full, un-aggregated OrderBook for a Product.
func (c *CoinbasePro) GetOrderBook(ctx context.Context, productID ProductID) (OrderBook, error) {
	var book OrderBook
	path := fmt.Sprintf("/%s/%s/%s/?level=3", coinbaseproProducts, productID, coinbaseproOrderbook)
	return book, c.API.Get(ctx, path, &book)
}

// GetProductTicker retrieves snapshot information about the last trade (tick), best bid/ask and 24h volume of a Product.
func (c *CoinbasePro) GetProductTicker(ctx context.Context, productID ProductID) (ProductTicker, error) {
	var ticker ProductTicker
	path := fmt.Sprintf("/%s/%s/%s", coinbaseproProducts, productID, coinbaseproTicker)
	return ticker, c.API.Get(ctx, path, &ticker)
}

// GetProductTrades retrieves a paginated list of the last trades of a Product.
func (c *CoinbasePro) GetProductTrades(ctx context.Context, productID ProductID, pagination PaginationParams) (ProductTrades, error) {
	var trades ProductTrades
	path := fmt.Sprintf("/%s/%s/%s/%s", coinbaseproProducts, productID, coinbaseproTrades, core.Query(pagination.Params()))
	return trades, c.API.Get(ctx, path, &trades)
}

// GetProductStats retrieves the 24hr stats for a Product. Volume is in base Currency units. Open, High, and Low are in quote Currency units.
func (c *CoinbasePro) GetProductStats(ctx context.Context, productID ProductID) (ProductStats, error) {
	var stats ProductStats
	path := fmt.Sprintf("/%s/%s/%s", coinbaseproProducts, productID, coinbaseproStats)
	return stats, c.API.Get(ctx, path, &stats)
}

// GetHistoricRates retrieves historic rates, as Candles, for a Product. Rates grouped buckets based on requested Granularity.
// If either one of the start or End fields are not provided then both fields will be ignored.
// The Granularity is limited to a set of supported Time slices, one of:
//   one minute, five minutes, fifteen minutes, one hour, six hours, or one day.
func (c *CoinbasePro) GetHistoricRates(ctx context.Context, productID string, filter HistoricRateFilter) (HistoricRates, error) {
	var history HistoricRates
	path := fmt.Sprintf("/%s/%s/%s/%s", coinbaseproProducts, productID, coinbaseproHistory, core.Query(filter.Params()))
	if err := c.API.Get(ctx, path, &history); err != nil {
		return history, err
	}
	return history, nil
}
