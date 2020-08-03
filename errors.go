package accountapi_client

import "fmt"

const (
	errMsgInvalidBaseUri       string = "baseuri %s is invalid"
	errMsgRemoteGatewayFailure string = "remote gateway failure for %s request %s with status %d"
	errMsgInvalidRequest       string = "remote gateway failure for %s request %s"
)

type ErrInvalidClientBaseUri struct {
	BaseUri string
}

func (e ErrInvalidClientBaseUri) Error() string {
	return fmt.Sprintf(errMsgInvalidBaseUri, e.BaseUri)
}

type ErrRemoteGatewayFailure struct {
	Method     string
	BaseUri    string
	StatusCode int
	Message    string
}

func (e ErrRemoteGatewayFailure) Error() string {
	return fmt.Sprintf(errMsgRemoteGatewayFailure, e.Method, e.BaseUri, e.StatusCode)
}

type ErrInvalidRequest struct {
	Method  string
	BaseUri string
}

func (e ErrInvalidRequest) Error() string {
	return fmt.Sprintf(errMsgInvalidRequest, e.Method, e.BaseUri)
}
