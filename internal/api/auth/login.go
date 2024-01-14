package auth

import (
	"context"
	"fmt"

	"github.com/lukmandev/nameless/gateway/internal/api/model"
	"github.com/lukmandev/nameless/gateway/internal/converter"
	// "github.com/lukmandev/nameless/gateway/internal/converter"
)

func (m AuthMutation) Login(ctx context.Context, input model.LoginInput) (*model.LoginResponse, error) {
	profile, _, accessToken, err := m.AuthService.Login(ctx, converter.ToLoginInputFromAuthApi(input))
	fmt.Println("Error", err)
	return &model.LoginResponse{
		AccessToken: accessToken,
		Profile:     converter.ToProfileFromService(profile),
	}, nil
}
