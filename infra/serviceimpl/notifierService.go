package serviceimpl

import (
	appdto "core_chat/application/notification/dto"
	"core_chat/application/notification/service"
	infradto "core_chat/infra/serviceimpl/dto"
	"encoding/json"
	"errors"
	"fmt"
)

type NotifierServiceImpl struct{}

func NewNotifierServiceImpl() service.NotifierService {
	return &NotifierServiceImpl{}
}

func (mns *NotifierServiceImpl) SendResponse(input appdto.ServerResponse, receiver string, wsManager service.WebSocketManager) error {
	resp := infradto.ServerResponse{
		Type:   input.Type,
		Status: input.Status,
	}

	data, err := json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("Failed to marshal notification data: %w", err)
	}

	if err := wsManager.Send(receiver, data); err != nil {
		return errors.New("Failed to send message")
	}
	return nil
}
