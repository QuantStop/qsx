package tdameritrade

import (
	"context"
	"github.com/quantstop/qsx/exchange"
	"golang.org/x/oauth2"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

const (
	apiURL   = "https://api.tdameritrade.com"
	tokenURL = "https://api.tdameritrade.com/v1/oauth2/token"
	authURL  = "https://auth.tdameritrade.com/auth"

	// Authenticated endpoints
	chains = "/v1/marketdata/chains"
)

var (
	authConfig = oauth2.Config{
		ClientID:     "XXXX-XXXX-XXXX-XXXX",
		ClientSecret: "YYYY-YYYY-YYYY-YYYY",
		RedirectURL:  "https://localhost/callback",
		Scopes:       []string{"all"},
		Endpoint: oauth2.Endpoint{
			AuthStyle: oauth2.AuthStyleInParams,
			AuthURL:   authURL,
			TokenURL:  tokenURL,
		},
	}
)

type TDAmeritrade struct {
	exchange.Exchange
}

func NewTDAmeritrade(config *exchange.Config) (exchange.IExchange, error) {

	httpClient := authConfig.Client(
		context.WithValue(context.Background(), oauth2.HTTPClient, &http.Client{Transport: &transport{}}),
		config.Token,
	)

	rl := rate.NewLimiter(rate.Every(time.Second), 10) // 10 requests per second

	api := exchange.New(
		httpClient,
		exchange.Options{
			ApiURL:  apiURL,
			Verbose: false,
		},
		rl,
	)

	return &TDAmeritrade{
		exchange.Exchange{
			Name: exchange.TDAmeritrade,
			Auth: config.Auth,
			API:  api,
		},
	}, nil
}

type transport struct{}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	r := req.Clone(req.Context())
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept", "application/json")
	return http.DefaultTransport.RoundTrip(r)
}
