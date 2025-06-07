package routes

import (
	"core_chat/application/chat/usecase"
	"core_chat/web/rest"
	"core_chat/web/rest/dto"
	"core_chat/web/rest/mapper"
	"core_chat/web/util"
	"net/http"
)

type ChatHandler struct {
	GetChatMessageUC usecase.GetChatMessageUseCase
}

func NewChatHandler(getChatMessageUC usecase.GetChatMessageUseCase) *ChatHandler {
	return &ChatHandler{
		GetChatMessageUC: getChatMessageUC,
	}
}

func (h *ChatHandler) GetChatMessage(w http.ResponseWriter, r *http.Request) {
	identifier, ok := util.GetIdentifier(r)
	if !ok {
		rest.SendResponse(w, 401, "Unauthorized: identifier not found")
		return
	}

	msgID := r.FormValue("message_id")
	if msgID == "" {
		rest.SendResponse(w, 400, "Message ID is required")
		return
	}

	message, err := h.GetChatMessageUC.Execute(msgID, identifier)

	if err != nil {
		rest.SendResponse(w, 400, "Get chat message failed: "+err.Error())
	}

	messageResult := mapper.ToMessageResult(message)

	res := dto.ChatMessageResult{
		Status:  200,
		Message: "Get chat message success",
		Data:    *messageResult,
	}

	rest.SendJSON(w, res)
}
