package auth

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/service/model"
)

func (s *serv) GetMe(ctx context.Context, input *model.GetMeInput) (*model.Profile, error) {
	profile, err := s.authClient.GetMe(ctx, input)
	if err != nil {
		return nil, err
	}
	return profile, nil
}
