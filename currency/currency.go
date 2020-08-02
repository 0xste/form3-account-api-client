package currency

type Currencies []Currency
type Currency string

const (
	// ISO 3166 country codes
	PoundSterling Currency = "GBP"
)

// validates a single country
func (c *Currency) Validate() error {
	for _, country := range *c.getAllCurrencies() {
		if country == *c {
			return nil
		}
	}
	return &ErrInvalidCurrency{
		ValidCurrencies: c.getAllCurrencies(),
		Currency:        c,
	}
}

// getAllCurrencies is a helper for returning all valid currencies
func (c *Currency) getAllCurrencies() *Currencies {
	return &Currencies{PoundSterling}
}
