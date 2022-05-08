package coinbasepro

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/quantstop/qsx/core"
	"golang.org/x/time/rate"
	"net/http"
	"strconv"
	"time"
)

const (

	// Base URL's
	coinbaseproAPIURL       = "https://api.pro.coinbase.com"
	coinbaseproWebsocketURL = "wss://ws-feed.exchange.coinbase.com"

	// Base URL's for sandbox environment
	coinbaseproSandboxWebsiteURL   = "https://public.sandbox.exchange.coinbase.com"
	coinbaseproSandboxRestAPIURL   = "https://api-public.sandbox.exchange.coinbase.com"
	coinbaseproSandboxWebsocketURL = "wss://ws-feed-public.sandbox.exchange.coinbase.com"
	coinbaseproSandboxFixAPIURL    = "tcp+ssl://fix-public.sandbox.exchange.coinbase.com:4198"

	// Endpoints
	coinbaseproAccounts                = "accounts"
	coinbaseproProducts                = "products"
	coinbaseproOrderbook               = "book"
	coinbaseproTicker                  = "ticker"
	coinbaseproTrades                  = "trades"
	coinbaseproHistory                 = "candles"
	coinbaseproStats                   = "stats"
	coinbaseproCurrencies              = "currencies"
	coinbaseproLedger                  = "ledger"
	coinbaseproHolds                   = "holds"
	coinbaseproOrders                  = "orders"
	coinbaseproFills                   = "fills"
	coinbaseproTransfers               = "transfers"
	coinbaseproReports                 = "reports"
	coinbaseproTime                    = "time"
	coinbaseproMarginTransfer          = "profiles/margin-transfer"
	coinbaseproPosition                = "position"
	coinbaseproPositionClose           = "position/close"
	coinbaseproPaymentMethod           = "payment-methods"
	coinbaseproPaymentMethodDeposit    = "deposits/payment-method"
	coinbaseproDepositCoinbase         = "deposits/coinbase-account"
	coinbaseproWithdrawalPaymentMethod = "withdrawals/payment-method"
	coinbaseproWithdrawalCoinbase      = "withdrawals/coinbase"
	coinbaseproWithdrawalCrypto        = "withdrawals/crypto"
	coinbaseproCoinbaseAccounts        = "coinbase-accounts"
	coinbaseproTrailingVolume          = "users/self/trailing-volume"
)

type CoinbasePro struct {
	core.Exchange
	Conn *websocket.Conn
}

func NewCoinbasePro(auth *core.Auth) (core.Qsx, error) {

	t := transport{
		authKey:        auth.Key,
		authPassphrase: auth.Passphrase,
		authSecret:     auth.Secret,
		timestamp: func() string {
			return strconv.FormatInt(time.Now().Unix(), 10)
		},
	}

	rl := rate.NewLimiter(rate.Every(time.Second), 10) // 10 requests per second

	api := core.New(
		&http.Client{
			Transport:     &t,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       0,
		},
		core.Options{
			ApiURL:  coinbaseproSandboxWebsiteURL,
			Verbose: false,
		},
		rl,
	)

	ws := &core.Dialer{
		URL: coinbaseproSandboxWebsocketURL,
	}

	return &CoinbasePro{
		core.Exchange{
			Name:      core.CoinbasePro,
			Crypto:    true,
			Auth:      auth,
			API:       api,
			Websocket: ws,
		},
		&websocket.Conn{},
	}, nil
}

type transport struct {
	authKey        string
	authPassphrase string
	authSecret     string
	timestamp      func() string
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {

	var b bytes.Buffer
	if req.Body != nil {
		err := json.NewEncoder(&b).Encode(req.Body)
		if err != nil {
			return nil, fmt.Errorf("qsx coinbase: error encoding content: %w", err)
		}
	}
	timestamp := t.timestamp()
	msg := fmt.Sprintf("%s%s%s%s", timestamp, req.Method, req.URL, b.Bytes())
	signature, err := core.SignSHA256HMAC(msg, t.authSecret)
	if err != nil {
		return nil, fmt.Errorf("qsx coinbase: error signing content: %w", err)
	}

	r := req.Clone(req.Context())
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept", "application/json")
	r.Header.Add("Accept-Charset", "UTF-8")
	r.Header.Add("User-Agent", "qsx v0.0.1")
	r.Header.Add("CB-ACCESS-KEY", t.authKey)
	r.Header.Add("CB-ACCESS-PASSPHRASE", t.authPassphrase)
	r.Header.Add("CB-ACCESS-TIMESTAMP", timestamp)
	r.Header.Add("CB-ACCESS-SIGN", signature)

	return http.DefaultTransport.RoundTrip(r)
}
