package services

import (
	"comm/crons"
	"comm/logger"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"security/hackerNews"
	"security/jwt"
	"security/proto/hot"
	"security/proto/token"
	"security/proto/verify"
	"time"
)

type Service struct {
	Log *logrus.Logger
	Rb  *redis.Client
	DB  *gorm.DB
}

func (s Service) GetSecurityUser(ctx context.Context, request *token.TokenRequest, response *token.TokenResponse) error {
	user, err := jwt.ParseUser(request.Token)
	if err != nil {
		logger.Error(s.Log, "凭证验证失败", err, request)
		response.Error = "凭证验证失败"
		return err
	}
	response.User = user
	return nil
}

func (s Service) GetSecurityToken(ctx context.Context, request *token.TokenRequest, response *token.TokenResponse) error {
	getToken, err := jwt.GetToken(request.User)
	if err != nil {
		logger.Error(s.Log, "生成错误", err, request)
		response.Error = "凭证生成失败"
		return err
	}
	response.Token = getToken
	return nil
}

const code string = "4555"

func (s Service) SendVerify(ctx context.Context, request *verify.VerifyRequest, response *verify.VerifyResponse) error {
	_, err := s.Rb.Get(ctx, request.Phone).Result()
	switch {
	case err == redis.Nil:
	case err != nil:
		response.Error = "短信发送失败"
		logger.Error(s.Log, "短信缓存错误", err, nil)
		return err
	}
	//todo:后续需要接入sms服务
	err = s.Rb.Set(ctx, request.Phone, code, time.Hour).Err()
	if err != nil {
		response.Error = "短信发送失败"
		logger.Error(s.Log, "短信发送失败", err, nil)
		return err
	}
	return nil

}

func (s Service) CheckVerify(ctx context.Context, request *verify.VerifyRequest, response *verify.VerifyResponse) error {
	code, err := s.Rb.Get(ctx, request.Phone).Result()
	if err != nil {
		response.Error = "验证码不正确"
	}
	switch {
	case err == redis.Nil:
	case err != nil:
		logger.Error(s.Log, "验证失败", err, nil)
	case code == " ":
	}
	if request.Code == code {
		response.Error = ""
		response.Check = true
	}
	return nil

}

func (s Service) GetSecurityHot(ctx context.Context, request *hot.HotRequest, response *hot.HotResponse) error {
	//var hottask = make(chan float64)
	crons.InserTimerTask(fmt.Sprintf("%s_id %s", request.ModelName, request.Id), 2*time.Minute, func(id string, err error) error {
		Hot := hackerNews.MakeHot(request)
		switch request.ModelName {
		case "post":
			s.DB.Table("post").Where("id", request.Id).Update("heat", Hot)
		case "discussion":
			s.DB.Table("discussion").Where("id", request.Id).Update("hot", Hot)
		}
		return nil
	})
	return nil
}
