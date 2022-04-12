package utils

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	jwt.StandardClaims
	Identifier string `json:"identifier"`
}

func GetToken(identifier string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Identifier: identifier,
	})
	return token.SignedString([]byte("qwerty12345"))
}

func ParseToken(tokenStr string, signingKey []byte) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("error with signing method of token")
		}
		return signingKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.Identifier, nil
	}
	return "", errors.New("invalid token")
}
