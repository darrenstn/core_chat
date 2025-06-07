package mapper

import (
	appdto "core_chat/application/chat/dto"
	infradto "core_chat/infra/serviceimpl/dto"
)

func ToInfraSendMessageInput(input appdto.SendMessageInput) infradto.SendMessageInput {
	return infradto.SendMessageInput{
		Receiver:      input.Receiver,
		Sender:        input.Sender,
		Type:          input.Type,
		Title:         input.Title,
		Body:          input.Body,
		Payload:       input.Payload,
		FirebaseToken: input.FirebaseToken,
	}
}
