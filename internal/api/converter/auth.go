package converter

import (
	"github.com/lukmandev/nameless/gateway/internal/api/model"
	serviceModel "github.com/lukmandev/nameless/gateway/internal/service/model"
)

func ToLoginInputFromAuthHandler(input model.LoginInput) *serviceModel.LoginInput {
	return &serviceModel.LoginInput{
		EmailOrUsername: input.EmailOrUsername,
		Password:        input.Password,
	}
}

func ToRegisterInputFromAuthHandler(input model.RegisterInput) *serviceModel.RegisterInput {
	return &serviceModel.RegisterInput{
		Email:    input.Email,
		Username: input.Username,
		Password: input.Password,
	}
}

func ToGetMeInputFromAuthHandler(accessToken string) *serviceModel.GetMeInput {
	return &serviceModel.GetMeInput{
		AccessToken: accessToken,
	}
}
