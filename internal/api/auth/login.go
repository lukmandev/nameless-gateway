package auth

import (
	"context"
	"errors"

	"github.com/lukmandev/nameless/gateway/internal/api/converter"
	"github.com/lukmandev/nameless/gateway/internal/api/model"
	"github.com/lukmandev/nameless/gateway/internal/api/syserrors"
	"github.com/lukmandev/nameless/gateway/internal/middleware"
	serviceErrors "github.com/lukmandev/nameless/gateway/internal/service/syserrors"
)

func (m AuthMutation) Login(ctx context.Context, input model.LoginInput) (*model.LoginResponse, error) {
	cookieAccess := middleware.GetCookieAccessFromCtx(ctx)
	profile, refreshToken, accessToken, err := m.AuthService.Login(ctx, converter.ToLoginInputFromAuthHandler(input))
	if err != nil {
		if errors.Is(err, serviceErrors.ErrWrongCredentials) {
			return nil, syserrors.ErrWrongCredentials
		}
		return nil, syserrors.ErrUnknown
	}
	cookieAccess.SetRefreshToken(refreshToken)

	return &model.LoginResponse{
		AccessToken: accessToken,
		Profile:     converter.ToProfileFromService(profile),
	}, nil
}
