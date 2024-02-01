package converter

import (
	"github.com/lukmandev/nameless/gateway/internal/service/model"

	authCommonDesc "github.com/lukmandev/nameless-auth/pkg/common_v1"
)

func ToProfileFromAuthDesc(input *authCommonDesc.Profile) *model.Profile {
	var avatarURL *string
	if input.AvatarUrl != nil {
		avatarURL = &input.AvatarUrl.Value
	}
	return &model.Profile{
		ID:               input.Id,
		Username:         input.Username,
		Email:            input.Email,
		Verified:         input.Verified,
		AvatarURL:        avatarURL,
		RegistrationDate: input.RegistrationDate.AsTime(),
	}
}
