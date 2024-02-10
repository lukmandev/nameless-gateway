package auth

import (
	movieDesc "github.com/lukmandev/nameless-movie/pkg/movie_v1"
	"github.com/lukmandev/nameless/gateway/internal/client"
)

type cli struct {
	movieClient movieDesc.MovieV1Client
}

func New(
	movieClient movieDesc.MovieV1Client,
) client.MovieServiceClient {
	return &cli{movieClient: movieClient}
}
