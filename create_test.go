package main

import (
	"context"
	"fmt"
	"form3-accountapi-client/account_type"
	bic2 "form3-accountapi-client/bic"
	"form3-accountapi-client/country"
	"form3-accountapi-client/currency"
	"form3-accountapi-client/uuid"
	"gopkg.in/h2non/gock.v1"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func createAccount(t time.Time, accountId, orgId uuid.UUID, bankId, bankIdCode string, bic bic2.BankIdentifierCode) AccountWrapper {
	return AccountWrapper{
		Data: Account{
			Attributes: AccountAttributes{
				AlternativeBankAccountNames: nil,
				BankId:                      bankId,
				BankIdCode:                  bankIdCode,
				Bic:                         bic,
				BaseCurrency:                currency.PoundSterling,
				Country:                     country.UnitedKingdomofGreatBritainandNorthernIrelandthe,
			},
			Id:             accountId,
			OrganisationId: orgId,
			AccountType:    account_type.TypeAccount,
		},
	}
}

func Test_accountClient_Create_ok(t *testing.T) {
	// arrange
	now := time.Now()
	accountId := uuid.MustUUID(uuid.FromStringV4("ad27e265-9605-4b4b-a0e5-3003ea9cc4dd"))
	orgId := uuid.MustUUID(uuid.FromStringV4("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"))
	bankId := "400300"
	bankIdCode := "GBDSC"
	bic := bic2.BankIdentifierCode("NWBKGB22")
	accountToCreate := createAccount(now, accountId, orgId, bankId, bankIdCode, bic)
	expectedRequestBody := `
	{
		"data": {
			"type": "accounts",
			"id": "ad27e265-9605-4b4b-a0e5-3003ea9cc4dd",
			"organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			"created_on": "0001-01-01T00:00:00Z",
			"modified_on": "0001-01-01T00:00:00Z",
			"attributes": {
            "alternative_bank_account_names": null,
				"country": "GB",
				"base_currency": "GBP",
				"bank_id": "400300",
				"bank_id_code": "GBDSC",
				"bic": "NWBKGB22"
			}
		}
	}`
	expectedResponseBody := fmt.Sprintf(`
	{
		"data": {
			"attributes": {
				"alternative_bank_account_names": null,
				"bank_id": "400300",
				"bank_id_code": "GBDSC",
				"base_currency": "GBP",
				"bic": "NWBKGB22",
				"country": "GB"
			},
			"created_on": "%s",
			"id": "ad27e265-9605-4b4b-a0e5-3003ea9cc4dd",
			"modified_on": "%s",
			"organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			"type": "accounts",
			"version": 0
		},
		"links": {
			"self": "/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dd"
		}
	}`, now.Format(time.RFC3339), now.Format(time.RFC3339))

	defer gock.Off()
	gock.New("http://server.com").
		Post("/v1/organisation/accounts").
		MatchType("json").
		BodyString(expectedRequestBody).
		Reply(http.StatusCreated).
		BodyString(expectedResponseBody)

	accountClient, err := NewAccountClient("http", "server.com", 80, 1)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}

	// act
	account, err := accountClient.Create(context.TODO(), &accountToCreate.Data)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}

	//assert

	expectedResponse := Account{
		Attributes:     accountToCreate.Data.Attributes,
		CreatedOn:      now,
		Id:             accountId,
		ModifiedOn:     now,
		OrganisationId: orgId,
		AccountType:    account_type.TypeAccount,
		Version:        0,
	}
	// todo something weird happening on deepequals for struct timestamps
	// revisit if time permits
	if expectedResponse.CreatedOn.Format(time.RFC3339) != now.Format(time.RFC3339) {
		t.Errorf("\ngot  :%v\nwant :%v", expectedResponse.CreatedOn, now)
	}
	if expectedResponse.Id != accountId {
		t.Errorf("\ngot  :%v\nwant :%v", expectedResponse.Id, accountId)
	}
	if expectedResponse.ModifiedOn != now {
		t.Errorf("\ngot  :%v\nwant :%v", expectedResponse.ModifiedOn, now)
	}
	if expectedResponse.OrganisationId != orgId {
		t.Errorf("\ngot  :%v\nwant :%v", expectedResponse.OrganisationId, orgId)
	}
	if !reflect.DeepEqual(account.Attributes, expectedResponse.Attributes) {
		t.Errorf("\ngot  :%v\nwant :%v", account.Attributes, expectedResponse.Attributes)
	}
}

func Test_accountClient_Create_duplicate(t *testing.T) {
	// arrange
	now := time.Now()
	accountId := uuid.MustUUID(uuid.FromStringV4("ad27e265-9605-4b4b-a0e5-3003ea9cc4dd"))
	orgId := uuid.MustUUID(uuid.FromStringV4("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"))
	bankId := "400300"
	bankIdCode := "GBDSC"
	bic := bic2.BankIdentifierCode("NWBKGB22")
	accountToCreate := createAccount(now, accountId, orgId, bankId, bankIdCode, bic)
	expectedRequestBody := `
	{
		"data": {
			"type": "accounts",
			"id": "ad27e265-9605-4b4b-a0e5-3003ea9cc4dd",
			"organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			"created_on": "0001-01-01T00:00:00Z",
			"modified_on": "0001-01-01T00:00:00Z",
			"attributes": {
            "alternative_bank_account_names": null,
				"country": "GB",
				"base_currency": "GBP",
				"bank_id": "400300",
				"bank_id_code": "GBDSC",
				"bic": "NWBKGB22"
			}
		}
	}`
	expectedResponseBody := `
	{
    "error_message": "Account cannot be created as it violates a duplicate constraint"
	}`

	defer gock.Off()
	gock.New("http://server.com").
		Post("/v1/organisation/accounts").
		MatchType("json").
		BodyString(expectedRequestBody).
		Reply(http.StatusConflict).
		BodyString(expectedResponseBody)

	accountClient, err := NewAccountClient("http", "server.com", 80, 1)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}

	// act
	_, err = accountClient.Create(context.TODO(), &accountToCreate.Data)

	//assert
	if _, ok := err.(ErrDuplicateAccount); ok {
		t.Errorf("expected duplicate account failure but was  %v", err)
	}

}
