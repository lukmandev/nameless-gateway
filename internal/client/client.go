package client

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/service/model"
)

type AuthServiceClient interface {
	Login(ctx context.Context, input *model.LoginInput) (*model.Profile, string, string, error)
	Register(ctx context.Context, input *model.RegisterInput) (*model.Profile, string, string, error)
	GetMe(ctx context.Context, input *model.GetMeInput) (*model.Profile, error)
}

type UserServiceClient interface {
	GetByID(ctx context.Context, id int64) (*model.PublicProfile, error)
}

type MovieServiceClient interface {
	GetByID(ctx context.Context, id string) (*model.Movie, error)
}

type ServiceClients struct {
	AuthServiceClient AuthServiceClient
	UserServiceClient UserServiceClient
}

func New(authClient AuthServiceClient, userClient UserServiceClient) *ServiceClients {
	return &ServiceClients{
		AuthServiceClient: authClient,
		UserServiceClient: userClient,
	}
}
