package auth

import (
	"context"

	"github.com/lukmandev/nameless-auth/pkg/auth_v1"
	"github.com/lukmandev/nameless/gateway/internal/client/auth/converter"
	"github.com/lukmandev/nameless/gateway/internal/model"
)

func (c *client) Login(ctx context.Context, input *model.LoginInput) (*model.Profile, string, string, error) {
	res, err := c.authClient.Login(ctx, &auth_v1.LoginRequest{
		Input: &auth_v1.LoginInput{
			EmailOrUsername: input.EmailOrUsername,
			Password:        input.Password,
		},
	})
	if err != nil {
		return nil, "", "", err
	}

	return converter.ToProfileFromAuthDesc(res.GetProfile()), res.GetRefreshToken(), res.GetAccessToken(), nil
}
