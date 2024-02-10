package auth

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/client/converter"
	"github.com/lukmandev/nameless/gateway/internal/service/model"
)

func (c *cli) GetByID(ctx context.Context, id string) (*model.Movie, error) {
	response, err := c.movieClient.GetByID(ctx, converter.ToGetMovieByIDFromService(id))
	if err != nil {
		return nil, err
	}
	return converter.ToMovieFromMovieDesc(response.GetMovie()), nil
}
