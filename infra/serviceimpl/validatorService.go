package serviceimpl

import (
	chatservice "core_chat/application/chat/service"
	personservice "core_chat/application/person/service"
	"net/mail"
	"regexp"
)

var (
	_ personservice.ValidatorService = (*ValidatorService)(nil)
	_ chatservice.ValidatorService   = (*ValidatorService)(nil)
)

type ValidatorService struct{}

func NewPersonValidatorService() personservice.ValidatorService {
	return &ValidatorService{}
}

func NewChatValidatorService() chatservice.ValidatorService {
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
