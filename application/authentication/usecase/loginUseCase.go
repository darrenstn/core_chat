package usecase

import (
	"core_chat/application/authentication/dto"
	"core_chat/application/authentication/mapper"
	"core_chat/application/authentication/repository"
	"core_chat/application/authentication/service"
	"time"
)

type LoginUseCase struct {
	PersonRepo   repository.PersonRepository
	TokenService service.TokenService
	HashService  service.HashService
}

func NewLoginUseCase(repo repository.PersonRepository, ts service.TokenService, hs service.HashService) *LoginUseCase {
	return &LoginUseCase{PersonRepo: repo, TokenService: ts, HashService: hs}
}

func (uc *LoginUseCase) Execute(identifier, password string) dto.LoginResult {
	person, err := uc.PersonRepo.GetPersonByIdentifier(identifier)
	if err != nil || !uc.HashService.CompareHash(person.Password, password) {
		return dto.LoginResult{Success: false, Message: "Invalid Credentials"}
	}

	expTime := time.Now().Add(5 * time.Minute)

	claims := mapper.ToClaimsDTO(person, expTime.Unix())

	token, err := uc.TokenService.GenerateToken(claims)
	if err != nil {
		return dto.LoginResult{Success: false, Message: "Failed to Generate Token"}
	}

	return dto.LoginResult{
		Success:    true,
		Token:      token,
		Message:    "Login Successful",
		Expiration: expTime,
	}
}
