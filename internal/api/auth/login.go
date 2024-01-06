package auth

import (
	"context"
	"fmt"

	"github.com/lukmandev/nameless/gateway/internal/api/model"
	"github.com/lukmandev/nameless/gateway/internal/converter"
	// "github.com/lukmandev/nameless/gateway/internal/converter"
)

func (m AuthMutation) Login(ctx context.Context, input model.LoginInput) (*model.LoginResponse, error) {
	fmt.Println("BEFORE RESPONSE", input, converter.ToLoginInputFromAuthApi(input), m.AuthService)
	profile, _, accessToken, err := m.AuthService.Login(ctx, converter.ToLoginInputFromAuthApi(input))
	fmt.Println("AFTER RESPOMSE", err)
	return &model.LoginResponse{
		AccessToken: accessToken,
		Profile:     converter.ToProfileFromService(profile),
	}, nil
}
