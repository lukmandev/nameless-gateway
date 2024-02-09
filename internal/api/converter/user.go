package converter

import (
	"github.com/lukmandev/nameless/gateway/internal/api/model"
	serviceModel "github.com/lukmandev/nameless/gateway/internal/service/model"
)

func ToProfileFromService(profile *serviceModel.Profile) *model.Profile {
	return &model.Profile{
		ID:               int(profile.ID),
		Username:         profile.Username,
		Email:            profile.Email,
		Verified:         profile.Verified,
		AvatarURL:        profile.AvatarURL,
		RegistrationDate: profile.RegistrationDate.String(),
	}
}

func ToPublicProfileFromService(profile *serviceModel.PublicProfile) *model.PublicProfile {
	return &model.PublicProfile{
		ID:               int(profile.ID),
		Username:         profile.Username,
		AvatarURL:        profile.AvatarURL,
		RegistrationDate: profile.RegistrationDate.String(),
	}
}
