package mapper

import (
	"core_chat/application/chat/entity"
	"core_chat/web/rest/dto"
)

func ToMessageResult(message *entity.Message) *dto.MessageResult {
	return &dto.MessageResult{
		ID:        message.ID,
		Receiver:  message.Receiver,
		Sender:    message.Sender,
		Type:      message.Type,
		Title:     message.Title,
		Body:      message.Body,
		Payload:   message.Payload,
		CreatedAt: message.CreatedAt,
		ReadAt:    message.ReadAt,
	}
}
