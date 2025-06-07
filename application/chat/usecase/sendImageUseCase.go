package usecase

import (
	chatdto "core_chat/application/chat/dto"
	"core_chat/application/chat/repository"
	chatservice "core_chat/application/chat/service"
	"errors"
)

type SendImageUseCase struct {
	repo        repository.ChatRepository
	wsManager   chatservice.WebSocketManager
	notifier    chatservice.PushNotifier
	dmService   chatservice.DirectMessageService
	msgNotifier chatservice.MessageNotifierService
}

func NewSendImageUseCase(
	repo repository.ChatRepository,
	wsManager chatservice.WebSocketManager,
	notifier chatservice.PushNotifier,
	dmService chatservice.DirectMessageService,
	msgNotifier chatservice.MessageNotifierService,
) *SendImageUseCase {
	return &SendImageUseCase{
		repo:        repo,
		wsManager:   wsManager,
		notifier:    notifier,
		dmService:   dmService,
		msgNotifier: msgNotifier,
	}
}

func (uc *SendImageUseCase) Execute(input chatdto.SendMessageInput, imagePath, imageName string) error {
	// Step 1: Check if receiver exists
	if !uc.repo.ExistsByIdentifier(input.Receiver) {
		return errors.New("Receiver doesn't exist")
	}

	fullImagePath := imagePath + imageName

	if !uc.repo.IsImageCanBeRetrieved(fullImagePath, input.Sender) {
		return errors.New("Image can't be retrieved")
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
			uc.dmService.Execute(input, uc.wsManager)
		}
		input.Type = "chat_notification"
		input.Title = "New Message"
		input.Body = ""
		input.Payload = messageID
		// Online but not in the same room: send WebSocket notification
		uc.msgNotifier.ExecuteViaWebSocket(input, uc.wsManager)
	}

	// Step 4: Receiver is offline, fallback to push notification
	return uc.msgNotifier.ExecuteViaPushNotifier(input, uc.notifier)
}
