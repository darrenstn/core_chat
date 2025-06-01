package usecase

import (
	"core_chat/application/notification/dto"
	"core_chat/application/notification/service"
	"errors"
)

type sendNotificationUseCase struct {
	notifier service.NotifierService
}

func NewSendNotificationUseCase(notifier service.NotifierService) *sendNotificationUseCase {
	return &sendNotificationUseCase{
		notifier: notifier,
	}
}

func (uc *sendNotificationUseCase) Execute(input dto.SendNotificationInput) error {
	if input.Receiver == "" || input.Type == "" {
		return errors.New("invalid notification input: missing required fields")
	}
	return uc.notifier.Send(input)
}
