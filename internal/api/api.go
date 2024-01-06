package api

import (
	"github.com/lukmandev/nameless/gateway/internal/api/auth"
	"github.com/lukmandev/nameless/gateway/internal/api/user"
	"github.com/lukmandev/nameless/gateway/internal/service"
)

type Resolver struct {
	authService service.AuthService
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
		UserQuery: user.UserQuery{},
	}
}

type Mutation struct {
	auth.AuthMutation
}
type Query struct {
	auth.AuthQuery
	user.UserQuery
}

func NewResolver(
	authService service.AuthService,
) Resolver {
	return Resolver{
		authService: authService,
	}
}
