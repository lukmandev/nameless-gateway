package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/lukmandev/nameless/gateway/internal/api/converter"
	"github.com/lukmandev/nameless/gateway/internal/api/loaders"
	"github.com/lukmandev/nameless/gateway/internal/api/model"
	"github.com/lukmandev/nameless/gateway/internal/service"
)

var (
	TalentLoaderKey = "talentLoaderKey"
)

func TalentLoaderMiddleware(talentService service.TalentService) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cfg := loaders.TalentLoaderConfig{
				MaxBatch: 100,
				Wait:     1 * time.Millisecond,
				Fetch: func(ids []string) ([]*model.Talent, []error) {
					fmt.Println("ids", ids)
					talents, err := talentService.GetByIDs(context.Background(), ids)
					if err != nil {
						errors := make([]error, 0)
						for i := 0; i < len(ids); i++ {
							errors = append(errors, err)
						}

						return nil, errors
					}
					return converter.ToTalentFromServiceList(talents), nil
				},
			}
			userLoader := loaders.NewTalentLoader(cfg)
			ctx := context.WithValue(r.Context(), TalentLoaderKey, userLoader)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
