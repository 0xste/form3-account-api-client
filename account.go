package accountapi_client

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	httpDefaultTimeout time.Duration = 5
	protocolHttp                     = "http"
	protocolHttps                    = "https"
	endpointHealth                   = "/health"
)

type accountClient struct {
	log        log.Logger
	baseUrl    *url.URL
	httpClient http.Client
}

// Creates a new instance of an account client
func NewAccountClient(protocol, host string, port, version int) (accountClient, error) {
	if err := isValidProtocol(protocol); err != nil {
		return accountClient{}, err
	}
	if err := isValidPort(port); err != nil {
		return accountClient{}, err
	}

	uri := fmt.Sprintf("%s://%s:%d/%d", protocol, host, port, version)
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
