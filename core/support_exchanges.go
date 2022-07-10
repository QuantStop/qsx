package core

type ExchangeName string

const (
	CoinbasePro  ExchangeName = "coinbasepro"
	Binance      ExchangeName = "binance"
	YFinance     ExchangeName = "yfinance"
	TDAmeritrade ExchangeName = "tdameritrade"
)

var SupportedExchanges = []ExchangeName{
	CoinbasePro,
}
