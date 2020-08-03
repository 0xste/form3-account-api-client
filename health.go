package accountapi_client

import (
	"context"
	"encoding/json"
	"net/http"
)

// hits the Health api of the account service
func (a *accountClient) GetHealth(ctx context.Context) (HealthResponse, error) {

	healthPath := a.baseUrl.String() + endpointHealth
	req, err := http.NewRequest(http.MethodGet, healthPath, nil)
	if err != nil {
		return HealthResponse{}, &ErrInvalidRequest{http.MethodGet, healthPath}
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return HealthResponse{}, &ErrRemoteGatewayFailure{
			http.MethodGet,
			healthPath,
			resp.StatusCode,
			"failed to connect to remote gateway",
		}
	}

	if resp.StatusCode != 200 {
		return HealthResponse{}, &ErrRemoteGatewayFailure{
			http.MethodGet,
			healthPath,
			resp.StatusCode,
			"invalid response code from api",
		}
	}

	var healthResponse HealthResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&healthResponse)
	if err != nil {
		return HealthResponse{}, &ErrRemoteGatewayFailure{
			Method:     http.MethodGet,
			BaseUri:    healthPath,
			StatusCode: http.StatusOK,
			Message:    "invalid response body from api",
		}
	}
	return healthResponse, nil
}

type HealthResponse struct {
	Status string `json:"status"`
}
