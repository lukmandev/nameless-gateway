package auth

import (
	"context"

	authDesc "github.com/lukmandev/nameless-auth/pkg/auth_v1"
	"github.com/lukmandev/nameless/gateway/internal/client/converter"
	"github.com/lukmandev/nameless/gateway/internal/service/model"
)

func (c *cli) Register(ctx context.Context, input *model.RegisterInput) (*model.Profile, string, string, error) {
	res, err := c.authClient.Register(ctx, &authDesc.RegisterRequest{
		Input: &authDesc.CreateUserInput{
			Email:    input.Email,
			Username: input.Username,
			Password: input.Password,
		},
	})
	if err != nil {
		return nil, "", "", err
	}

	return converter.ToProfileFromAuthDesc(res.GetProfile()), res.GetRefreshToken(), res.GetAccessToken(), nil
}
