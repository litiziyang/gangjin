package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"security/proto/token"
	"time"
)

func GetToken(user *token.TokenUser) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, buildClaims(user)).SignedString(key)

}

func buildClaims(user *token.TokenUser) *Claims {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Hour * time.Duration(24) * time.Duration(30))
	return &Claims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "it-is-gangjin",
		},
	}
}

func ParseUser(token string) (*token.TokenUser, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	claim, ok := tokenClaims.Claims.(*Claims)
	if !ok {
		return nil, errors.New("token解析失败")
	}
	return claim.User, nil
}
