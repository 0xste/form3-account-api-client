package client

import (
	"context"
	"fmt"
	"form3-accountapi-client/uuid"
	"log"
	"net/http"
)

func (a *accountClient) Delete(ctx context.Context, accountId uuid.UUID, version int64) error {
	accountPath := fmt.Sprintf("%s%s/%s?version=%d", a.baseUrl.String(), endpointAccounts, accountId.String(), version)
	requestMethod := http.MethodDelete
	req, err := http.NewRequest(requestMethod, accountPath, nil)
	if err != nil {
		return &ErrInvalidRequest{requestMethod, endpointAccounts, err.Error()}
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return newGenericAccountError(requestMethod, accountPath, err.Error(), resp.StatusCode)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return &ErrAccountNotFound{accountId, version}
	}
	if resp.StatusCode != http.StatusNoContent {
		return newGenericAccountError(requestMethod, accountPath, "invalid response code from api", resp.StatusCode)
	}

	log.Printf("successfully deleted account %s", accountId)
	return nil
}
