package auth

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/api/model"
)

func (m AuthQuery) GetMe(ctx context.Context) (*model.GetMeResponse, error) {
	return &model.GetMeResponse{
		Profile: &model.User{},
	}, nil
}
