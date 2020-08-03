package account_type

const (
	errMsgInvalidAccountType string = "account type %s is invalid"
)

type ErrInvalidAccountType struct {
	AccountType AccountType
}

func (e ErrInvalidAccountType) Error() string {
	return errMsgInvalidAccountType
}
