package util

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"project/pkg/setting"
	"time"
)

var jwtSecret = []byte(setting.JWTSecret)
type Claims struct {
	Username 		string `json:"username"`
	Password		string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	expiredTime := time.Now().Add(24 * time.Hour).Unix()

	claim := &Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expiredTime,
			Issuer: "blog",
		},
	}

	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodES256, claim)
	token, err := tokenClaim.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaim, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaim == nil {
		return nil, errors.New("parse jwt happen some err")
	}

	if claims, ok := tokenClaim.Claims.(*Claims); ok && tokenClaim.Valid {
		return claims, nil
	}
	return nil, err
}