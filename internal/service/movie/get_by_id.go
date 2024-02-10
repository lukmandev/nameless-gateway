package user

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/service/model"
)

func (s *serv) GetByID(ctx context.Context, id string, userID int64) (*model.Movie, error) {
	movie, err := s.movieClient.GetByID(ctx, id, userID)
	return movie, err
}
