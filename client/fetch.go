package client

import (
	"context"
	"encoding/json"
	"form3-accountapi-client/uuid"
	"io/ioutil"
	"log"
	"net/http"
)

// retrieves an account
func (a *accountClient) Fetch(ctx context.Context, accountId uuid.UUID) (Account, error) {

	accountPath := a.baseUrl.String() + endpointAccounts + "/" + accountId.String()
	requestMethod := http.MethodGet

	resp, err := a.httpClient.Get(accountPath)
	if err != nil {
		return Account{}, newGenericAccountError(requestMethod, accountPath, err.Error(), resp.StatusCode)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return Account{}, &ErrAccountNotFound{accountId, -1} // todo tidy this up
	}
	if resp.StatusCode != http.StatusOK {
		return Account{}, newGenericAccountError(requestMethod, accountPath, "invalid response code from api", resp.StatusCode)
	}

	var accountResponse AccountWrapper
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Account{}, newGenericAccountError(requestMethod, accountPath, err.Error(), resp.StatusCode)
	}
	err = json.Unmarshal(body, &accountResponse)
	if err != nil {
		return Account{}, newGenericAccountError(requestMethod, accountPath, err.Error(), resp.StatusCode)
	}

	if err := accountResponse.Data.Validate(); err != nil {
		return Account{}, newGenericAccountError(requestMethod, accountPath, err.Error(), resp.StatusCode)
	}

	log.Printf("successfully retrieved account %s", accountId)
	return accountResponse.Data, nil
}

