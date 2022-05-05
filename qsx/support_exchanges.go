package qsx

type ExchangeName string

const (
	CoinbasePro ExchangeName = "coinbasepro"
	Binance     ExchangeName = "binance"
	YFinance    ExchangeName = "yfinance"
)

var SupportedExchanges = []ExchangeName{
	CoinbasePro,
}
