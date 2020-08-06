package bic

import (
	"regexp"
	"strings"
)

const patternBic string = "([a-zA-Z]{4})([a-zA-Z]{2})(([2-9a-zA-Z]{1})([0-9a-np-zA-NP-Z]{1}))((([0-9a-wy-zA-WY-Z]{1})([0-9a-zA-Z]{2}))|([xX]{3})|)"

type BankIdentifierCodes []BankIdentifierCode
type BankIdentifierCode string

func (b *BankIdentifierCode) String() string {
	return string(*b)
}

func (b *BankIdentifierCode) Validate() error {
	if strings.TrimSpace(b.String()) == "" {
		return &ErrInvalidBankIdentifierCode{*b}
	}
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
