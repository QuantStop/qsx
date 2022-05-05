package coinbasepro

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/quantstop/qsx/qsx"
	"math"
	"strings"
	"time"
)

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
		return fmt.Errorf("timeperiod(%ds) is invalid", t)
	}
}

type TimePeriodParam time.Duration

func (t TimePeriodParam) Validate() error {
	return t.TimePeriod().Valid()
}

func (t TimePeriodParam) TimePeriod() TimePeriod {
	return TimePeriod(int(math.Round(time.Duration(t).Seconds())))
}

func (t *TimePeriodParam) UnmarshalJSON(b []byte) error {
	var s string
	// quote bytes so that marshaller properly scans a number followed by a string as a single string
	err := json.Unmarshal([]byte(fmt.Sprintf("%q", b)), &s)
	if err != nil {
		return err
	}
	d, err := time.ParseDuration(strings.ReplaceAll(s, "\"", ""))
	if err != nil {
		return err
	}
	*t = TimePeriodParam(d)
	return nil
}

// HistoricRateFilter holds filters historic rates for a product by date and sets the granularity of the response.
// If either one of the start or end fields are not provided then both fields will be ignored.
// If a custom time range is not declared then one ending now is selected.
// The granularity field must be one of the following values:
//  {60, 300, 900, 3600, 21600, 86400}.
// Otherwise, the request will be rejected. These values correspond to time slices representing:
// one minute, five minutes, fifteen minutes, one hour, six hours, and one day, respectively.
// If data points are readily available, the response may contain as many as 300 candles and some candles
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

type HistoricRates struct {
	Candles []*Candle
}

func (h *HistoricRates) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &h.Candles)
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

func (c *Candle) UnmarshalJSON(b []byte) error {
	var tmp []json.RawMessage
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	if len(tmp) != 6 {
		return fmt.Errorf("a Candle must have 6 elements, only found %d", len(tmp))
	}
	var rawTime int64
	if err := json.Unmarshal(tmp[0], &rawTime); err != nil {
		return err
	}
	c.Time = Time(time.Unix(rawTime, 0).UTC())
	if err := json.Unmarshal(tmp[1], &c.Low); err != nil {
		return err
	}
	if err := json.Unmarshal(tmp[2], &c.High); err != nil {
		return err
	}
	if err := json.Unmarshal(tmp[3], &c.Open); err != nil {
		return err
	}
	if err := json.Unmarshal(tmp[4], &c.Close); err != nil {
		return err
	}
	return json.Unmarshal(tmp[5], &c.Volume)
}

// GetHistoricRates retrieves historic rates, as Candles, for a Product. Rates grouped buckets based on requested Granularity.
// If either one of the start or End fields are not provided then both fields will be ignored.
// The Granularity is limited to a set of supported Time slices, one of:
//   one minute, five minutes, fifteen minutes, one hour, six hours, or one day.
func (c *CoinbasePro) GetHistoricRates(ctx context.Context, productID string, filter HistoricRateFilter) (HistoricRates, error) {
	var history HistoricRates
	path := fmt.Sprintf("%s/%s/%s/%s", coinbaseproProducts, productID, coinbaseproHistory, qsx.Query(filter.Params()))
	if err := c.API.Get(ctx, path, &history); err != nil {
		return history, err
	}
	return history, nil
}
