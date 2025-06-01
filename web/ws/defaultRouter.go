package ws

import "encoding/json"

type DefaultRouter struct {
	SendMessageUC chat.SendMessageUseCase
}

func (r *DefaultRouter) Route(identifier string, data []byte) {
	var base struct {
		Type string `json:"type"` // e.g., "chat_message", "typing", etc.
	}
	_ = json.Unmarshal(data, &base)

	switch base.Type {
	case "chat_message":
		var input chat.SendMessageInput
		_ = json.Unmarshal(data, &input)
		r.SendMessageUC.Execute(input)
	}
}
