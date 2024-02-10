package converter

import (
	movieDesc "github.com/lukmandev/nameless-movie/pkg/movie_v1"
	"github.com/lukmandev/nameless/gateway/internal/service/model"
)

func ToGetMovieByIDFromService(id string) *movieDesc.GetByIDRequest {
	return &movieDesc.GetByIDRequest{
		Id: id,
	}
}

func ToMovieFromMovieDesc(protoMovie *movieDesc.Movie) *model.Movie {
	serviceMovie := &model.Movie{
		ID:                protoMovie.GetId(),
		Title:             protoMovie.GetTitle(),
		Description:       protoMovie.GetDescription(),
		PosterURL:         protoMovie.GetPosterUrl(),
		AverageRating:     protoMovie.GetAverageRating(),
		ReleaseDate:       protoMovie.GetReleaseDate().AsTime(),
		DurationInSeconds: int(protoMovie.GetDurationInSeconds()),
		MpaaRating:        protoMovie.GetMpaaRating(),
	}

	// Convert Keywords
	serviceMovie.Keywords = make([]string, len(protoMovie.GetKeywords()))
	for i, kw := range protoMovie.GetKeywords() {
		serviceMovie.Keywords[i] = kw
	}

	// Convert Directors, Screenwriters, Producers, Cinematographers, Composers, Roles
	serviceMovie.Directors = ToDirectorsFromMovieDesc(protoMovie.GetDirectors())
	serviceMovie.Screenwriters = ToScreenwritersFromMovieDesc(protoMovie.GetScreenwriters())
	serviceMovie.Producers = ToProducersFromMovieDesc(protoMovie.GetProducers())
	serviceMovie.Cinematographers = ToCinematographersFromMovieDesc(protoMovie.GetCinematographers())
	serviceMovie.Composers = ToComposersFromMovieDesc(protoMovie.GetComposers())
	serviceMovie.Roles = ToRolesFromMovieDesc(protoMovie.GetRoles())

	return serviceMovie
}

func ToDirectorFromMovieDesc(protoDirector *movieDesc.Director) *model.Director {
	if protoDirector == nil {
		return nil
	}

	return &model.Director{
		TalentID: protoDirector.GetTalentId(),
	}
}

func ToDirectorsFromMovieDesc(protoDirectors []*movieDesc.Director) []*model.Director {
	serviceDirectors := make([]*model.Director, len(protoDirectors))
	for i, protoDirector := range protoDirectors {
		serviceDirectors[i] = ToDirectorFromMovieDesc(protoDirector)
	}
	return serviceDirectors
}

func ToScreenwriterFromMovieDesc(protoScreenwriter *movieDesc.Screenwriter) *model.Screenwriter {
	if protoScreenwriter == nil {
		return nil
	}

	return &model.Screenwriter{
		TalentID: protoScreenwriter.GetTalentId(),
	}
}

func ToScreenwritersFromMovieDesc(protoScreenwriters []*movieDesc.Screenwriter) []*model.Screenwriter {
	serviceScreenwriters := make([]*model.Screenwriter, len(protoScreenwriters))
	for i, protoScreenwriter := range protoScreenwriters {
		serviceScreenwriters[i] = ToScreenwriterFromMovieDesc(protoScreenwriter)
	}
	return serviceScreenwriters
}

func ToProducerFromMovieDesc(protoProducer *movieDesc.Producer) *model.Producer {
	if protoProducer == nil {
		return nil
	}

	return &model.Producer{
		TalentID: protoProducer.GetTalentId(),
	}
}

func ToProducersFromMovieDesc(protoProducers []*movieDesc.Producer) []*model.Producer {
	serviceProducers := make([]*model.Producer, len(protoProducers))
	for i, protoProducer := range protoProducers {
		serviceProducers[i] = ToProducerFromMovieDesc(protoProducer)
	}
	return serviceProducers
}

func ToCinematographerFromMovieDesc(input *movieDesc.Cinematographer) *model.Cinematographer {
	if input == nil {
		return nil
	}

	return &model.Cinematographer{
		TalentID: input.GetTalentId(),
	}
}
func ToCinematographersFromMovieDesc(input []*movieDesc.Cinematographer) []*model.Cinematographer {
	response := make([]*model.Cinematographer, len(input))
	for i, protoProducer := range input {
		response[i] = ToCinematographerFromMovieDesc(protoProducer)
	}
	return response
}

func ToComposerFromMovieDesc(input *movieDesc.Composer) *model.Composer {
	if input == nil {
		return nil
	}

	return &model.Composer{
		TalentID: input.GetTalentId(),
	}
}

func ToComposersFromMovieDesc(input []*movieDesc.Composer) []*model.Composer {
	response := make([]*model.Composer, len(input))
	for i, protoProducer := range input {
		response[i] = ToComposerFromMovieDesc(protoProducer)
	}
	return response
}

func ToRoleFromMovieDesc(input *movieDesc.Role) *model.Role {
	if input == nil {
		return nil
	}

	return &model.Role{
		TalentID:   input.GetTalentId(),
		IsStarring: input.GetIsStarring(),
	}
}

func ToRolesFromMovieDesc(input []*movieDesc.Role) []*model.Role {
	response := make([]*model.Role, len(input))
	for i, protoProducer := range input {
		response[i] = ToRoleFromMovieDesc(protoProducer)
	}
	return response
}
