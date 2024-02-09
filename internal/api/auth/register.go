package auth

import (
	"context"
	"errors"

	"github.com/lukmandev/nameless/gateway/internal/api/converter"
	"github.com/lukmandev/nameless/gateway/internal/api/model"
	"github.com/lukmandev/nameless/gateway/internal/middleware"
)

func (m AuthMutation) Register(ctx context.Context, input model.RegisterInput) (*model.RegisterResponse, error) {
	cookieAccess := middleware.CookieAccessFromCtx(ctx)
	profile, refreshToken, accessToken, err := m.AuthService.Register(ctx, converter.ToRegisterInputFromAuthHandler(input))
	if err != nil {
		return nil, errors.New("SOME ERROR error")
	}
	cookieAccess.SetRefreshToken(refreshToken)

	return &model.RegisterResponse{
		AccessToken: accessToken,
		Profile:     converter.ToProfileFromService(profile),
	}, nil
}
