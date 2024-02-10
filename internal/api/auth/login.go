package auth

import (
	"context"
	"errors"

	"github.com/lukmandev/nameless/gateway/internal/api/converter"
	"github.com/lukmandev/nameless/gateway/internal/api/model"
	"github.com/lukmandev/nameless/gateway/internal/middleware"
)

func (m AuthMutation) Login(ctx context.Context, input model.LoginInput) (*model.LoginResponse, error) {
	cookieAccess := middleware.GetCookieAccessFromCtx(ctx)
	profile, refreshToken, accessToken, err := m.AuthService.Login(ctx, converter.ToLoginInputFromAuthHandler(input))
	if err != nil {
		return nil, errors.New("SOME ERROR error")
	}
	cookieAccess.SetRefreshToken(refreshToken)

	return &model.LoginResponse{
		AccessToken: accessToken,
		Profile:     converter.ToProfileFromService(profile),
	}, nil
}
