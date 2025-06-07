package repository

import (
	"core_chat/application/chat/dto"
	"core_chat/application/chat/entity"
)

type ChatRepository interface {
	SaveImageMetadata(metadata dto.ChatImageMetadata) error
	ExistsByIdentifier(identifier string) bool
	IsImageCanBeRetrieved(imagePath, identifier string) bool
	SaveMessage(input dto.SendMessageInput) (string, error)
	FindChatMessage(msgID, identifier string) (*entity.Message, error)
	MarkMessageAsRead(msgID string, receiver string) error
}
