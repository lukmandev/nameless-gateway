package client

import (
	"github.com/lukmandev/nameless-auth/pkg/auth_v1"
	"github.com/lukmandev/nameless-auth/pkg/user_v1"
)

type ServiceClients struct {
	AuthClient auth_v1.AuthV1Client
	UserClient user_v1.UserV1Client
}
