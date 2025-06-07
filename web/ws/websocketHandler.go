package ws

import (
	"core_chat/application/websocket/service"
	"core_chat/infra/serviceimpl/adapter"
	"core_chat/web/util"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebSocketHandler struct {
	wsManager service.WebSocketManager
}

func NewWebSocketHandler(wsManager service.WebSocketManager) *WebSocketHandler {
	return &WebSocketHandler{
		wsManager: wsManager,
	}
}

func (wsh *WebSocketHandler) HandleWebSocketConn(w http.ResponseWriter, r *http.Request) {
	identifier, ok := util.GetIdentifier(r)

	if !ok {
		http.Error(w, "Unauthorized: identifier not found", 401)
		log.Printf("WebSocket upgrade failed: %v", "Failed get identifier")
		return
	}

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}

	var wsConn service.WebSocketConn = &adapter.GorillaConn{Conn: conn}

	wsh.wsManager.AddClient(identifier, wsConn)
}
