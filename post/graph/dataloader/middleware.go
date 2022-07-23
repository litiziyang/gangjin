package dataloader

import (
	"context"
	"gorm.io/gorm"
	"net/http"
)

func Middleware(db *gorm.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
			PostById: *GetPostLoader(db),
		})
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
