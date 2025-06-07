package usecase

import (
	"core_chat/application/chat/dto"
	"core_chat/application/chat/repository"
)

type GetChatImageUseCase struct {
	ChatRepo repository.ChatRepository
}

func NewGetChatImageUseCase(repo repository.ChatRepository) *GetChatImageUseCase {
	return &GetChatImageUseCase{
		ChatRepo: repo,
	}
}

func (uc *GetChatImageUseCase) Execute(imageName, imagePath, identifier string) dto.ChatImageResult {
	fullImagePath := imagePath + imageName
	if !uc.ChatRepo.IsImageCanBeRetrieved(fullImagePath, identifier) {
		return dto.ChatImageResult{
			Success: false,
			Message: "Image not found or not accessible",
		}
	}

	return dto.ChatImageResult{
		Success:     true,
		PicturePath: fullImagePath,
	}
}
