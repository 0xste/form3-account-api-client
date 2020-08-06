package account_type

import "fmt"

const (
	errMsgInvalidAccountType string = "account type %s is invalid"
)

type ErrInvalidAccountType struct {
	AccountType AccountType
}

func (e ErrInvalidAccountType) Error() string {
	return fmt.Sprintf(errMsgInvalidAccountType, e.AccountType.String())
}
