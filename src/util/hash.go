package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	str, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	return string(str), nil
}
