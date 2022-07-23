package jwt

import (
	"comm/tool"
	"github.com/dgrijalva/jwt-go"
	"security/proto/token"
)

type Claims struct {
	User *token.TokenUser
	jwt.StandardClaims
}

var key []byte

func init() {
	key = []byte(tool.GetEnvDefault("JWT_KEY", "IWY@*3JUI#d309HhefzX2WpLtPKtD!hn"))
}
