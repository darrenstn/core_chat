package service

import "core_chat/application/websocket/service"

type WebSocketManager interface {
	AddClient(identifier string, conn service.WebSocketConn)
	Send(identifier string, data []byte) error
	IsOnline(identifier string) bool
}
