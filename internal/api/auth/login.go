package auth

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/api/model"
)

func (m AuthMutation) Login(ctx context.Context, input model.LoginInput) (*model.LoginResponse, error) {
	return &model.LoginResponse{
		AccessToken: "",
		Profile:     &model.User{},
	}, nil
}
