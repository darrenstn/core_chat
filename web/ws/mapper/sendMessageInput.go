package mapper

import (
	appdto "core_chat/application/chat/dto"
	wsdto "core_chat/web/ws/dto"
)

func ToAppSendMessageInput(input wsdto.SendMessageInput) appdto.SendMessageInput {
	return appdto.SendMessageInput{
		Receiver:      input.Receiver,
		Sender:        input.Sender,
		Type:          input.Type,
		Title:         input.Title,
		Body:          input.Body,
		Payload:       input.Payload,
		FirebaseToken: input.FirebaseToken,
	}
}
