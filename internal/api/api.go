package api

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/api/auth"
	"github.com/lukmandev/nameless/gateway/internal/api/model"
	"github.com/lukmandev/nameless/gateway/internal/api/movie"
	"github.com/lukmandev/nameless/gateway/internal/api/user"
	"github.com/lukmandev/nameless/gateway/internal/service"
)

type Resolver struct {
	authService service.AuthService
	userService service.UserService
}

type directorResolver struct {
	*Resolver
}

func (r *Resolver) Director() DirectorResolver {
	return &directorResolver{r}
}

func (r *directorResolver) Talent(ctx context.Context, obj *model.Director) (*model.Talent, error) {
	return nil, nil
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
	}
}

type Mutation struct {
	auth.AuthMutation
	movie.MovieMutation
}
type Query struct {
	auth.AuthQuery
	user.UserQuery
	movie.MovieQuery
}

func NewResolver(
	authService service.AuthService,
	userService service.UserService,
) Resolver {
	return Resolver{
		authService: authService,
		userService: userService,
	}
}
