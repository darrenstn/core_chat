package service

type WebSocketManager interface {
	UpdateToken(identifier, newToken string)
	CloseConnection(identifier string)
}
