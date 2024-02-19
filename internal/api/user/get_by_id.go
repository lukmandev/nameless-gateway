package user

import (
	"context"
	"errors"

	"github.com/lukmandev/nameless/gateway/internal/api/converter"
	"github.com/lukmandev/nameless/gateway/internal/api/model"
	"github.com/lukmandev/nameless/gateway/internal/api/syserrors"
	serviceErrors "github.com/lukmandev/nameless/gateway/internal/service/syserrors"
)

func (q UserQuery) GetUserByID(ctx context.Context, id int) (*model.PublicProfile, error) {
	profile, err := q.UserService.GetByID(ctx, int64(id))
	if err != nil {
		if errors.Is(err, serviceErrors.ErrUserNotFound) {
			return nil, syserrors.ErrUserNotFound
		}
		return nil, syserrors.ErrUnknown
	}
	return converter.ToPublicProfileFromService(profile), nil
}
