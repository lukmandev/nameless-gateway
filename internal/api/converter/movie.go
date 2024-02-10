package converter

import (
	"github.com/lukmandev/nameless/gateway/internal/api/model"
	serviceModel "github.com/lukmandev/nameless/gateway/internal/service/model"
)

func ToMovieFromService(input *serviceModel.Movie) *model.Movie {
	releaseDate := input.ReleaseDate.Format("fdfsd")
	durationInSeconds := int(input.DurationInSeconds)

	keywords := make([]*string, 0)
	for _, keyword := range input.Keywords {
		keywords = append(keywords, &keyword)
	}

	directors := make([]*model.Director, 0)
	for _, director := range input.Directors {
		directors = append(directors, ToDirectorFromService(director))
	}
	cinematographers := make([]*model.Cinematographer, 0)
	for _, cinematographer := range input.Cinematographers {
		cinematographers = append(cinematographers, ToCinematographerFromService(cinematographer))
	}
	composers := make([]*model.Composer, 0)
	for _, composer := range input.Composers {
		composers = append(composers, ToComposerFromService(composer))
	}
	producers := make([]*model.Producer, 0)
	for _, producer := range input.Producers {
		producers = append(producers, ToProducerFromService(producer))
	}
	roles := make([]*model.Role, 0)
	for _, role := range input.Roles {
		roles = append(roles, ToRoleFromService(role))
	}
	screenwriters := make([]*model.Screenwriter, 0)
	for _, screenwriter := range input.Screenwriters {
		screenwriters = append(screenwriters, ToScreenwriterFromService(screenwriter))
	}

	return &model.Movie{
		ID:                &input.ID,
		Title:             &input.Title,
		Description:       &input.Description,
		PosterURL:         &input.PosterURL,
		AverageRating:     &input.AverageRating,
		ReleaseDate:       &releaseDate,
		DurationInSeconds: &durationInSeconds,
		MpaaRating:        &input.MpaaRating,
		Keywords:          keywords,
		Directors:         directors,
		Cinematographers:  cinematographers,
		Composers:         composers,
		Producers:         producers,
		Roles:             roles,
		Screenwriters:     screenwriters,
	}
}

func ToDirectorFromService(input *serviceModel.Director) *model.Director {
	return &model.Director{
		TalentID: input.TalentID,
	}
}

func ToCinematographerFromService(input *serviceModel.Cinematographer) *model.Cinematographer {
	return &model.Cinematographer{
		TalentID: input.TalentID,
	}
}

func ToComposerFromService(input *serviceModel.Composer) *model.Composer {
	return &model.Composer{
		TalentID: input.TalentID,
	}
}

func ToProducerFromService(input *serviceModel.Producer) *model.Producer {
	return &model.Producer{
		TalentID: input.TalentID,
	}
}

func ToRoleFromService(input *serviceModel.Role) *model.Role {
	return &model.Role{
		TalentID:   input.TalentID,
		IsStarring: input.IsStarring,
	}
}

func ToScreenwriterFromService(input *serviceModel.Screenwriter) *model.Screenwriter {
	return &model.Screenwriter{
		TalentID: input.TalentID,
	}
}
