package user

import (
	"github.com/lukmandev/nameless/gateway/internal/client"
	"github.com/lukmandev/nameless/gateway/internal/service"
)

type serv struct {
	userClient client.UserServiceClient
}

func NewService(userClient client.UserServiceClient) service.UserService {
	return &serv{
		userClient: userClient,
	}
}
