package currency

type Currencies []Currency
type Currency string

const (
	// ISO 3166 country codes
	PoundSterling Currency = "GBP"
)

func (c *Currency) String() string {
	return string(*c)
}

// validates a single country
func (c *Currency) Validate() error {
	allCurrencies := getAllCurrencies()
	for _, currency := range allCurrencies {
		if currency.String() == c.String() {
			return nil
		}
	}
	return &ErrInvalidCurrency{
		ValidCurrencies: allCurrencies,
		Currency:        *c,
	}
}

// getAllCurrencies is a helper for returning all valid currencies
func getAllCurrencies() Currencies {
	return Currencies{PoundSterling}
}
