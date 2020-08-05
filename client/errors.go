package client

import (
	"fmt"
	"form3-accountapi-client/uuid"
)

const (
	errMsgInvalidBaseUri       string = "baseuri %s is invalid"
	errMsgRemoteGatewayFailure string = "remote gateway failure, %s for %s request %s with status %d"
	errMsgInvalidRequest       string = "remote gateway failure %s for %s request %s"
	errMsgAccountNotFound      string = "account-id %s and version %d not found"
	errMsgDuplicateAccount      string = "duplicate account, cannot create with id %s"
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
	return fmt.Sprintf(errMsgRemoteGatewayFailure, e.Message, e.Method, e.BaseUri, e.StatusCode)
}

type ErrInvalidRequest struct {
	Method  string
	BaseUri string
	ErrMsg string
}

func (e ErrInvalidRequest) Error() string {
	return fmt.Sprintf(errMsgInvalidRequest, e.ErrMsg, e.Method, e.BaseUri)
}

type ErrAccountNotFound struct {
	AccountId uuid.UUID
	AccountVersion int64
}

func (e ErrAccountNotFound) Error() string {
	return fmt.Sprintf(errMsgAccountNotFound, e.AccountId, e.AccountVersion)
}

type ErrDuplicateAccount struct {
	Account *Account
}

func (e ErrDuplicateAccount) Error() string {
	return fmt.Sprintf(errMsgDuplicateAccount, e.Account.Id)
}

