package auth

import (
	"context"
	"errors"

	"github.com/lukmandev/nameless/gateway/internal/api/converter"
	apiModel "github.com/lukmandev/nameless/gateway/internal/api/model"
	"github.com/lukmandev/nameless/gateway/internal/api/syserrors"
	"github.com/lukmandev/nameless/gateway/internal/middleware"
	serviceErrors "github.com/lukmandev/nameless/gateway/internal/service/syserrors"
)

func (q AuthQuery) GetMe(ctx context.Context) (*apiModel.GetMeResponse, error) {
	cookieAccess := middleware.GetCookieAccessFromCtx(ctx)

	profile, err := q.AuthService.GetMe(ctx, converter.ToGetMeInputFromAuthHandler(cookieAccess.AccessToken))
	if err != nil {
		if errors.Is(err, serviceErrors.ErrUserNotFound) {
			return nil, syserrors.ErrTokenExpired
		}
		return nil, syserrors.ErrUnknown
	}
	return &apiModel.GetMeResponse{
		Profile: converter.ToProfileFromService(profile),
	}, nil
}
