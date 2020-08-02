package accountapi_client

import (
	"context"
	"fmt"
	"net/url"
)

type accountClient struct {
	baseUrl *url.URL
}

func NewAccountClient(protocol, host, port, version string) (accountClient, error) {
	baseUrl, err := url.ParseRequestURI(fmt.Sprintf("%s://%s:%s/%s", protocol, host, port, version))
	if err != nil {
		return accountClient{}, err
	}
	return accountClient{
		baseUrl: baseUrl,
	}, nil
}

func (a *accountClient) getHealth(ctx context.Context) {

}
