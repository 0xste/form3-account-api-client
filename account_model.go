package accountapi_client

import (
	"form3-accountapi-client/currency"
	"form3-accountapi-client/uuid"
	"time"
)

type accountType string

const typeAccount accountType = "accounts"

type account struct {
	AccountType      accountType
	Id               uuid.UUID
	CreatedOn        time.Time
	ModifiedOn       time.Time
	OrganisationUnit uuid.UUID
	Version          uint64
	Attributes       accountAttributes
}

type BankIdCode string

func (a *accountType) Validate() error {
	return nil
}

type BusinessIdentifierCode string

func (b BusinessIdentifierCode) Validate() error {
	//ISO 9362 validation
	return nil
}

type accountAttributes struct {
	alternativeBankAccountNames []string
	bankId                      int64
	bankIdCode                  BankIdCode
	baseCurrency                currency.Currency
	bic                         BusinessIdentifierCode
	//            "bic": "NWBKGB22",
	//            "country": "GB"
}

type accountHateoas struct {
	data  account
	links map[string]string
}
