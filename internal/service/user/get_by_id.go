package user

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/service/model"
)

func (s *serv) GetByID(ctx context.Context, id int64) (*model.PublicProfile, error) {
	profile, err := s.userClient.GetByID(ctx, id)
	return profile, err
}
