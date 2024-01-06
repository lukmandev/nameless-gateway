package client

import (
	"github.com/lukmandev/nameless/gateway/internal/client/auth"
)

type ServiceClients struct {
	AuthServiceClient auth.AuthServiceClient
}

func New(authClient auth.AuthServiceClient) *ServiceClients {
	return &ServiceClients{
		AuthServiceClient: authClient,
	}
}
