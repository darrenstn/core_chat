package usecase

import (
	"core_chat/application/authentication/dto"
	"core_chat/application/authentication/repository"
	"core_chat/application/authentication/service"
)

type RefreshTokenUseCase struct {
	PersonRepo   repository.PersonRepository
	TokenService service.TokenService
	WsManager    service.WebSocketManager
}

func NewRefreshTokenUseCase(repo repository.PersonRepository, ts service.TokenService, wsManager service.WebSocketManager) *RefreshTokenUseCase {
	return &RefreshTokenUseCase{
		PersonRepo:   repo,
		TokenService: ts,
		WsManager:    wsManager,
	}
}

func (uc *RefreshTokenUseCase) Execute(identifier string) dto.AuthResult {
	person, err := uc.PersonRepo.GetPersonByIdentifier(identifier)
	if err != nil {
		return dto.AuthResult{Success: false, Message: "Invalid Credential"}
	}
	res := generateAuthResult(person, uc.TokenService, "Refresh Token Success")

	uc.WsManager.UpdateToken(identifier, res.Token)

	return res
}
