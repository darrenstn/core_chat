package service

import "core_chat/application/notification/dto"

type NotifierService interface {
	Send(input dto.SendNotificationInput) error
	IsOnline(identifier string) bool
	SetWebSocketManager(ws WebSocketManager)
}
