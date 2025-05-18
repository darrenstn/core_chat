package usecase

import (
	"core_chat/application/authentication/repository"
	"core_chat/application/authentication/service"
	"net/http"
)

type LoginUseCase struct {
	PersonRepo   repository.PersonRepository
	TokenService service.TokenService
	HashService  service.HashService
	HTTPService  service.HTTPService
}

func NewLoginUseCase(repo repository.PersonRepository, ts service.TokenService, hs service.HashService, hsrv service.HTTPService) *LoginUseCase {
	return &LoginUseCase{PersonRepo: repo, TokenService: ts, HashService: hs, HTTPService: hsrv}
}

func (uc *LoginUseCase) Execute(w http.ResponseWriter, identifier, password string) bool {
	person, err := uc.PersonRepo.GetPersonByIdentifier(identifier)
	if err != nil || !uc.HashService.CompareHash(person.Password, password) {
		return false
	}
	token, err := uc.TokenService.GenerateToken(person.Identifier, person.Role, person.EmailValidated)
	if err != nil {
		return false
	}

	uc.HTTPService.SetAuthCookie(w, token)
	return true
}
