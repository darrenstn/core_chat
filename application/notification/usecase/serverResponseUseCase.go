package usecase

import (
	"core_chat/application/notification/dto"
	"core_chat/application/notification/service"
)

type ServerResponseUseCase struct {
	wsManager   service.WebSocketManager
	notifierSvc service.NotifierService
}

func NewServerResponseUseCase(wsManager service.WebSocketManager, notifierSvc service.NotifierService) *ServerResponseUseCase {
	return &ServerResponseUseCase{
		wsManager:   wsManager,
		notifierSvc: notifierSvc,
	}
}

func (uc *ServerResponseUseCase) Execute(receiver string, status string, responseType string) error {
	resp := dto.ServerResponse{
		Type:   responseType,
		Status: status,
	}

	return uc.notifierSvc.SendResponse(resp, receiver, uc.wsManager)
}
