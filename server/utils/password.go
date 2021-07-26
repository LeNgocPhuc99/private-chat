package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func CreatePassword(passwordString string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(passwordString), 8)
	if err != nil {
		return "", errors.New("error while hash password")
	}
	return string(hashPass), nil
}

func CommparePassword(password, hashPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return errors.New("password don't match")
	}

	return nil
}
