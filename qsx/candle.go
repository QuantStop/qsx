package qsx

import "time"

type Candle struct {
	Close  float64
	High   float64
	Low    float64
	Open   float64
	Time   time.Time
	Volume float64
}
