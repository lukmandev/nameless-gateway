package talent

import "github.com/lukmandev/nameless/gateway/internal/service"

type TalentMutation struct {
	AuthService service.AuthService
}

type TalentQuery struct {
	TalentService service.TalentService
}
