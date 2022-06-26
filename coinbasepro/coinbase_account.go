package coinbasepro

import (
	"context"
	"fmt"
)

type CoinbaseAccount struct {
	Active                 bool                   `json:"active"`
	Balance                float64                `json:"balance"`
	Currency               CurrencyName           `json:"currency"`
	ID                     string                 `json:"id"`
	Name                   string                 `json:"name"`
	Primary                bool                   `json:"primary"`
	Type                   AccountType            `json:"type"`
	WireDepositInformation WireDepositInformation `json:"wire_deposit_information"`
	SEPADepositInformation SEPADepositInformation `json:"sepa_deposit_information"`
}

type AccountType string

const (
	CoinbaseAccountTypeFiat   AccountType = "fiat"
	CoinbaseAccountTypeWallet AccountType = "wallet"
)

type WireDepositInformation struct {
	Code           string  `json:"code"`
	Name           string  `json:"name"`
	AccountAddress string  `json:"account_address"`
	AccountName    string  `json:"account_name"`
	AccountNumber  string  `json:"account_number"`
	BankAddress    string  `json:"bank_address"`
	BankCountry    Country `json:"bank_country"`
	Reference      string  `json:"reference"`
	RoutingNumber  string  `json:"routing_number"`
}

type Country struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type SEPADepositInformation struct {
	AccountAddress  string `json:"account_address"`
	AccountName     string `json:"account_name"`
	BankAddress     string `json:"bank_address"`
	BankCountryName string `json:"bank_country_name"`
	BankName        string `json:"bank_name"`
	IBAN            string `json:"iban"`
	Reference       string `json:"reference"`
	Swift           string `json:"swift"`
}

// ListCoinbaseAccounts retrieves the list of CoinbaseAccounts available for the current Profile. The list is not paginated.
func (c *CoinbasePro) ListCoinbaseAccounts(ctx context.Context) ([]CoinbaseAccount, error) {
	var coinbaseAccounts []CoinbaseAccount
	path := fmt.Sprintf("/%s/", coinbaseproCoinbaseAccounts)
	return coinbaseAccounts, c.API.Get(ctx, path, &coinbaseAccounts)
}
