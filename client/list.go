package client

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (a *accountClient) List(ctx context.Context, limit int, offset int) ([]Account, error) {
	accountPath := a.baseUrl.String() + endpointAccounts
	requestMethod := http.MethodGet
	req, err := http.NewRequest(requestMethod, accountPath, nil)
	if err != nil {
		return []Account{}, &ErrInvalidRequest{requestMethod, endpointAccounts, err.Error()}
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return []Account{}, newGenericAccountError(requestMethod, accountPath, err.Error(), resp.StatusCode)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []Account{}, newGenericAccountError(requestMethod, accountPath, "invalid response code from api", resp.StatusCode)
	}

	var accountsResponse AccountsWrapper
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Account{}, newGenericAccountError(requestMethod, accountPath, err.Error(), resp.StatusCode)
	}
	err = json.Unmarshal(body, &accountsResponse)
	if err != nil {
		return []Account{}, newGenericAccountError(requestMethod, accountPath, err.Error(), resp.StatusCode)
	}

	return paginate(accountsResponse.Data, limit, offset), nil

}

func paginate(data []Account, limit, offset int) []Account {
	if offset > len(data) {
		offset = len(data)
	}
	end := offset + limit
	if end > len(data) {
		end = len(data)
	}
	return data[offset:end]
}

type AccountsWrapper struct{
	Data []Account
	Links map[string]string
}