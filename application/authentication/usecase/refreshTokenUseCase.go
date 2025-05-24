package usecase

import (
	"core_chat/application/authentication/dto"
	"core_chat/application/authentication/repository"
	"core_chat/application/authentication/service"
)

type RefreshTokenUseCase struct {
	PersonRepo   repository.PersonRepository
	TokenService service.TokenService
}

func NewRefreshTokenUseCase(repo repository.PersonRepository, ts service.TokenService) *RefreshTokenUseCase {
	return &RefreshTokenUseCase{
		PersonRepo:   repo,
		TokenService: ts,
	}
}

func (uc *RefreshTokenUseCase) Execute(identifier string) dto.AuthResult {
	person, err := uc.PersonRepo.GetPersonByIdentifier(identifier)
	if err != nil {
		return dto.AuthResult{Success: false, Message: "Invalid Credential"}
	}

	return generateAuthResult(person, uc.TokenService, "Refresh Token Success")
}
