package movie

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/api/model"
)

func (m MovieMutation) CreateMovie(ctx context.Context) (*model.CreateMovieResponse, error) {
	return &model.CreateMovieResponse{}, nil
}
