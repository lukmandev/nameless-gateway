package auth

import "github.com/lukmandev/nameless/gateway/internal/service"

type AuthMutation struct {
	authService service.AuthService
}

type AuthQuery struct {
	authService service.AuthService
}

func NewQuery() AuthQuery {
	return AuthQuery{}
}
