package middleware

import (
	"context"
	"net/http"
	"time"
)

var (
	cookieName = "rt"

	CookieAccessKey = "cookie-access"
)

type CookieAccess struct {
	Writer     http.ResponseWriter
	IsLoggedIn bool
	UserId     uint64
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
			cookieAccess := &CookieAccess{
				Writer: w,
			}
			ctx := context.WithValue(r.Context(), CookieAccessKey, cookieAccess)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func CookieAccessFromCtx(ctx context.Context) *CookieAccess {
	access := ctx.Value(CookieAccessKey).(*CookieAccess)
	return access
}
