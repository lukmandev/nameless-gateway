package service

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/service/model"
)

type UserService interface {
	GetByID(ctx context.Context, id int64) (*model.PublicProfile, error)
}

type AuthService interface {
	Login(ctx context.Context, input *model.LoginInput) (*model.Profile, string, string, error)
	Register(ctx context.Context, input *model.RegisterInput) (*model.Profile, string, string, error)
	GetMe(ctx context.Context, input *model.GetMeInput) (*model.Profile, error)
}
