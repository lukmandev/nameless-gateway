package converter

import (
	apiModel "github.com/lukmandev/nameless/gateway/internal/api/model"
	"github.com/lukmandev/nameless/gateway/internal/model"
)

func ToLoginInputFromAuthApi(input apiModel.LoginInput) *model.LoginInput {
	return &model.LoginInput{
		EmailOrUsername: input.EmailOrUsername,
		Password:        input.Password,
	}
}

func ToProfileFromService(profile *model.Profile) *apiModel.Profile {
	return &apiModel.Profile{
		ID:               int(profile.ID),
		Username:         profile.Username,
		Email:            profile.Email,
		Verified:         profile.Verified,
		RegistrationDate: profile.RegistrationDate.String(),
	}
}
