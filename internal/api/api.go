package api

import (
	"github.com/lukmandev/nameless/gateway/internal/api/auth"
	"github.com/lukmandev/nameless/gateway/internal/api/movie"
	"github.com/lukmandev/nameless/gateway/internal/api/talent"
	"github.com/lukmandev/nameless/gateway/internal/api/user"
	"github.com/lukmandev/nameless/gateway/internal/service"
)

type Resolver struct {
	authService   service.AuthService
	userService   service.UserService
	talentService service.TalentService
	movieService  service.MovieService
}

func (r *Resolver) Mutation() MutationResolver {
	return Mutation{
		AuthMutation: auth.AuthMutation{
			AuthService: r.authService,
		},
	}
}

func (r *Resolver) Query() QueryResolver {
	return Query{
		AuthQuery: auth.AuthQuery{
			AuthService: r.authService,
		},
		UserQuery: user.UserQuery{
			UserService: r.userService,
		},
		TalentQuery: talent.TalentQuery{
			TalentService: r.talentService,
		},
		MovieQuery: movie.MovieQuery{
			MovieService: r.movieService,
		},
	}
}

type Mutation struct {
	auth.AuthMutation
	movie.MovieMutation
	talent.TalentMutation
}
type Query struct {
	auth.AuthQuery
	user.UserQuery
	talent.TalentQuery
	movie.MovieQuery
}

func NewResolver(
	authService service.AuthService,
	userService service.UserService,
	movieService service.MovieService,
	talentService service.TalentService,
) Resolver {
	return Resolver{
		authService:   authService,
		userService:   userService,
		movieService:  movieService,
		talentService: talentService,
	}
}
