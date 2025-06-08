package service

import "core_chat/application/websocket/service"

type WebSocketManager interface {
	AddClient(identifier string, conn service.WebSocketConn, token string)
	Send(identifier string, data []byte) error
	IsOnline(identifier string) bool
	JoinRoom(chatID, userID string)
	GenerateChatID(person1, person2 string) string
	LeaveRoom(chatID, userID string)
	IsPersonInRoom(chatID, userID string) bool
}
