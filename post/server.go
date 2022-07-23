package main

import (
	"comm/auth"
	"comm/crons"
	"comm/directive"
	es2 "comm/es"
	"comm/tool"
	"github.com/sirupsen/logrus"
	"go-micro.dev/v4"
	"net/http"
	"os"
	"post/graph"
	"post/graph/dataloader"
	"post/graph/generated"
	"post/graph/model"
	"security/proto/hot"
	"security/proto/token"
	"security/proto/verify"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "4002"

var log = logrus.New()

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	//定时任务
	err := crons.InitProcessTimer()
	if err != nil {
		log.Error(err)
	}

	//获取es客户端
	clientEs, err := es2.GetESClient()
	if err != nil {
		log.Error(err)
	}

	var tokenSrv token.TokenService
	var verifySrv verify.VerifyService
	var hotSrv hot.HotService

	db, log, _ := tool.CreateServices(&tool.ServerConfig{
		Name: "post",
		TestingFunc: func(service *micro.Service) {
			verifySrv = verify.NewVerifyService("security", (*service).Client())
			tokenSrv = token.NewTokenService("security", (*service).Client())
			hotSrv = hot.NewHotService("security", (*service).Client())
		},
		Models: func() []any {
			return []any{&model.Post{}}
		},
	})

	config := generated.Config{Resolvers: &graph.Resolver{
		Db:              db,
		Log:             log,
		SecurityService: tokenSrv,
		VerifyService:   verifySrv,
		Es:              clientEs,
		HotServise:      hotSrv,
	},
	}
	config.Directives.Auth = directive.Auth()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	router := auth.GetRoute(&tokenSrv)
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, dataloader.Middleware(db, router)))
}
