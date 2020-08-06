package country

import "fmt"

const (
	errMsgInvalidCountry string = "invalid country %s"
)

// ErrInvalidCountry is returned if a a country is invalid.
type ErrInvalidCountry struct {
	Country Country
}

func (e *ErrInvalidCountry) Error() string {
	return fmt.Sprintf(errMsgInvalidCountry, e.Country.String())
}
