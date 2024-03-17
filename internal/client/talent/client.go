package talent

import (
	talentDesc "github.com/lukmandev/nameless-movie/pkg/talent_v1"
	"github.com/lukmandev/nameless/gateway/internal/client"
)

type cli struct {
	talentClient talentDesc.TalentV1Client
}

func New(
	talentClient talentDesc.TalentV1Client,
) client.TalentServiceClient {
	return &cli{talentClient: talentClient}
}
