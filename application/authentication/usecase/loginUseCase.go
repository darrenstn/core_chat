package usecase

import (
	"core_chat/application/authentication/dto"
	"core_chat/application/authentication/repository"
	"core_chat/application/authentication/service"
)

type LoginUseCase struct {
	PersonRepo   repository.PersonRepository
	TokenService service.TokenService
	HashService  service.HashService
}

func NewLoginUseCase(repo repository.PersonRepository, ts service.TokenService, hs service.HashService) *LoginUseCase {
	return &LoginUseCase{PersonRepo: repo, TokenService: ts, HashService: hs}
}

func (uc *LoginUseCase) Execute(identifier, password string) dto.AuthResult {
	person, err := uc.PersonRepo.GetPersonByIdentifier(identifier)
	if err != nil || !uc.HashService.CompareHash(person.Password, password) {
		return dto.AuthResult{Success: false, Message: "Invalid Credentials"}
	}

	return generateAuthResult(person, uc.TokenService, "Login Success")
}
