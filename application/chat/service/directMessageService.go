package service

import (
	"core_chat/application/chat/dto"
)

type DirectMessageService interface {
	Execute(input dto.SendMessageInput, wsManager WebSocketManager) error
}
