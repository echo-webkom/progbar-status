package security

import "errors"

func ValidateToken(token string) error {
	if token == "" {
		return errors.New("authorization header is required")
	}

	return nil
}
