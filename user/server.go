package main

import (
	"comm/auth"
	"comm/directive"
	"comm/tool"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	logrus "github.com/sirupsen/logrus"
	"go-micro.dev/v4"
	"net/http"
	"os"
	"security/proto/token"
	"security/proto/verify"
	"user/graph"
	"user/graph/dataloader"
	"user/graph/generated"
	"user/graph/model"
)

const defaultPort = "4001"

var log = logrus.New()

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	var tokenSrv token.TokenService
	var verifySrv verify.VerifyService

	db, log, _ := tool.CreateServices(&tool.ServerConfig{
		Name: "user",
		TestingFunc: func(service *micro.Service) {
			verifySrv = verify.NewVerifyService("security", (*service).Client())
			tokenSrv = token.NewTokenService("security", (*service).Client())
		},
		Models: func() []any {
			return []any{&model.User{}}
		},
	})
	c := generated.Config{Resolvers: &graph.Resolver{
		Db:              db,
		Log:             log,
		SecurityService: tokenSrv,
		VerifyService:   verifySrv,
	}}
	c.Directives.Auth = directive.Auth()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	router := auth.GetRoute(&tokenSrv)
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, dataloader.Middleware(db, router)))
}
