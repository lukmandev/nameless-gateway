package talent

import (
	"context"

	"github.com/lukmandev/nameless/gateway/internal/client/converter"
	"github.com/lukmandev/nameless/gateway/internal/service/model"
)

func (c *cli) GetByID(ctx context.Context, id string) (*model.Talent, error) {
	response, err := c.talentClient.GetByID(ctx, converter.ToGetTalentByIDFromService(id))
	if err != nil {
		return nil, err
	}
	return converter.ToTalentFromTalentDesc(response.GetTalent()), nil
}
