package auth

import (
	"context"
	"errors"

	"github.com/lukmandev/nameless/gateway/internal/api/converter"
	"github.com/lukmandev/nameless/gateway/internal/api/model"
)

func (m AuthMutation) Login(ctx context.Context, input model.LoginInput) (*model.LoginResponse, error) {
	profile, _, accessToken, err := m.AuthService.Login(ctx, converter.ToLoginInputFromAuthApi(input))
	if err != nil {
		return nil, errors.New("SOME ERROR error")
	}
	return &model.LoginResponse{
		AccessToken: accessToken,
		Profile:     converter.ToProfileFromService(profile),
	}, nil
}
