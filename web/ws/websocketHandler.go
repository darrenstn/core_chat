package ws

import (
	"core_chat/application/websocket/service"
	"core_chat/infra/serviceimpl/adapter"
	"net/http"

	"github.com/gorilla/websocket"
)

func HandleWebSocketConn(w http.ResponseWriter, r *http.Request, identifier string, wsManager service.WebSocketManager) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not upgrade to WebSocket", http.StatusBadRequest)
		return
	}

	var wsConn service.WebSocketConn = &adapter.GorillaConn{Conn: conn}

	wsManager.AddClient(identifier, wsConn)
}
