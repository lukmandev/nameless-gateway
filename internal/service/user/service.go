package user

import (
	"github.com/lukmandev/nameless/gateway/internal/service"
)

type serv struct {
}

func NewService() service.UserService {
	return serv{}
}
