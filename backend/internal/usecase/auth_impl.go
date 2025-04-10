package usecase

import (
	"golang.org/x/crypto/bcrypt"
)

type authRepositoryImpl struct {
}

func NewAuth() Auth {
	return &authRepositoryImpl{}
}

func (r *authRepositoryImpl) HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

func (r *authRepositoryImpl) ComparePassword(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil

}
