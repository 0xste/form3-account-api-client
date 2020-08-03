package bic

import "fmt"

const (
	errMsgInvalidBusinessIdentifierCode string = "business identifier code %s is invalid"
	errMsgFailureToCompileRegex         string = "failed to compile regex"
)

// ErrInvalidBankIdentifierCode is returned if BIC for a company is invalid.
type ErrInvalidBankIdentifierCode struct {
	BusinessIdentifierCode BankIdentifierCode
}

func (e *ErrInvalidBankIdentifierCode) Error() string {
	return fmt.Sprintf(errMsgInvalidBusinessIdentifierCode, e.BusinessIdentifierCode.String())
}

// ErrFailureToGenerateUUID is returned if a UUID is not able to be generated.
type ErrFailureToCompileRegex struct{}

func (e *ErrFailureToCompileRegex) Error() string {
	return errMsgFailureToCompileRegex
}
