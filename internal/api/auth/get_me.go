package auth

import (
	"context"

	apiModel "github.com/lukmandev/nameless/gateway/internal/api/model"
)

func (q AuthQuery) GetMe(ctx context.Context) (*apiModel.GetMeResponse, error) {
	return &apiModel.GetMeResponse{
		Profile: &apiModel.Profile{},
	}, nil
}
