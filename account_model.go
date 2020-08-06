package main

import (
	"form3-accountapi-client/account_type"
	"form3-accountapi-client/bic"
	"form3-accountapi-client/country"
	"form3-accountapi-client/currency"
	"form3-accountapi-client/uuid"
	"strconv"
	"time"
)

type Account struct {
	Attributes     AccountAttributes        `json:"attributes"`
	CreatedOn      time.Time                `json:"created_on,omitempty"`
	Id             uuid.UUID                `json:"id"`
	ModifiedOn     time.Time                `json:"modified_on,omitempty"`
	OrganisationId uuid.UUID                `json:"organisation_id"`
	AccountType    account_type.AccountType `json:"type"`
	Version        int64                    `json:"version,omitempty"`
}

func (a *Account) WithAccountType(accountType account_type.AccountType) *Account {
	a.AccountType = accountType
	return a
}
func (a *Account) WithId(id uuid.UUID) *Account {
	a.Id = id
	return a
}
func (a *Account) WithCreatedOn(createdOn time.Time) *Account {
	a.CreatedOn = createdOn
	return a
}
func (a *Account) WithModifiedOn(modifiedOn time.Time) *Account {
	a.ModifiedOn = modifiedOn
	return a
}
func (a *Account) WithOrganisationId(ou uuid.UUID) *Account {
	a.OrganisationId = ou
	return a
}
func (a *Account) WithVersion(ver int64) *Account {
	a.Version = ver
	return a
}
func (a *Account) WithAttributes(attr AccountAttributes) *Account {
	a.Attributes = attr
	return a
}

func (a *Account) Validate() error {
	if err := a.AccountType.Validate(); err != nil {
		return &ErrBadAccountRequest{"AccountType", a.AccountType}
	}
	if _, err := uuid.FromStringV4(a.Id.String()); err != nil {
		return &ErrBadAccountRequest{"Id", a.Version}
	}
	if _, err := uuid.FromStringV4(a.OrganisationId.String()); err != nil {
		return &ErrBadAccountRequest{"OrganisationId", a.OrganisationId}
	}
	if err := a.Attributes.Validate(); err != nil {
		return err
	}
	return nil
}

type AccountAttributes struct {
	AlternativeBankAccountNames []string               `json:"alternative_bank_account_names"`
	BankId                      string                 `json:"bank_id"`
	BankIdCode                  string                 `json:"bank_id_code"`
	Bic                         bic.BankIdentifierCode `json:"bic"`
	BaseCurrency                currency.Currency      `json:"base_currency"`
	Country                     country.Country        `json:"country"`
}

func (a *AccountAttributes) WithAlternativeBankAccountNames(abn []string) *AccountAttributes {
	a.AlternativeBankAccountNames = abn
	return a
}
func (a *AccountAttributes) WithBankId(b string) *AccountAttributes {
	a.BankId = b
	return a
}
func (a *AccountAttributes) WithBankIdCode(b string) *AccountAttributes {
	a.BankIdCode = b
	return a
}

func (a *AccountAttributes) WithBic(b bic.BankIdentifierCode) *AccountAttributes {
	a.Bic = b
	return a
}
func (a *AccountAttributes) WithBaseCurrency(c currency.Currency) *AccountAttributes {
	a.BaseCurrency = c
	return a
}
func (a *AccountAttributes) WithCountry(c country.Country) *AccountAttributes {
	a.Country = c
	return a
}

func (a *AccountAttributes) Validate() error {
	if bankId, err := strconv.Atoi(a.BankId); err != nil || bankId <= 0 {
		return &ErrBadAccountRequest{"Attributes.BankId", a.BankId}
	}
	if a.BaseCurrency.Validate() != nil {
		return &ErrBadAccountRequest{"Attributes.BaseCurrency", a.BaseCurrency}
	}
	if a.Country.Validate() != nil {
		return &ErrBadAccountRequest{"Attributes.Country", a.Country}
	}
	return nil
}

type AccountWrapper struct {
	Data  Account           `json:"data"`
	Links map[string]string `json:"links,omitempty"`
}
