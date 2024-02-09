package auth

import (
	userDesc "github.com/lukmandev/nameless-auth/pkg/user_v1"
	"github.com/lukmandev/nameless/gateway/internal/client"
)

type cli struct {
	userClient userDesc.UserV1Client
}

func New(
	userClient userDesc.UserV1Client,
) client.UserServiceClient {
	return &cli{userClient: userClient}
}
