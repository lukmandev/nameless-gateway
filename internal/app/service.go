package app

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/service"
	authService "github.com/lukmandev/nameless/gateway/internal/service/auth"
	movieService "github.com/lukmandev/nameless/gateway/internal/service/movie"
	talentService "github.com/lukmandev/nameless/gateway/internal/service/talent"
	userService "github.com/lukmandev/nameless/gateway/internal/service/user"
)

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = authService.NewService(s.ServiceClients().AuthServiceClient)
	}

	return s.authService
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(s.ServiceClients().UserServiceClient)
	}

	return s.userService
}

func (s *serviceProvider) MovieService(ctx context.Context) service.MovieService {
	if s.movieService == nil {
		s.movieService = movieService.NewService(s.ServiceClients().MovieServiceClient)
	}

	return s.movieService
}

func (s *serviceProvider) TalentService(ctx context.Context) service.TalentService {
	if s.talentService == nil {
		s.talentService = talentService.NewService(s.ServiceClients().TalentServiceClient)
	}

	return s.talentService
}
