package auth

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/service/model"
)

func (s *serv) Register(ctx context.Context, input *model.RegisterInput) (*model.Profile, string, string, error) {
	profile, refreshToken, accessToken, err := s.clients.AuthServiceClient.Register(ctx, input)
	if err != nil {
		return nil, "", "", err
	}
	return profile, refreshToken, accessToken, nil
}
