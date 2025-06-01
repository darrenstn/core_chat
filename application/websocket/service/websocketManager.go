package service

type WebSocketManager interface {
	AddClient(identifier string, conn WebSocketConn)
	Send(identifier string, data []byte) error
	IsOnline(identifier string) bool
	JoinRoom(chatID, userID string)
	LeaveRoom(chatID, userID string)
	IsUserInRoom(chatID, userID string) bool
}
