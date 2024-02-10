package converter

import (
	"github.com/lukmandev/nameless/gateway/internal/api/model"
	serviceModel "github.com/lukmandev/nameless/gateway/internal/service/model"
)

func ToTalentFromService(input *serviceModel.Talent) *model.Talent {
	return &model.Talent{
		ID:   input.Id,
		Name: input.Name,
	}
}

func ToTalentFromServiceList(input []*serviceModel.Talent) []*model.Talent {
	result := make([]*model.Talent, 0)
	for i := range input {
		result = append(result, ToTalentFromService(input[i]))
	}

	return result
}
