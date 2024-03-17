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

func (c *cli) Register(ctx context.Context, input *model.RegisterInput) (*model.Profile, string, string, error) {
	res, err := c.authClient.Register(ctx, &authDesc.RegisterRequest{
		Input: &authDesc.CreateUserInput{
			Email:    input.Email,
			Username: input.Username,
			Password: input.Password,
		},
	})
	if err != nil {
		status, ok := status.FromError(err)
		if ok {
			if status.Code() == codes.AlreadyExists {
				return nil, "", "", syserrors.ErrEmailOrUsernameAlreadyInUse
			}
		}
		return nil, "", "", syserrors.ErrUnknown
	}

	return converter.ToProfileFromAuthDesc(res.GetProfile()), res.GetRefreshToken(), res.GetAccessToken(), nil
}
