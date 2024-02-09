package app

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/service"
	authService "github.com/lukmandev/nameless/gateway/internal/service/auth"
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
