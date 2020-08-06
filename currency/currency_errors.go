package currency

import (
	"fmt"
)

const errMsgInvalidCurrency string = "currency is invalid, must be one of '%p' but is '%s'"

// ErrInvalidCurrency is returned if Currency for a country is invalid.
type ErrInvalidCurrency struct {
	ValidCurrencies Currencies
	Currency        Currency
}

func (e *ErrInvalidCurrency) Error() string {
	return fmt.Sprintf(errMsgInvalidCurrency, e.ValidCurrencies, e.Currency.String())
}
