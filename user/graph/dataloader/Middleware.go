package dataloader

import (
	"context"
	"gorm.io/gorm"
	"net/http"
)

func Middleware(db *gorm.DB, next http.Handler) http.Handler {
	// return a middleware that injects the loader to the request context
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCtx := context.WithValue(r.Context(), loadersKey, &Loaders{
			UserById: *GetUserLoader(db),
		})
		r = r.WithContext(nextCtx)
		next.ServeHTTP(w, r)
	})
}
