package usecase

import (
	"core_chat/application/chat/repository"
	chatservice "core_chat/application/chat/service"
	"errors"
)

type LeaveChatRoomUseCase struct {
	repo      repository.ChatRepository
	wsManager chatservice.WebSocketManager
}

func NewLeaveChatRoomUseCase(
	repo repository.ChatRepository,
	wsManager chatservice.WebSocketManager,
) *LeaveChatRoomUseCase {
	return &LeaveChatRoomUseCase{
		repo:      repo,
		wsManager: wsManager,
	}
}

func (uc *LeaveChatRoomUseCase) Execute(identifier, person string) error {
	if !uc.repo.ExistsByIdentifier(person) {
		return errors.New("Person doesn't exist")
	}
	chatID := uc.wsManager.GenerateChatID(identifier, person)
	uc.wsManager.LeaveRoom(chatID, identifier)

	return nil
}
