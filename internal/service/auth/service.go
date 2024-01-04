package auth

import (
	"github.com/lukmandev/nameless/gateway/internal/client"
	"github.com/lukmandev/nameless/gateway/internal/service"
)

type serv struct {
	clients *client.ServiceClients
}

func NewService(clients *client.ServiceClients) service.AuthService {
	return &serv{
		clients: clients,
	}
}
