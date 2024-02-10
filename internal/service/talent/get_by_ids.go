package talent

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/service/model"
)

func (s *serv) GetByIDs(ctx context.Context, ids []string) ([]*model.Talent, error) {
	talents, err := s.talentClient.GetByIDs(ctx, ids)
	return talents, err
}
