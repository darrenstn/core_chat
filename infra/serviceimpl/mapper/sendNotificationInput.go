package mapper

import (
	appdto "core_chat/application/notification/dto"
	infradto "core_chat/infra/serviceimpl/dto"
)

func ToInfraSendNotificationInput(input appdto.SendNotificationInput) infradto.SendNotificationInput {
	return infradto.SendNotificationInput{
		Receiver:      input.Receiver,
		Sender:        input.Sender,
		Type:          input.Type,
		Title:         input.Title,
		Body:          input.Body,
		Payload:       input.Payload,
		FirebaseToken: input.FirebaseToken,
	}
}
