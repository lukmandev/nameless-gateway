package converter

import (
	"github.com/lukmandev/nameless/gateway/internal/service/model"

	authCommonDesc "github.com/lukmandev/nameless-auth/pkg/common_v1"
)

func ToProfileFromAuthDesc(input *authCommonDesc.Profile) *model.Profile {
	return &model.Profile{
		ID:               input.Id,
		Username:         input.Username,
		Email:            input.Email,
		Verified:         input.Verified,
		AvatarURL:        &input.AvatarUrl.Value,
		RegistrationDate: input.RegistrationDate.AsTime(),
	}
}
