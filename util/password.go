package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func EncodePassword(rawPassword string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	return string(hash)
}

func ValidatePassword(encodePassword, rawPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(encodePassword), []byte(rawPassword))
	return err
}
