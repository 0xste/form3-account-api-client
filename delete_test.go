package main

import (
	"context"
	"form3-accountapi-client/uuid"
	"gopkg.in/h2non/gock.v1"
	"net/http"
	"testing"
)

func Test_accountClient_Delete_ok_or_non_existent(t *testing.T) {
	// arrange
	accountId := uuid.MustUUID(uuid.FromStringV4("ad27e265-9605-4b4b-a0e5-3003ea9cc4dd"))
	accountVersion := int64(0)

	defer gock.Off()
	gock.New("http://server.com").
		Delete("/v1/organisation/accounts/" + accountId.String()).
		Reply(http.StatusNoContent).
		BodyString("")

	accountClient, err := NewAccountClient("http", "server.com", 80, 1)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}

	// act
	err = accountClient.Delete(context.TODO(), accountId, accountVersion)

	//assert
	if _, ok := err.(ErrDuplicateAccount); ok {
		t.Errorf("expected duplicate account failure but was  %v", err)
	}

}

func Test_accountClient_Delete_existing_non_existent_version(t *testing.T) {
	// arrange
	accountId := uuid.MustUUID(uuid.FromStringV4("ad27e265-9605-4b4b-a0e5-3003ea9cc4dd"))
	accountVersion := int64(99)
	response := `{
    	"error_message": "invalid version"
	}`

	defer gock.Off()
	gock.New("http://server.com").
		Delete("/v1/organisation/accounts/" + accountId.String()).
		Reply(http.StatusNotFound).
		BodyString(response)

	accountClient, err := NewAccountClient("http", "server.com", 80, 1)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}

	// act
	err = accountClient.Delete(context.TODO(), accountId, accountVersion)

	//assert
	if _, ok := err.(ErrBadAccountRequest); ok {
		t.Errorf("expected bad request failure but was  %v", err)
	}

}
