package serviceimpl

import (
	appchatdto "core_chat/application/chat/dto"
	"core_chat/application/chat/service"
	"core_chat/infra/serviceimpl/mapper"
	"encoding/json"
	"errors"
	"fmt"
)

type MessageNotifierServiceImpl struct{}

func NewMessageNotifierServiceImpl() service.MessageNotifierService {
	return &MessageNotifierServiceImpl{}
}

func (mns *MessageNotifierServiceImpl) ExecuteViaWebSocket(input appchatdto.SendMessageInput, wsManager service.WebSocketManager) error {
	infraNotificationInput := mapper.ToInfraSendMessageInput(input)

	data, err := json.Marshal(infraNotificationInput)
	if err != nil {
		return fmt.Errorf("Failed to marshal notification data: %w", err)
	}

	if err := wsManager.Send(infraNotificationInput.Receiver, data); err != nil {
		return errors.New("Failed to send message")
	}
	return nil
}

func (mns *MessageNotifierServiceImpl) ExecuteViaPushNotifier(input appchatdto.SendMessageInput, pushNotifier service.PushNotifier) error {
	pushNotificationInput := mapper.ToSendNotificationInput(input)

	if err := pushNotifier.Send(pushNotificationInput); err != nil {
		return fmt.Errorf("Failed to push notification: %w", err)
	}

	return nil
}
