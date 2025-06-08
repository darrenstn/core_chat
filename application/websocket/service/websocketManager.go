package service

type WebSocketManager interface {
	SetRouter(router WebSocketRouter)
	AddClient(identifier string, conn WebSocketConn, token string)
	Send(identifier string, data []byte) error
	IsOnline(identifier string) bool
	JoinRoom(chatID, userID string)
	GenerateChatID(person1, person2 string) string
	LeaveRoom(chatID, userID string)
	IsPersonInRoom(chatID, userID string) bool
	UpdateToken(identifier, newToken string)
	CloseConnection(identifier string)
}
