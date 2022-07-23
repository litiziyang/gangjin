package tool

import (
	"comm/database"
	"comm/logger"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/sirupsen/logrus"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"gorm.io/gorm"
)

var log = logrus.New()

type (
	ServerConfig struct {
		Name           string
		LogLevel       *string
		TestingFunc    func(*micro.Service)
		DebugFunc      func(*micro.Service)
		ProdFunc       func(*micro.Service)
		Models         func() []any
		ServiceBuilder func(name string, reg registry.Registry) *micro.Service
	}
)

var db *gorm.DB

func CreateServices(c *ServerConfig) (*gorm.DB, *logrus.Logger, *micro.Service) {
	register := GetEnvDefault("REGISTER", "127.0.0.1:8500")
	es := GetEnvDefault("ELASTICSEARCH", "127.0.0.1:9200")

	loglevel := "debug"
	if c.LogLevel != nil {
		loglevel = *c.LogLevel
	}
	err := logger.SetupLog(es, loglevel, c.Name)
	if err != nil {
		log.Fatal("log初始化失败", err)
	}

	reg := consul.NewRegistry(
		func(options *registry.Options) {
			options.Addrs = []string{register}
		})

	var mco micro.Service
	if c.ServiceBuilder != nil {
		mco = *c.ServiceBuilder(c.Name, reg)
	} else {
		mco = micro.NewService(
			micro.Name(c.Name),
			micro.Registry(reg),
		)
	}
	mco.Init()
	if c.Models != nil {
		db, err = database.InitDb()
		if err != nil {
			log.Panic(err)
		}
	}
	ev := GetEnvDefault("ENVIRONMENT", "testing")
	switch ev {
	case "testing":
		if c.TestingFunc != nil {
			c.TestingFunc(&mco)
			if c.Models != nil {
				err = database.Migrate(c.Models()...)
				if err != nil {
					log.Panic(err)
				}
			}
		}
	case "debug":
		if c.DebugFunc != nil {
			c.DebugFunc(&mco)
			if c.Models != nil {
				err = database.Migrate(c.Models()...)
				if err != nil {
					log.Panic(err)
				}
			}
		}
		fallthrough
	case "prod":
		if c.ProdFunc != nil {
			c.ProdFunc(&mco)
		}
	}
	return db, log, &mco
}
