package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pimp13/server-chi/pkg/config"
)

var jwtKey = []byte(config.Envs.JWTKey)

func MakeToken(userID uint) (string, error) {
	exp := time.Second * time.Duration(config.Envs.JWTExpirationInSecond)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"ID":  userID,
		"exp": time.Now().Add(exp).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %s", err.Error())
	}

	return token, nil
}
