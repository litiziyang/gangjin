package graph

import (
	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"security/proto/hot"
	"security/proto/token"
	"security/proto/verify"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	Db              *gorm.DB
	Log             *logrus.Logger
	SecurityService token.TokenService
	VerifyService   verify.VerifyService
	Es              *elastic.Client
	HotServise      hot.HotService
}
