package core

import "time"

type Candle struct {
	Close  float64   `json:"close"`
	High   float64   `json:"high"`
	Low    float64   `json:"low"`
	Open   float64   `json:"open"`
	Time   time.Time `json:"time"`
	Volume float64   `json:"volume"`
}
