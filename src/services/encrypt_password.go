package services

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string) (string, error) {
	salt := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(bytes), err
}
