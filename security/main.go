package main

import (
	"comm/redis"
	"comm/tool"
	"security/proto/token"
	"security/proto/verify"
	"security/services"
)

func main() {

	_, logger, srv := tool.CreateServices(&tool.ServerConfig{
		Name: "security",
	})
	if err := token.RegisterTokenServiceHandler((*srv).Server(), &services.Service{
		Log: logger,
		Rb:  redis.GetRedisClient(),
	}); err != nil {
		logger.Fatal("服务绑定失败", err)
	}
	if err := verify.RegisterVerifyServiceHandler((*srv).Server(), &services.Service{
		Log: logger,
		Rb:  redis.GetRedisClient(),
	}); err != nil {
		logger.Fatal("服务绑定失败", err)
	}
	if err := (*srv).Run(); err != nil {
		logger.Fatal("服务运行失败", err)
	}
}
