package utils

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

const issuer = "micro_test_mall" //签名
var (
	UserExpireDuration  = time.Hour
	AdminExpireDuration = time.Hour

	UserSecretKey  = []byte("user_token_key")
	AdminSecretKey = []byte("admin_token_key")
)

type UserTokenClaims struct {
	jwt.StandardClaims //jwt标准字段
	//自定义用户信息
	UserName string `json:"user_name"`
	UserId   int    `json:"user_id"`
}

type Admin struct {
}

// 生成token
func GenToken(userId int, UserName string, expired time.Duration,
	secret_key []byte) (string, error) {
	user := UserTokenClaims{
		StandardClaims: jwt.StandardClaims{
			//过期时间（在某一个时间点过期）  相对于 现在时间 + 传递的过期时间
			ExpiresAt: time.Now().Add(expired).Unix(),
			Issuer:    issuer,
		},
		UserName: UserName,
		UserId:   userId,
	}

	//生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	//加密token并返回
	return token.SignedString(secret_key)
}

// 认证token
func AuthToken(tokenString string, secret_key []byte) (*UserTokenClaims, error) {
	//var userClaims UserTokenClaims
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &UserTokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return secret_key, nil
		})

	if err != nil {
		return nil, err
	}
	claims, is_ok := token.Claims.(*UserTokenClaims)
	//验证token
	if !token.Valid {
		return nil, errors.New("token valid err")
	}
	if !is_ok {
		return nil, errors.New("to claims err")
	}

	return claims, nil
}
