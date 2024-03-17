package movie

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/api/converter"
	"github.com/lukmandev/nameless/gateway/internal/api/model"
	"github.com/lukmandev/nameless/gateway/internal/middleware"
)

func (m MovieQuery) GetMovieByID(ctx context.Context, id string) (*model.Movie, error) {
	profileID := 0

	profile, err := middleware.GetUser(ctx, m.AuthService)
	if nil == err {
		profileID = profile.ID
	}
	movie, err := m.MovieService.GetByID(ctx, id, int64(profileID))
	if err != nil {
		return nil, err
	}
	return converter.ToMovieFromService(movie), nil
}
