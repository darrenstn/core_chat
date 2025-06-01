package serviceimpl

import (
	"core_chat/application/websocket/service"
	"errors"
	"sync"

	"github.com/gorilla/websocket"
)

type connection struct {
	conn *websocket.Conn
	mu   sync.Mutex // Prevent concurrent writes
}

type WebSocketManagerImpl struct {
	clients   map[string]*connection
	chatRooms map[string]map[string]bool // chatID -> userID set
	lock      sync.RWMutex
	router    service.WebSocketRouter
}

func NewWebSocketManager(router service.WebSocketRouter) *WebSocketManagerImpl {
	return &WebSocketManagerImpl{
		clients: make(map[string]*connection),
		router:  router,
	}
}

func (wsm *WebSocketManagerImpl) AddClient(identifier string, conn *websocket.Conn) {
	wsm.lock.Lock()
	wsm.clients[identifier] = &connection{conn: conn}
	wsm.lock.Unlock()
	go wsm.listen(identifier, conn)
}

func (wsm *WebSocketManagerImpl) Send(identifier string, data []byte) error {
	wsm.lock.RLock()
	client, ok := wsm.clients[identifier]
	wsm.lock.RUnlock()
	if !ok {
		return errors.New("user is offline")
	}

	client.mu.Lock()
	defer client.mu.Unlock()
	return client.conn.WriteMessage(websocket.TextMessage, data)
}

func (wsm *WebSocketManagerImpl) IsOnline(identifier string) bool {
	wsm.lock.RLock()
	defer wsm.lock.RUnlock()
	_, ok := wsm.clients[identifier]
	return ok
}

func (wsm *WebSocketManagerImpl) listen(identifier string, conn *websocket.Conn) {
	defer func() {
		conn.Close()
		wsm.lock.Lock()
		delete(wsm.clients, identifier)
		wsm.lock.Unlock()
	}()

	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			break
		}

		if wsm.router != nil {
			wsm.router.Route(identifier, data)
		}
	}
}

func (wsm *WebSocketManagerImpl) JoinRoom(chatID, userID string) {
	wsm.lock.Lock()
	defer wsm.lock.Unlock()
	if wsm.chatRooms[chatID] == nil {
		wsm.chatRooms[chatID] = make(map[string]bool)
	}
	wsm.chatRooms[chatID][userID] = true
}

func (wsm *WebSocketManagerImpl) LeaveRoom(chatID, userID string) {
	wsm.lock.Lock()
	defer wsm.lock.Unlock()
	if room, ok := wsm.chatRooms[chatID]; ok {
		delete(room, userID)
		if len(room) == 0 {
			delete(wsm.chatRooms, chatID)
		}
	}
}

func (wsm *WebSocketManagerImpl) IsUserInRoom(chatID, userID string) bool {
	wsm.lock.RLock()
	defer wsm.lock.RUnlock()
	return wsm.chatRooms[chatID][userID]
}
