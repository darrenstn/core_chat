package service

import "core_chat/application/notification/dto"

type NotifierService interface {
	SendResponse(input dto.ServerResponse, receiver string, wsManager WebSocketManager) error
}
