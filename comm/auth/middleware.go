package auth

import (
	"context"
	"fmt"
	"net/http"
	"security/proto/token"
	"strings"
)

var userCtxKey = &contextKey{"USER"}

type contextKey struct {
	mode string
}

func Middleware(service *token.TokenService) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			headerToken := r.Header.Get("Authorization")
			if headerToken == "" {
				next.ServeHTTP(w, r)
				return
			}
			headerToken = strings.Replace(headerToken, "Bearer", "", 1)
			res, err := (*service).GetSecurityUser(r.Context(), &token.TokenRequest{Token: headerToken})
			if err != nil {
				fmt.Println("Token拦截器错误：", err)
				next.ServeHTTP(w, r)
				return
			}
			ctx := context.WithValue(r.Context(), userCtxKey, res.User)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)

		})
	}
}

func User(ctx context.Context) *token.TokenUser {
	user, ok := ctx.Value(userCtxKey).(*token.TokenUser)
	fmt.Println("登录用户：", user)
	if !ok {
		return nil
	}
	return user
}
