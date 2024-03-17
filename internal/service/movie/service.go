package user

import (
	"github.com/lukmandev/nameless/gateway/internal/client"
	"github.com/lukmandev/nameless/gateway/internal/service"
)

type serv struct {
	movieClient client.MovieServiceClient
}

func NewService(movieClient client.MovieServiceClient) service.MovieService {
	return &serv{
		movieClient: movieClient,
	}
}
