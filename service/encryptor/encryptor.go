package encryptor

import (
	"golang.org/x/crypto/bcrypt"
)

type encryptor struct{}

func New() *encryptor {
	return &encryptor{}
}

func (e *encryptor) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (e *encryptor) ComparePasswordAndHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
