package utils

import (
	"bytes"
	"sister-backend/exception"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(pass string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	exception.PanicWhenError(err)

	return bytes.NewBuffer(hash).String()
}

func ValidatePassword(pass string, hash string) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	exception.PanicValidationWhenError(err)
}
