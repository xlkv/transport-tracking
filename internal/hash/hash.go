package hash

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 12)
}

func ComparePassword(hasedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hasedPassword), []byte(password))
}
