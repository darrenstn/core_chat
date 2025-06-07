package usecase

import (
	"core_chat/application/chat/repository"
	chatservice "core_chat/application/chat/service"
	"errors"
)

type JoinChatRoomUseCase struct {
	repo      repository.ChatRepository
	wsManager chatservice.WebSocketManager
}

func NewJoinChatRoomUseCase(
	repo repository.ChatRepository,
	wsManager chatservice.WebSocketManager,
) *JoinChatRoomUseCase {
	return &JoinChatRoomUseCase{
		repo:      repo,
		wsManager: wsManager,
	}
}

func (uc *JoinChatRoomUseCase) Execute(identifier, person string) error {
	if !uc.repo.ExistsByIdentifier(person) {
		return errors.New("Person doesn't exist")
	}
	chatID := uc.wsManager.GenerateChatID(identifier, person)
	uc.wsManager.JoinRoom(chatID, identifier)

	return nil
}
