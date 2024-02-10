package movie

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/api/model"
)

func (m MovieQuery) GetMovieByID(ctx context.Context, id string) (*model.Movie, error) {
	// movie, err := m.MovieService.GetByID(ctx, id)
	return &model.Movie{}, nil
}
