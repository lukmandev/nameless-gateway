package converter

import (
	talentDesc "github.com/lukmandev/nameless-movie/pkg/talent_v1"
	"github.com/lukmandev/nameless/gateway/internal/service/model"
)

func ToGetTalentByIDFromService(id string) *talentDesc.GetByIDRequest {
	return &talentDesc.GetByIDRequest{
		Id: id,
	}
}

func ToGetTalentsByIDsFromService(ids []string) *talentDesc.GetByIDsRequest {
	return &talentDesc.GetByIDsRequest{
		Ids: ids,
	}
}

func ToTalentFromTalentDesc(input *talentDesc.Talent) *model.Talent {
	career := make([]*model.TalentCareer, len(input.Career))
	for _, car := range input.Career {
		career = append(career, &model.TalentCareer{
			Id: car.Id,
		})
	}
	return &model.Talent{
		Id:        input.Id,
		Name:      input.Name,
		BirthDate: input.BirthDate.AsTime(),
		PhotoUrl:  input.PhotoUrl,
		Career:    career,
	}
}

func ToTalentFromTalentDescList(input []*talentDesc.Talent) []*model.Talent {
	result := make([]*model.Talent, 0)
	for i := range input {
		result = append(result, ToTalentFromTalentDesc(input[i]))
	}

	return result
}
