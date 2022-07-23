package testService

import (
	rediss "comm/redis"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"security/jwt"
	"security/proto/token"
	"testing"
	"time"
)

func TestVerify(t *testing.T) {
	//users := &User{Phone: "12345678", cods: "4555", password: " "}
	strings := GetSecurityUser(context.TODO())
	if strings == 1 {
		t.Error(strings)
	}
	t.Log(strings)
}

type User struct {
	Phone    string
	cods     string
	password string
}

const code string = "4555"

func sendMessage(ctx context.Context, user *User) string {
	Rb := rediss.GetRedisClient()
	_, err := Rb.Get(ctx, user.Phone).Result()
	switch {
	case err == redis.Nil:
	case err != nil:
		fmt.Println("服务失败")
	}
	//todo:后续需要接入sms服务
	err = Rb.Set(ctx, user.Phone, code, time.Hour).Err()
	if err != nil {
		fmt.Println("服务失败")
	}
	result, err := Rb.Get(ctx, user.Phone).Result()
	if err != nil {
		fmt.Println("redis,失败")
	}
	return result
}

func checkVerity(ctx context.Context, user *User) string {
	Rb := rediss.GetRedisClient()
	err := Rb.Set(ctx, user.Phone, code, time.Hour).Err()
	if err != nil {
		fmt.Println("服务失败")
	}
	code, err := Rb.Get(ctx, user.Phone).Result()
	if err != nil {
		fmt.Println(err)
	}
	switch {
	case err == redis.Nil:
	case err != nil:
		fmt.Println("验证失败")
	case code == "":
	}
	if user.cods == code {

	}
	return "成功"

}

func GetSecurityToken(ctx context.Context) string {
	getToken, err := jwt.GetToken(&token.TokenUser{Id: 1})
	if err != nil {
		fmt.Println("生成失败")
		return "失败"
	}

	return getToken
}

func GetSecurityUser(ctx context.Context) uint64 {
	user, err := jwt.ParseUser("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjoyfSwiZXhwIjoxNjU1NDU1NzQ3LCJpc3MiOiJpdC1pcy1nYW5namluIn0._OxV6JEtbfazwIeUbwqYqH907PlnoKf8rUjFbNeyVpE")
	if err != nil {
		fmt.Println(err)
		return 1
	}

	return user.Id
}
