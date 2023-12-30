package user

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/api/model"
)

func (m UserQuery) GetByID(ctx context.Context) (*model.User, error) {
	return &model.User{}, nil
}
