package movie

import "github.com/lukmandev/nameless/gateway/internal/service"

type MovieMutation struct {
	MovieService service.MovieService
}

type MovieQuery struct {
	MovieService service.MovieService
	AuthService  service.AuthService
}
