package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword - хэширует пароль
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// VerifyPassword - сверяет хэш пароля и сам пароль
func VerifyPassword(hashedPassword string, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))

	return err == nil
}
