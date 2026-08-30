package core_encrypt

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", fmt.Errorf("failed to encrypt password")
	}

	return string(hashedPassword), nil
}

func CompareHashAndPassword(hash string, password string) error {
	if err := bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	); err != nil {
		return fmt.Errorf("wrong password")
	}

	return nil
}
