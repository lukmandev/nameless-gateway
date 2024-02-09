package auth

import (
	"github.com/lukmandev/nameless/gateway/internal/client"
	"github.com/lukmandev/nameless/gateway/internal/service"
)

type serv struct {
	authClient client.AuthServiceClient
}

func NewService(authClient client.AuthServiceClient) service.AuthService {
	return &serv{
		authClient: authClient,
	}
}
