package accountapi_client

import (
	"form3-accountapi-client/account_type"
	"form3-accountapi-client/bic"
	"form3-accountapi-client/country"
	"form3-accountapi-client/currency"
	"form3-accountapi-client/uuid"
	"time"
)

type Account struct {
	AccountType      account_type.AccountType
	Id               uuid.UUID
	CreatedOn        time.Time
	ModifiedOn       time.Time
	OrganisationUnit uuid.UUID
	Version          uint64
	Attributes       AccountAttributes
}

func (a *Account) Validate() error {
	if err := a.AccountType.Validate(); err != nil {
		return &ErrBadAccountRequest{"AccountType",a.AccountType}
	}
	if a.CreatedOn.After(time.Now().UTC()) {
		return &ErrBadAccountRequest{"CreatedOn",a.CreatedOn}
	}
	if a.ModifiedOn.After(time.Now().UTC()) {
		return &ErrBadAccountRequest{"ModifiedOn",a.ModifiedOn}
	}
	if a.Version < 1 {
		return &ErrBadAccountRequest{"Version",a.Version}
	}
	return nil
}

type AccountAttributes struct {
	alternativeBankAccountNames []string
	bankId                      int64
	bankIdCode                  bic.BankIdentifierCode
	baseCurrency                currency.Currency
	country                     country.Country
}

type accountHateoas struct {
	data  Account
	links map[string]string
}
