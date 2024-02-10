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

type TalentServiceClient interface {
	GetByID(ctx context.Context, id string) (*model.Talent, error)
	GetByIDs(ctx context.Context, ids []string) ([]*model.Talent, error)
}

type ServiceClients struct {
	AuthServiceClient   AuthServiceClient
	UserServiceClient   UserServiceClient
	TalentServiceClient TalentServiceClient
	MovieServiceClient  MovieServiceClient
}

func New(authClient AuthServiceClient, userClient UserServiceClient, movieClient MovieServiceClient, talentClient TalentServiceClient) *ServiceClients {
	return &ServiceClients{
		AuthServiceClient:   authClient,
		UserServiceClient:   userClient,
		TalentServiceClient: talentClient,
		MovieServiceClient:  movieClient,
	}
}
