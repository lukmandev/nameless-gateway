package movie

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/api/converter"
	"github.com/lukmandev/nameless/gateway/internal/api/model"
)

func (m MovieQuery) GetMovieByID(ctx context.Context, id string) (*model.Movie, error) {
	movie, err := m.MovieService.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return converter.ToMovieFromService(movie), nil
}
