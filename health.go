package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// hits the Health api of the account service
func (a *accountClient) Health(ctx context.Context) error {

	healthPath := a.baseUrl.String() + endpointHealth
	requestMethod := http.MethodGet

	resp, err := a.httpClient.Get(healthPath)
	if err != nil {
		return &ErrRemoteGatewayFailure{
			requestMethod,
			healthPath,
			0,
			"failed to connect to remote gateway",
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return &ErrRemoteGatewayFailure{
			requestMethod,
			healthPath,
			resp.StatusCode,
			"invalid response code from api",
		}
	}

	var healthResponse HealthResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&healthResponse)
	if err != nil {
		return &ErrRemoteGatewayFailure{
			Method:     requestMethod,
			BaseUri:    healthPath,
			StatusCode: resp.StatusCode,
			Message:    "invalid response body from api",
		}
	}

	if strings.ToUpper(healthResponse.Status) != "UP" {
		return &ErrRemoteGatewayFailure{
			Method:     requestMethod,
			BaseUri:    healthPath,
			StatusCode: resp.StatusCode,
			Message:    fmt.Sprintf("status is %s in response from api", healthResponse.Status),
		}
	}

	return nil
}

type HealthResponse struct {
	Status string `json:"status"`
}
