package api

import (
	"github.com/lukmandev/nameless/gateway/internal/api/auth"
	"github.com/lukmandev/nameless/gateway/internal/api/user"
)

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return Mutation{}
}

func (r *Resolver) Query() QueryResolver {
	return Query{}
}

type Mutation struct {
	auth.AuthMutation
}
type Query struct {
	auth.AuthQuery
	user.UserQuery
}

func NewResolver() Resolver {
	return Resolver{}
}
