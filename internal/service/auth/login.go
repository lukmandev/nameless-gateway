package auth

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/service/model"
)

func (s *serv) Login(ctx context.Context, input *model.LoginInput) (*model.Profile, string, string, error) {
	profile, refreshToken, accessToken, err := s.clients.AuthServiceClient.Login(ctx, input)
	if err != nil {
		return nil, "", "", err
	}
	return profile, refreshToken, accessToken, nil
}
