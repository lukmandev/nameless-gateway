package auth

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/client/converter"
	"github.com/lukmandev/nameless/gateway/internal/service/model"
	"github.com/lukmandev/nameless/gateway/internal/service/syserrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *cli) GetByID(ctx context.Context, id int64) (*model.PublicProfile, error) {
	response, err := c.userClient.GetById(ctx, converter.ToGetUserByIDFromService(id))
	if err != nil {
		status, ok := status.FromError(err)
		if ok {
			if status.Code() == codes.NotFound {
				return nil, syserrors.ErrUserNotFound
			}
		}
		return nil, syserrors.ErrUnknown
	}
	return converter.ToPublicProfileFromUserDesc(response.GetUser()), nil
}
