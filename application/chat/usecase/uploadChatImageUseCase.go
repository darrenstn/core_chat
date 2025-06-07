package usecase

import (
	"core_chat/application/chat/dto"
	"core_chat/application/chat/repository"
	"core_chat/application/chat/service"
	"errors"
)

type UploadChatImageUseCase struct {
	AntivirusSvc   service.AntivirusService
	ValidatorSvc   service.ValidatorService
	ChatRepository repository.ChatRepository
}

func NewUploadChatImageUseCase(antivirusSvc service.AntivirusService, chatRepo repository.ChatRepository, validatorSvc service.ValidatorService) *UploadChatImageUseCase {
	return &UploadChatImageUseCase{AntivirusSvc: antivirusSvc, ChatRepository: chatRepo, ValidatorSvc: validatorSvc}
}

func (uc *UploadChatImageUseCase) Execute(filePath, sender, receiver string) error {
	if err := uc.AntivirusSvc.ScanImage(filePath); err != nil {
		return err
	}

	if !uc.ValidatorSvc.IsIdentifierValid(receiver) {
		return errors.New("Invalid receiver identifier")
	}

	if !uc.ChatRepository.ExistsByIdentifier(receiver) {
		return errors.New("Receiver does not exist")
	}

	metadata := dto.ChatImageMetadata{
		Sender:    sender,
		Receiver:  receiver,
		ImagePath: filePath,
	}

	return uc.ChatRepository.SaveImageMetadata(metadata)
}
