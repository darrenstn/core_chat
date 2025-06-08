package usecase

import (
	"core_chat/application/authentication/service"
)

type LogoutUseCase struct {
	WsManager service.WebSocketManager
}

func NewLogoutUseCase(wsManager service.WebSocketManager) *LogoutUseCase {
	return &LogoutUseCase{WsManager: wsManager}
}

func (uc *LogoutUseCase) Execute(identifier string) {
	uc.WsManager.CloseConnection(identifier)
}
