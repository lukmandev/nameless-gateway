package user

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/api/converter"
	"github.com/lukmandev/nameless/gateway/internal/api/model"
)

func (q UserQuery) GetByID(ctx context.Context, id int) (*model.PublicProfile, error) {
	profile, err := q.UserService.GetByID(ctx, int64(id))
	if err != nil {
		return nil, err
	}
	return converter.ToPublicProfileFromService(profile), nil
}
