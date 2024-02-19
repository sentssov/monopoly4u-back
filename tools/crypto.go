package tools

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func CompareHash(hashedPassword []byte, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	return err == nil
}
