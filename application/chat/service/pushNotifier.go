package service

import "core_chat/application/pushnotification/dto"

type PushNotifier interface {
	Send(input dto.SendNotificationInput) error
}
