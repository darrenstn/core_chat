package usecase

import (
	chatdto "core_chat/application/chat/dto"
	"core_chat/application/chat/repository"
	chatservice "core_chat/application/chat/service"
	"errors"
)

type SendMessageUseCase struct {
	repo        repository.ChatRepository
	wsManager   chatservice.WebSocketManager
	notifier    chatservice.PushNotifier
	dmService   chatservice.DirectMessageService
	msgNotifier chatservice.MessageNotifierService
}

func NewSendMessageUseCase(
	repo repository.ChatRepository,
	wsManager chatservice.WebSocketManager,
	notifier chatservice.PushNotifier,
	dmService chatservice.DirectMessageService,
	msgNotifier chatservice.MessageNotifierService,
) *SendMessageUseCase {
	return &SendMessageUseCase{
		repo:        repo,
		wsManager:   wsManager,
		notifier:    notifier,
		dmService:   dmService,
		msgNotifier: msgNotifier,
	}
}

func (uc *SendMessageUseCase) Execute(input chatdto.SendMessageInput) error {
	// Step 1: Check if receiver exists
	if !uc.repo.ExistsByIdentifier(input.Receiver) {
		return errors.New("Receiver doesn't exist")
	}

	// Step 2: Save message to repository
	messageID, err := uc.repo.SaveMessage(input)
	if err != nil {
		return err
	}

	// Step 3: Check if receiver is online
	if uc.wsManager.IsOnline(input.Receiver) {
		chatID := uc.wsManager.GenerateChatID(input.Receiver, input.Sender)
		if uc.wsManager.IsPersonInRoom(chatID, input.Receiver) {
			// Same room: send real-time message
			err = uc.dmService.Execute(input, uc.wsManager)
			if err != nil {
				return err
			}
			uc.repo.MarkMessageAsRead(messageID, input.Receiver)
			return err
		}
		input.Type = "chat_notification"
		input.Title = "New Message"
		input.Body = ""
		input.Payload = messageID
		// Online but not in the same room: send WebSocket notification
		return uc.msgNotifier.ExecuteViaWebSocket(input, uc.wsManager)
	}

	// Step 4: Receiver is offline, fallback to push notification
	return uc.msgNotifier.ExecuteViaPushNotifier(input, uc.notifier)
}
