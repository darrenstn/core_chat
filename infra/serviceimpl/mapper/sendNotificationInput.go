package mapper

import (
	appchatdto "core_chat/application/chat/dto"
	appnotificationdto "core_chat/application/pushnotification/dto"
)

func ToSendNotificationInput(input appchatdto.SendMessageInput) appnotificationdto.SendNotificationInput {
	return appnotificationdto.SendNotificationInput{
		Receiver:      input.Receiver,
		Sender:        input.Sender,
		Type:          input.Type,
		Title:         input.Title,
		Body:          input.Body,
		Payload:       input.Payload,
		FirebaseToken: input.FirebaseToken,
	}
}
