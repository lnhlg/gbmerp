package biz

import (
	"gbmerp/app/interface/internal/conf"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/golang-jwt/jwt"
	jwtv4 "github.com/golang-jwt/jwt/v4"
)

type AuthUseCase struct {
	key string
}

type Token struct {
	Ciphertext string
	Expire     int64
}

func NewAuthUseCase(conf *conf.Auth) *AuthUseCase {
	return &AuthUseCase{
		key: conf.ApiKey,
	}
}

type MyClaims struct {
	UserName string `json:"username"`
	jwtv4.StandardClaims
}

func (receiver AuthUseCase) Auth(
	username string,
	expire int64,
) (*Token, error) {
	myClaims := MyClaims{
		username,
		jwtv4.StandardClaims{
			ExpiresAt: expire, //设置JWT过期时间
		},
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)

	token, err := claims.SignedString([]byte(receiver.key))
	if err != nil {
		return nil, err
	}

	return &Token{
		Ciphertext: token,
		Expire:     expire,
	}, nil
}

func (receiver AuthUseCase) CheckJWT(jwtToken string) (map[string]interface{}, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(receiver.key), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		result := make(map[string]interface{}, 2)
		result["username"] = claims["username"]
		result["differentiate"] = claims["differentiate"]
		return result, nil
	} else {
		return nil, errors.Forbidden("CheckJWT", "token type error")
	}
}
