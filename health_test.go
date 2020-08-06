package main

import (
	"context"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func Test_accountClient_GetHealth_up(t *testing.T) {
	defer gock.Off()
	gock.New("http://server.com").
		Get("/v1/health").
		Reply(200).
		JSON(map[string]string{"status": "up"})

	accountClient, err := NewAccountClient("http", "server.com", 80, 1)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}

	err = accountClient.Health(context.TODO())
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
}

func Test_accountClient_GetHealth_down(t *testing.T) {
	defer gock.Off()
	gock.New("http://server.com").
		Get("/v1/health").
		Reply(200).
		JSON(map[string]string{"status": "down"})

	accountClient, err := NewAccountClient("http", "server.com", 80, 1)
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}

	err = accountClient.Health(context.TODO())
	if _, ok := err.(ErrRemoteGatewayFailure); ok {
		t.Errorf("expected remote gateway failure but was  %v", err)
	}
}
