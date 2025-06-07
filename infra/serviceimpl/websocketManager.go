package serviceimpl

import (
	chatservice "core_chat/application/chat/service"
	notificationservice "core_chat/application/notification/service"
	postservice "core_chat/application/post/service"
	wsservice "core_chat/application/websocket/service"
	"errors"
	"sort"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	_ wsservice.WebSocketManager           = (*WebSocketManagerImpl)(nil)
	_ chatservice.WebSocketManager         = (*WebSocketManagerImpl)(nil)
	_ postservice.WebSocketManager         = (*WebSocketManagerImpl)(nil)
	_ notificationservice.WebSocketManager = (*WebSocketManagerImpl)(nil)
)

type connection struct {
	conn wsservice.WebSocketConn
	mu   sync.Mutex // Prevent concurrent writes
}

type WebSocketManagerImpl struct {
	clients   map[string]*connection
	chatRooms map[string]map[string]bool // chatID -> userID set
	lock      sync.RWMutex
	router    wsservice.WebSocketRouter
}

func NewWebSocketManager(_ wsservice.WebSocketRouter) wsservice.WebSocketManager {
	return GetWebSocketManager()
}

var (
	once      sync.Once
	singleton *WebSocketManagerImpl
)

// this must be called once early, e.g. during main/init/DI
func InitWebSocketManagerImpl(router wsservice.WebSocketRouter) {
	once.Do(func() {
		singleton = &WebSocketManagerImpl{
			clients:   make(map[string]*connection),
			chatRooms: make(map[string]map[string]bool),
			router:    router,
		}
	})
}

func GetWebSocketManager() *WebSocketManagerImpl {
	if singleton == nil {
		panic("WebSocketManager is not initialized. Call InitWebSocketManager first.")
	}
	return singleton
}

// Factory methods
func NewChatWebSocketManager(_ wsservice.WebSocketRouter) chatservice.WebSocketManager {
	return GetWebSocketManager()
}

func NewPostWebSocketManager(_ wsservice.WebSocketRouter) postservice.WebSocketManager {
	return GetWebSocketManager()
}

func NewNotificationWebSocketManager(_ wsservice.WebSocketRouter) notificationservice.WebSocketManager {
	return GetWebSocketManager()
}

func (wsm *WebSocketManagerImpl) SetRouter(router wsservice.WebSocketRouter) {
	wsm.router = router
}

func (wsm *WebSocketManagerImpl) AddClient(identifier string, conn wsservice.WebSocketConn) {
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

func (wsm *WebSocketManagerImpl) listen(identifier string, conn wsservice.WebSocketConn) {
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

func (wsm *WebSocketManagerImpl) JoinRoom(chatID, identifier string) {
	wsm.lock.Lock()
	defer wsm.lock.Unlock()
	if wsm.chatRooms[chatID] == nil {
		wsm.chatRooms[chatID] = make(map[string]bool)
	}
	wsm.chatRooms[chatID][identifier] = true
}

func (wsm *WebSocketManagerImpl) GenerateChatID(person1, person2 string) string {
	chatID := []string{person1, person2}
	sort.Strings(chatID)
	return chatID[0] + ":" + chatID[1]
}

func (wsm *WebSocketManagerImpl) LeaveRoom(chatID, identifier string) {
	wsm.lock.Lock()
	defer wsm.lock.Unlock()
	if room, ok := wsm.chatRooms[chatID]; ok {
		delete(room, identifier)
		if len(room) == 0 {
			delete(wsm.chatRooms, chatID)
		}
	}
}

func (wsm *WebSocketManagerImpl) IsPersonInRoom(chatID, identifier string) bool {
	wsm.lock.RLock()
	defer wsm.lock.RUnlock()
	return wsm.chatRooms[chatID][identifier]
}
