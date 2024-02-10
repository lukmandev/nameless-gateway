package talent

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/api/converter"
	"github.com/lukmandev/nameless/gateway/internal/api/model"
)

func (q TalentQuery) GetTalentByID(ctx context.Context, ids string) (*model.Talent, error) {
	talent, err := q.TalentService.GetByID(ctx, ids)
	if err != nil {
		return nil, err
	}
	return converter.ToTalentFromService(talent), nil
}
