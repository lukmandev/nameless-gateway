package service

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/model"
)

type UserService interface {
}

type AuthService interface {
	Login(ctx context.Context, input *model.LoginInput) (*model.Profile, string, string, error)
}
