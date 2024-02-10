package talent

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/api/converter"
	"github.com/lukmandev/nameless/gateway/internal/api/model"
)

func (q TalentQuery) GetTalentsByIDs(ctx context.Context, ids []string) ([]*model.Talent, error) {
	talents, err := q.TalentService.GetByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}
	return converter.ToTalentFromServiceList(talents), nil
}
