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

func (m AuthMutation) Register(ctx context.Context, input model.RegisterInput) (*model.RegisterResponse, error) {
	cookieAccess := middleware.GetCookieAccessFromCtx(ctx)
	profile, refreshToken, accessToken, err := m.AuthService.Register(ctx, converter.ToRegisterInputFromAuthHandler(input))
	if err != nil {
		if errors.Is(err, serviceErrors.ErrEmailOrUsernameAlreadyInUse) {
			return nil, syserrors.ErrWrongCredentials
		}
		return nil, syserrors.ErrUnknown
	}
	cookieAccess.SetRefreshToken(refreshToken)

	return &model.RegisterResponse{
		AccessToken: accessToken,
		Profile:     converter.ToProfileFromService(profile),
	}, nil
}
