package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/lukmandev/nameless/gateway/internal/api/converter"
	"github.com/lukmandev/nameless/gateway/internal/api/model"
	"github.com/lukmandev/nameless/gateway/internal/service"
)

var (
	cookieName = "rt"

	CookieAccessKey = "cookie-access"
)

type CookieAccess struct {
	Writer     http.ResponseWriter
	IsLoggedIn bool
	UserId     uint64

	AccessToken string
}

func (access *CookieAccess) SetRefreshToken(token string) {
	http.SetCookie(access.Writer, &http.Cookie{
		Name:     cookieName,
		Value:    token,
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24 * 15),
	})
}

func AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authorizationHeader := r.Header["Authorization"]
			accessToken := ""
			if len(authorizationHeader) > 0 {
				tokenParts := strings.Split(authorizationHeader[0], " ")
				if len(tokenParts) == 2 && tokenParts[0] == "Bearer" {
					accessToken = tokenParts[1]
				}
			}
			cookieAccess := &CookieAccess{
				Writer:      w,
				AccessToken: accessToken,
			}
			ctx := context.WithValue(r.Context(), CookieAccessKey, cookieAccess)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func GetCookieAccessFromCtx(ctx context.Context) *CookieAccess {
	access := ctx.Value(CookieAccessKey).(*CookieAccess)
	return access
}

func GetUser(ctx context.Context, authService service.AuthService) (*model.Profile, error) {
	access, ok := ctx.Value(CookieAccessKey).(*CookieAccess)
	if !ok {
		return nil, errors.New("not found")
	}

	profile, err := authService.GetMe(ctx, converter.ToGetMeInputFromAuthHandler(access.AccessToken))
	if err != nil {
		return nil, err
	}

	return converter.ToProfileFromService(profile), nil
}
