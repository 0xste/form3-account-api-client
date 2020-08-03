package bic

import (
	"regexp"
)

const patternBic string = "([a-zA-Z]{2,5})"

type BankIdentifierCodes []BankIdentifierCode
type BankIdentifierCode string

func (b *BankIdentifierCode) String() string {
	return string(*b)
}

func (b *BankIdentifierCode) Validate() error {
	bic := b.String()
	bicRegex, err := regexp.Compile(patternBic)
	if err != nil {
		return &ErrFailureToCompileRegex{}
	}
	if !bicRegex.MatchString(bic) {
		return &ErrInvalidBankIdentifierCode{*b}
	}
	return nil
}
