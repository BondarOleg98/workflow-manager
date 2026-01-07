package util

import (
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

func HashPassword(password string) (string, error) {
	slog.Info("hashing the password")
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), err
}

func VerifyPassword(hashedPassword, providedPassword string) error {
	slog.Debug("verify the password")
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(providedPassword))
}
