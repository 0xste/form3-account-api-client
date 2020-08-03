package accountapi_client

import (
	"form3-accountapi-client/account_type"
	"form3-accountapi-client/bic"
	"form3-accountapi-client/country"
	"form3-accountapi-client/currency"
	"form3-accountapi-client/uuid"
	"time"
)

type account struct {
	AccountType      account_type.AccountType
	Id               uuid.UUID
	CreatedOn        time.Time
	ModifiedOn       time.Time
	OrganisationUnit uuid.UUID
	Version          uint64
	Attributes       AccountAttributes
}

type AccountAttributes struct {
	alternativeBankAccountNames []string
	bankId                      int64
	bankIdCode                  bic.BankIdentifierCode
	baseCurrency                currency.Currency
	country                     country.Country
}

type accountHateoas struct {
	data  account
	links map[string]string
}
