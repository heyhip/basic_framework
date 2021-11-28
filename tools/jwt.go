package tools

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type Jwt struct {
}

// 秘钥，加密数据
func (this Jwt) GenerateToken(secret []byte, claims jwt.MapClaims) (token string, err error) {
	// 创建一个新的令牌对象，指定签名方法和声明
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	return
}

// 解析token
func (this Jwt) ParseToken(token string, secret []byte) (jwt.MapClaims, error) {
	t, e := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		return claims, nil
	}
	return nil, e
}
