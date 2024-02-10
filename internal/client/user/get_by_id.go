package auth

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/client/converter"
	"github.com/lukmandev/nameless/gateway/internal/service/model"
)

func (c *cli) GetByID(ctx context.Context, id int64) (*model.PublicProfile, error) {
	response, err := c.userClient.GetById(ctx, converter.ToGetUserByIDFromService(id))
	if err != nil {
		return nil, err
	}
	return converter.ToPublicProfileFromUserDesc(response.GetUser()), nil
}
