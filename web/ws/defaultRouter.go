package ws

import (
	chatapp "core_chat/application/chat/usecase"
	notificationapp "core_chat/application/notification/usecase"
	"core_chat/application/websocket/service"
	"core_chat/web/ws/dto"
	"core_chat/web/ws/mapper"
	"encoding/json"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

type DefaultRouter struct {
	SendMessageUC    chatapp.SendMessageUseCase
	SendImageUC      chatapp.SendImageUseCase
	JoinChatRoomUC   chatapp.JoinChatRoomUseCase
	LeaveChatRoomUC  chatapp.LeaveChatRoomUseCase
	ServerResponseUC notificationapp.ServerResponseUseCase
}

func NewDefaultRouter(sendMessageUC *chatapp.SendMessageUseCase, sendImageUC *chatapp.SendImageUseCase, joinChatRoomUC *chatapp.JoinChatRoomUseCase, leaveChatRoomUC *chatapp.LeaveChatRoomUseCase, serverResponseUC *notificationapp.ServerResponseUseCase) service.WebSocketRouter {
	return &DefaultRouter{
		SendMessageUC:    *sendMessageUC,
		SendImageUC:      *sendImageUC,
		JoinChatRoomUC:   *joinChatRoomUC,
		LeaveChatRoomUC:  *leaveChatRoomUC,
		ServerResponseUC: *serverResponseUC,
	}
}

func (r *DefaultRouter) Route(identifier, tokenStr string, data []byte) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte("test"), nil
	})

	if err != nil || !token.Valid {
		r.ServerResponseUC.Execute(identifier, "Token not valid", "server_error_response")
		return
	}

	var base struct {
		Type string `json:"type"` // e.g., "chat_message", "typing", etc.
	}
	_ = json.Unmarshal(data, &base)

	switch base.Type {
	case "chat_message":
		var input dto.SendMessageInput
		_ = json.Unmarshal(data, &input)
		input.Sender = identifier
		appInput := mapper.ToAppSendMessageInput(input)
		err := r.SendMessageUC.Execute(appInput)
		if err != nil {
			r.ServerResponseUC.Execute(identifier, "Send message failed: "+err.Error(), "server_error_response")
			return
		}
		r.ServerResponseUC.Execute(identifier, "Send message success", "server_success_response")

	case "image":
		var input dto.SendMessageInput
		_ = json.Unmarshal(data, &input)
		input.Sender = identifier
		appInput := mapper.ToAppSendMessageInput(input)

		imagePath := os.Getenv("DEFAULT_CHAT_IMAGE_DIR")

		imageName := ""

		parts := strings.Split(appInput.Payload, "=")
		if len(parts) > 1 {
			imageName = parts[1]
			imageChatUrl := os.Getenv("IMAGE_CHAT_URL")
			baseChatImgUrl := imageChatUrl + "?image_name"
			if baseChatImgUrl == parts[0] { //Check if the base url matched with base url on the payload
				err := r.SendImageUC.Execute(appInput, imagePath, imageName)
				if err != nil {
					r.ServerResponseUC.Execute(identifier, "Send message failed: "+err.Error(), "server_error_response")
					return
				}
				r.ServerResponseUC.Execute(identifier, "Send message success", "server_success_response")
			}
		}

	case "join_room":
		var input struct {
			Person string `json:"person"`
		}
		_ = json.Unmarshal(data, &input)
		err := r.JoinChatRoomUC.Execute(identifier, input.Person)
		if err != nil {
			r.ServerResponseUC.Execute(identifier, "Join room failed: "+err.Error(), "server_error_response")
			return
		}
		r.ServerResponseUC.Execute(identifier, "Join room success", "server_success_response")

	case "leave_room":
		var input struct {
			Person string `json:"person"`
		}
		_ = json.Unmarshal(data, &input)
		err := r.LeaveChatRoomUC.Execute(identifier, input.Person)
		if err != nil {
			r.ServerResponseUC.Execute(identifier, "Leave room failed: "+err.Error(), "server_error_response")
			return
		}
		r.ServerResponseUC.Execute(identifier, "Leave room success", "server_success_response")
	}
}
