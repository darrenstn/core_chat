package serviceimpl

import (
	"core_chat/application/person/service"
	"net/mail"
	"regexp"
)

type ValidatorService struct{}

func NewValidatorService() service.ValidatorService {
	return &ValidatorService{}
}

var (
	identifierRegex = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	// Password: 4–12 characters, only letters, numbers, @
	passwordRegex = regexp.MustCompile(`^[a-zA-Z0-9@]{4,12}$`)
)

func (s *ValidatorService) IsIdentifierValid(identifier string) bool {
	return identifierRegex.MatchString(identifier)
}

func (s *ValidatorService) IsEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (s *ValidatorService) IsPasswordValid(password string) bool {
	return passwordRegex.MatchString(password)
}
