package auth

import (
	"context"

	authDesc "github.com/lukmandev/nameless-auth/pkg/auth_v1"
	"github.com/lukmandev/nameless/gateway/internal/client/converter"
	"github.com/lukmandev/nameless/gateway/internal/service/model"
	"github.com/lukmandev/nameless/gateway/internal/service/syserrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *cli) GetMe(ctx context.Context, input *model.GetMeInput) (*model.Profile, error) {
	res, err := c.authClient.GetMe(ctx, &authDesc.GetMeRequest{
		AccessToken: input.AccessToken,
	})
	if err != nil {
		status, ok := status.FromError(err)
		if ok {
			if status.Code() == codes.NotFound {
				return nil, syserrors.ErrUserNotFound
			}
		}
		return nil, syserrors.ErrUnknown
	}

	return converter.ToProfileFromAuthDesc(res.GetProfile()), nil
}
