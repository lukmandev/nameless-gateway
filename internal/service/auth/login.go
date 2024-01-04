package auth

import (
	"context"

	"github.com/lukmandev/nameless-auth/pkg/auth_v1"
	"github.com/lukmandev/nameless/gateway/internal/model"
)

func (s *serv) Login(ctx context.Context, input *model.LoginInput) (*model.Profile, string, string, error) {
	loginInput := &auth_v1.LoginInput{
		EmailOrUsername: input.EmailOrUsername,
		Password:        input.Password,
	}
	loginRequest := &auth_v1.LoginRequest{
		Input: loginInput,
	}
	loginResponse, err := s.clients.AuthClient.Login(ctx, loginRequest)

	if err != nil {
		return nil, "", "", nil
	}
	return &model.Profile{}, loginResponse.RefreshToken, loginResponse.AccessToken, nil
}
