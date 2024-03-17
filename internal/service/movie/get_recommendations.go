package user

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/service/model"
)

func (s *serv) GetRecommendations(ctx context.Context, limit int32, userID int64) ([]*model.Movie, error) {
	movies, err := s.movieClient.GetRecommendations(ctx, limit, userID)
	return movies, err
}
