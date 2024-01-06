package auth

import (
	"context"

	"github.com/lukmandev/nameless-auth/pkg/auth_v1"
	"github.com/lukmandev/nameless-auth/pkg/user_v1"
	"github.com/lukmandev/nameless/gateway/internal/model"
)

type AuthServiceClient interface {
	Login(ctx context.Context, input *model.LoginInput) (*model.Profile, string, string, error)
}

type client struct {
	authClient auth_v1.AuthV1Client
	userClient user_v1.UserV1Client
}

func New(
	authClient auth_v1.AuthV1Client,
	userClient user_v1.UserV1Client,
) *client {
	return &client{authClient: authClient, userClient: userClient}
}
