package app

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/service"
	authService "github.com/lukmandev/nameless/gateway/internal/service/auth"
)

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = authService.NewService(s.ServiceClients())
	}

	return s.authService
}
