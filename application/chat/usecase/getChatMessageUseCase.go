package usecase

import (
	"core_chat/application/chat/entity"
	"core_chat/application/chat/repository"
)

type GetChatMessageUseCase struct {
	ChatRepo repository.ChatRepository
}

func NewGetChatMessageUseCase(repo repository.ChatRepository) *GetChatMessageUseCase {
	return &GetChatMessageUseCase{
		ChatRepo: repo,
	}
}

func (uc *GetChatMessageUseCase) Execute(messageID, identifier string) (*entity.Message, error) {
	result, err := uc.ChatRepo.FindChatMessage(messageID, identifier)
	if err != nil {
		return nil, err
	}

	if result.ReadAt == "" {
		if result.Receiver == identifier {
			err = uc.ChatRepo.MarkMessageAsRead(result.ID, result.Receiver)
			if err != nil {
				return nil, err
			}
			result, err = uc.ChatRepo.FindChatMessage(messageID, identifier)
			if err != nil {
				return nil, err
			}
		}
	}

	return result, nil
}
