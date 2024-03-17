package talent

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/client/converter"
	"github.com/lukmandev/nameless/gateway/internal/service/model"
)

func (c *cli) GetByIDs(ctx context.Context, ids []string) ([]*model.Talent, error) {
	response, err := c.talentClient.GetByIDs(ctx, converter.ToGetTalentsByIDsFromService(ids))
	if err != nil {
		return nil, err
	}
	return converter.ToTalentFromTalentDescList(response.GetTalents()), nil
}
