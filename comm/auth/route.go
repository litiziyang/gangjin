package auth

import (
	"github.com/go-chi/chi"
	"security/proto/token"
)

func GetRoute(tokenSrv *token.TokenService) *chi.Mux {
	router := chi.NewRouter()
	router.Use(Middleware(tokenSrv))
	return router
}
