package auth

import (
	"context"

	authDesc "github.com/lukmandev/nameless-auth/pkg/auth_v1"
	userDesc "github.com/lukmandev/nameless-auth/pkg/user_v1"
	"github.com/lukmandev/nameless/gateway/internal/service/model"
)

type AuthServiceClient interface {
	Login(ctx context.Context, input *model.LoginInput) (*model.Profile, string, string, error)
	Register(ctx context.Context, input *model.RegisterInput) (*model.Profile, string, string, error)
}

type client struct {
	authClient authDesc.AuthV1Client
	userClient userDesc.UserV1Client
}

func New(
	authClient authDesc.AuthV1Client,
	userClient userDesc.UserV1Client,
) *client {
	return &client{authClient: authClient, userClient: userClient}
}
