package service

type WebSocketRouter interface {
	Route(identifier, token string, data []byte)
}
