package serviceimpl

import (
	"core_chat/application/authentication/service"

	"golang.org/x/crypto/bcrypt"
)

type bcryptHashService struct{}

func NewBcryptHashService() service.HashService {
	return &bcryptHashService{}
}

func (s *bcryptHashService) CompareHash(hashed, plain string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain)) == nil
}

func (s *bcryptHashService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
