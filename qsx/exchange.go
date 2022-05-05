package qsx

// Exchange is the base type for all supported exchanges
// It implements common functionality and traits that all exchanges share
type Exchange struct {
	Name     ExchangeName
	isCrypto bool

	Auth *Auth
	API  *Client
}

// GetName implements the Qsx interface, and returns the Exchange's Name
func (base *Exchange) GetName() ExchangeName {
	return base.Name
}

func (base *Exchange) IsCrypto() bool {
	return base.isCrypto
}
