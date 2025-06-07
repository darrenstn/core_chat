package serviceimpl

import (
	"core_chat/application/chat/dto"
	"core_chat/application/chat/service"
	"core_chat/infra/serviceimpl/mapper"
	"encoding/json"
	"errors"
	"fmt"
)

type DirectMessageServiceImpl struct{}

func NewDirectMessageServiceImpl() service.DirectMessageService {
	return &DirectMessageServiceImpl{}
}

func (dms *DirectMessageServiceImpl) Execute(input dto.SendMessageInput, wsManager service.WebSocketManager) error {
	infraInput := mapper.ToInfraSendMessageInput(input)

	data, err := json.Marshal(infraInput)
	if err != nil {
		return fmt.Errorf("failed to marshal notification data: %w", err)
	}

	if err := wsManager.Send(infraInput.Receiver, data); err != nil {
		return errors.New("Failed to send message")
	}
	return nil
}
