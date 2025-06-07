package service

type WebSocketManager interface {
	Send(identifier string, data []byte) error
}
