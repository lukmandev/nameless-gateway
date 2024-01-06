package auth

import "github.com/lukmandev/nameless/gateway/internal/service"

type AuthMutation struct {
	AuthService service.AuthService
}

type AuthQuery struct {
	AuthService service.AuthService
}
