package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (a *accountClient) Create(ctx context.Context, account *Account) (Account, error) {
	accountPath := a.baseUrl.String() + endpointAccounts
	requestMethod := http.MethodPost
	acc, err := json.Marshal(&AccountWrapper{
		Data: *account,
	})
	if err != nil {
		return Account{}, &ErrInvalidRequest{requestMethod, endpointAccounts, err.Error()}
	}
	req, err := http.NewRequest(requestMethod, accountPath, bytes.NewBuffer(acc))
	if err != nil {
		return Account{}, &ErrInvalidRequest{requestMethod, endpointAccounts, err.Error()}
	}
	resp, err := a.httpClient.Do(req)
	if err != nil {
		return Account{}, newGenericAccountError(requestMethod, accountPath, err.Error(), resp.StatusCode)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusConflict {
		return Account{}, &ErrDuplicateAccount{account}
	}
	if resp.StatusCode == http.StatusBadRequest{
		all, err := ioutil.ReadAll(resp.Body)
		if err != nil{
			return Account{}, newGenericAccountError(requestMethod, accountPath, err.Error(), resp.StatusCode)
		}
		var respErr map[string]string
		err = json.Unmarshal(all, &respErr)
		if err != nil{
			return Account{}, newGenericAccountError(requestMethod, accountPath, err.Error(), resp.StatusCode)
		}
		errStr, ok := respErr["error_message"]
		if !ok{
			return Account{}, newGenericAccountError(requestMethod, accountPath, err.Error(), resp.StatusCode)
		}

		return Account{}, &ErrInvalidRequest{
			Method:  requestMethod,
			BaseUri: accountPath,
			ErrMsg:  errStr,
		}
	}

	if resp.StatusCode != http.StatusCreated {
		return Account{}, newGenericAccountError(requestMethod, accountPath, "invalid response code from api", resp.StatusCode)
	}

	accountResponse, err := parseAccountResponse(requestMethod, accountPath, resp)
	if err != nil {
		return Account{}, err
	}
	return accountResponse, nil
}

