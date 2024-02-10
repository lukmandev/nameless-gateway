package user

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/service/model"
)

func (s *serv) GetByID(ctx context.Context, id string) (*model.Movie, error) {
	profile, err := s.movieClient.GetByID(ctx, id)
	return profile, err
}
