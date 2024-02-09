package converter

import (
	commonDesc "github.com/lukmandev/nameless-auth/pkg/common_v1"
	userDesc "github.com/lukmandev/nameless-auth/pkg/user_v1"
	"github.com/lukmandev/nameless/gateway/internal/service/model"
)

func ToGetUserByIdFromService(id int64) *userDesc.GetRequest {
	return &userDesc.GetRequest{
		Id: id,
	}
}

func ToPublicProfileFromUserDesc(input *commonDesc.PublicProfile) *model.PublicProfile {
	var avatarURL *string
	if input.AvatarUrl != nil {
		avatarURL = &input.AvatarUrl.Value
	}
	return &model.PublicProfile{
		ID:               input.Id,
		Username:         input.Username,
		AvatarURL:        avatarURL,
		RegistrationDate: input.RegistrationDate.AsTime(),
	}
}
