package util

import (
	// "github.com/pimp13/server-chi/pkg/config"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// password += config.Envs.Salt
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckHashPassword(password, hash string) bool {
	// password += config.Envs.Salt
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
