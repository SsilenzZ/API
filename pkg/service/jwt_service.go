package service

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type JwtI interface {
	GenerateAccessToken(id int) (t string)
}

type Jwt struct {
}

type JWTCustomClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

func (J Jwt) GenerateAccessToken(id int) (t string) {
	claims := &JWTCustomClaims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return
	}
	return t
}
