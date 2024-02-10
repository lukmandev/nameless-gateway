package talent

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/service/model"
)

func (s *serv) GetByID(ctx context.Context, id string) (*model.Talent, error) {
	talent, err := s.talentClient.GetByID(ctx, id)
	return talent, err
}
