package auth

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/api/converter"
	apiModel "github.com/lukmandev/nameless/gateway/internal/api/model"
	"github.com/lukmandev/nameless/gateway/internal/middleware"
)

func (q AuthQuery) GetMe(ctx context.Context) (*apiModel.GetMeResponse, error) {
	cookieAccess := middleware.CookieAccessFromCtx(ctx)

	profile, err := q.AuthService.GetMe(ctx, converter.ToGetMeInputFromAuthHandler(cookieAccess.AccessToken))
	if err != nil {
		return nil, err
	}
	return &apiModel.GetMeResponse{
		Profile: converter.ToProfileFromService(profile),
	}, nil
}
