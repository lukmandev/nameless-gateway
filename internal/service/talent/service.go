package talent

import (
	"github.com/lukmandev/nameless/gateway/internal/client"
	"github.com/lukmandev/nameless/gateway/internal/service"
)

type serv struct {
	talentClient client.TalentServiceClient
}

func NewService(talentClient client.TalentServiceClient) service.TalentService {
	return &serv{
		talentClient: talentClient,
	}
}
