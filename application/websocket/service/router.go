package service

type WebSocketRouter interface {
	Route(identifier string, data []byte)
}
