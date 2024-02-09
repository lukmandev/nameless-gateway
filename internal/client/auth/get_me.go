package auth

import (
	"context"

	authDesc "github.com/lukmandev/nameless-auth/pkg/auth_v1"
	"github.com/lukmandev/nameless/gateway/internal/client/converter"
	"github.com/lukmandev/nameless/gateway/internal/service/model"
)

func (c *cli) GetMe(ctx context.Context, input *model.GetMeInput) (*model.Profile, error) {
	res, err := c.authClient.GetMe(ctx, &authDesc.GetMeRequest{
		AccessToken: input.AccessToken,
	})
	if err != nil {
		return nil, err
	}

	return converter.ToProfileFromAuthDesc(res.GetProfile()), nil
}
