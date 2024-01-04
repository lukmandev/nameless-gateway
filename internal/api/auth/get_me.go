package auth

import (
	"context"

	apiModel "github.com/lukmandev/nameless/gateway/internal/api/model"
	"github.com/lukmandev/nameless/gateway/internal/model"
)

func (q AuthQuery) GetMe(ctx context.Context) (*apiModel.GetMeResponse, error) {
	q.authService.Login(ctx, &model.LoginInput{})
	return &apiModel.GetMeResponse{
		Profile: &apiModel.User{},
	}, nil
}
