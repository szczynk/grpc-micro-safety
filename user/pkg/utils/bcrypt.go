package utils

import (
	"user/pkg/grpc_errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	if len(password) == 0 {
		return "", grpc_errors.ErrNoPassword
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func CheckPasswordHash(hash string, password string) error {
	if len(password) == 0 {
		return grpc_errors.ErrNoPassword
	}

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err
}
