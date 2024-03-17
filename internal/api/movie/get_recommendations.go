package movie

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/api/converter"
	"github.com/lukmandev/nameless/gateway/internal/api/model"
	"github.com/lukmandev/nameless/gateway/internal/middleware"
)

func (m MovieQuery) GetMovieRecommendations(ctx context.Context, limit int) ([]*model.Movie, error) {
	profile, err := middleware.GetUser(ctx, m.AuthService)
	if err != nil {
		return nil, err
	}
	movies, err := m.MovieService.GetRecommendations(ctx, int32(limit), int64(profile.ID))
	if err != nil {
		return nil, err
	}
	return converter.ToMovieFromServiceList(movies), nil
}
