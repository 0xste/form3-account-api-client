package main

import (
	"context"
	"encoding/json"
	"fmt"
	"form3-accountapi-client/uuid"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	httpDefaultTimeout time.Duration = 5
	protocolHttp       string        = "http"
	protocolHttps      string        = "https"
	endpointHealth     string        = "/health"
	endpointAccounts   string        = "/organisation/accounts"
)

type accountClient struct {
	baseUrl    *url.URL
	httpClient http.Client
}

type AccountClient interface {
	Fetch(ctx context.Context, accountId uuid.UUID) (Account, error)
	List(ctx context.Context, limit int, offset int) ([]Account, error)
	Create(ctx context.Context, account *Account) (Account, error)
	Delete(ctx context.Context, accountId uuid.UUID, version int64) error
}

// Creates a new instance of an account client
func NewAccountClient(protocol, host string, port, version int) (accountClient, error) {
	if err := isValidProtocol(protocol); err != nil {
		return accountClient{}, err
	}
	if err := isValidPort(port); err != nil {
		return accountClient{}, err
	}

	uri := fmt.Sprintf("%s://%s:%d/v%d", protocol, host, port, version)
	baseUrl, err := url.ParseRequestURI(uri)
	if err != nil {
		return accountClient{}, &ErrInvalidClientBaseUri{uri}
	}

	return accountClient{
		baseUrl: baseUrl,
		httpClient: http.Client{
			Timeout: httpDefaultTimeout * time.Second,
		},
	}, nil
}

func isValidPort(port int) error {
	if port < 1 || port > 99999 {
		return &ErrInvalidClientBaseUri{strconv.Itoa(port)}
	}
	return nil
}

func isValidProtocol(protocol string) error {
	for _, validProtocol := range []string{protocolHttp, protocolHttps} {
		if validProtocol == protocol {
			return nil
		}
	}
	return ErrInvalidClientBaseUri{protocol}
}

func parseAccountResponse(requestMethod, accountPath string, resp *http.Response) (Account, error) {
	var accountResponse AccountWrapper
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
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
	return accountResponse.Data, nil
}

func newGenericAccountError(method, baseUri, message string, statusCode int) *ErrRemoteGatewayFailure {
	return &ErrRemoteGatewayFailure{
		Method:     method,
		BaseUri:    baseUri,
		StatusCode: statusCode,
		Message:    message,
	}
}
