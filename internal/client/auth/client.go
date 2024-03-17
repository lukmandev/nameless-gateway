package auth

import (
	authDesc "github.com/lukmandev/nameless-auth/pkg/auth_v1"
	userDesc "github.com/lukmandev/nameless-auth/pkg/user_v1"
	"github.com/lukmandev/nameless/gateway/internal/client"
)

type cli struct {
	authClient authDesc.AuthV1Client
	userClient userDesc.UserV1Client
}

func New(
	authClient authDesc.AuthV1Client,
) client.AuthServiceClient {
	return &cli{authClient: authClient}
}
