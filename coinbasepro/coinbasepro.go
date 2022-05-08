package coinbasepro

import (
	"context"
	"github.com/quantstop/qsx/core"
	"sync"
	"time"
)

// This file holds all methods that implement the main Qsx interface.
// This is served as an adapter to the Coinbase client, which adheres strictly to the Coinbase API.
// Here is where the data is formatted into the common type defined by Qsx.

func (c *CoinbasePro) GetHistoricalCandles(ctx context.Context, productID string) ([]core.Candle, error) {
	var candles []core.Candle
	coinbaseCandles, err := c.GetHistoricRates(ctx, productID, HistoricRateFilter{
		Granularity: 60,
		End:         Time{},
		Start:       Time{},
	})
	if err != nil {
		return nil, err
	}

	for _, cbCandle := range coinbaseCandles.Candles {
		candles = append(candles, core.Candle{
			Close:  cbCandle.Close,
			High:   cbCandle.High,
			Low:    cbCandle.Low,
			Open:   cbCandle.Open,
			Time:   time.Time(cbCandle.Time),
			Volume: cbCandle.Volume,
		})
	}

	return candles, nil
}

func normalizeHistory(history HistoricRates, filter HistoricRateFilter) HistoricRates {
	normalizedHistory := HistoricRates{nil}
	//var now = time.Now()
	var nextCandleTime Time

	for index, candle := range history.Candles {

		normalizedHistory.Candles = append(normalizedHistory.Candles, candle)

		if index >= len(history.Candles)-1 {
			break
		} else {
			nextCandleTime = history.Candles[index+1].Time
		}

		_, _, _, _, min, _ := diff(time.Time(candle.Time), time.Time(nextCandleTime))
		switch filter.Granularity {
		case TimePeriod1Minute:
			if min > 1 {
				/*newTime := time.Time(candle.Time).Add(time.Minute)
				newCandle := &Candle{
					Close:  candle.Close,
					High:   candle.Close,
					Low:    candle.Close,
					Open:   candle.Close,
					Time: 	Time(newTime),
					Volume: 0,
				}
				normalizedHistory.Candles = insertCandle(normalizedHistory.Candles, index, newCandle)*/

				for i := 0; i < min-1; i++ {
					newTime := time.Time(candle.Time).Add(time.Minute)
					newCandle := &Candle{
						Close:  candle.Close,
						High:   candle.Close,
						Low:    candle.Close,
						Open:   candle.Close,
						Time:   Time(newTime),
						Volume: 0,
					}
					normalizedHistory.Candles = insertCandle(normalizedHistory.Candles, index+i, newCandle)
				}
				continue
			}
			continue

		case TimePeriod5Minutes:

		case TimePeriod15Minutes:

		case TimePeriod1Hour:

		case TimePeriod6Hours:

		case TimePeriod1Day:

		}
	}
	return normalizedHistory
}

func insertCandle(a []*Candle, index int, value *Candle) []*Candle {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func diff(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}

func (c *CoinbasePro) WatchFeed(shutdown chan struct{}, wg *sync.WaitGroup, product string, feed interface{}) error {

	// create a new subscription request
	prods := []ProductID{ProductID(product)}
	channelNames := []ChannelName{
		ChannelNameHeartbeat,
		ChannelNameLevel2,
	}
	channels := []Channel{
		{
			Name:       ChannelNameMatches,
			ProductIDs: []ProductID{ProductID(product)},
		},
	}
	subReq := NewSubscriptionRequest(prods, channelNames, channels)

	return c.Watch(shutdown, wg, subReq, feed.(*Feed))
}

func (c *CoinbasePro) ListProducts(ctx context.Context) ([]core.Product, error) {
	products, err := c.ListCoinbaseProducts(ctx)
	if err != nil {
		return nil, err
	}
	var returnArr []core.Product
	for _, product := range products {
		returnArr = append(returnArr, core.Product{
			ID:             product.ID,
			BaseCurrency:   string(product.BaseCurrency),
			QuoteCurrency:  string(product.QuoteCurrency),
			BaseMinSize:    product.BaseMinSize,
			BaseMaxSize:    product.BaseMaxSize,
			QuoteIncrement: product.QuoteIncrement,
			BaseIncrement:  product.BaseIncrement,
			DisplayName:    product.DisplayName,
			MinMarketFunds: product.MinMarketFunds,
			MaxMarketFunds: product.MaxMarketFunds,
			MarginEnabled:  product.MarginEnabled,
			PostOnly:       product.PostOnly,
			LimitOnly:      product.LimitOnly,
			CancelOnly:     product.CancelOnly,
			Status:         string(product.Status),
			StatusMessage:  product.StatusMessage,
			AuctionMode:    product.AuctionMode,
		})
	}

	return returnArr, nil
}
