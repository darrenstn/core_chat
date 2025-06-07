package service

import (
	"core_chat/application/chat/dto"
)

type MessageNotifierService interface {
	ExecuteViaWebSocket(input dto.SendMessageInput, wsManager WebSocketManager) error
	ExecuteViaPushNotifier(input dto.SendMessageInput, pushNotifier PushNotifier) error
}
