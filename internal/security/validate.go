package security

import (
	"errors"
	"os"
)

func ValidateToken(token string) error {
	apiKey := os.Getenv("API_KEY")

	if token != apiKey {
		return errors.New("authorization header is required")
	}

	return nil
}
