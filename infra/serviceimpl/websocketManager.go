package serviceimpl

import (
	authservice "core_chat/application/authentication/service"
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
	_ authservice.WebSocketManager         = (*WebSocketManagerImpl)(nil)
)

type connection struct {
	conn   wsservice.WebSocketConn
	mu     sync.Mutex // Prevent concurrent writes
	token  string
	once   sync.Once
	closed bool
}

func (c *connection) getToken() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.token
}

func (c *connection) setToken(newToken string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.token = newToken
}

func (c *connection) SafeClose() {
	c.once.Do(func() {
		c.mu.Lock()
		defer c.mu.Unlock()
		if !c.closed {
			c.conn.Close()
			c.closed = true
		}
	})
}

func (c *connection) SafeWrite(data []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.closed {
		return errors.New("connection is closed")
	}
	return c.conn.WriteMessage(websocket.TextMessage, data)
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

func NewAuthenticationWebSocketManager(_ wsservice.WebSocketRouter) authservice.WebSocketManager {
	return GetWebSocketManager()
}

func (wsm *WebSocketManagerImpl) SetRouter(router wsservice.WebSocketRouter) {
	wsm.lock.Lock()
	defer wsm.lock.Unlock()
	wsm.router = router
}

func (wsm *WebSocketManagerImpl) AddClient(identifier string, conn wsservice.WebSocketConn, token string) {
	var oldClient *connection
	client := &connection{conn: conn, token: token}

	wsm.lock.Lock()
	if oc, exists := wsm.clients[identifier]; exists {
		oldClient = oc
	}
	wsm.clients[identifier] = client
	wsm.lock.Unlock()

	if oldClient != nil {
		oldClient.SafeClose()
	}

	go wsm.listen(identifier, client)
}

func (wsm *WebSocketManagerImpl) Send(identifier string, data []byte) error {
	wsm.lock.RLock()
	client, ok := wsm.clients[identifier]
	wsm.lock.RUnlock()
	if !ok {
		return errors.New("user is offline")
	}

	return client.SafeWrite(data)
}

func (wsm *WebSocketManagerImpl) IsOnline(identifier string) bool {
	wsm.lock.RLock()
	defer wsm.lock.RUnlock()
	_, ok := wsm.clients[identifier]
	return ok
}

func (wsm *WebSocketManagerImpl) listen(identifier string, client *connection) {
	defer wsm.CloseConnection(identifier)

	for {
		_, data, err := client.conn.ReadMessage()
		if err != nil {
			break
		}

		wsm.lock.RLock()
		router := wsm.router
		wsm.lock.RUnlock()

		if router != nil {
			token := client.getToken()

			router.Route(identifier, token, data)
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
	if room, ok := wsm.chatRooms[chatID]; ok {
		return room[identifier]
	}
	return false
}

func (wsm *WebSocketManagerImpl) UpdateToken(identifier, newToken string) {
	wsm.lock.Lock()
	defer wsm.lock.Unlock()

	client, ok := wsm.clients[identifier]
	if ok {
		client.setToken(newToken)
	}
}

func (wsm *WebSocketManagerImpl) CloseConnection(identifier string) {
	wsm.lock.Lock()
	client, exists := wsm.clients[identifier]
	if exists {
		delete(wsm.clients, identifier)

		// Clean up user from all chat rooms
		for chatID := range wsm.chatRooms {
			delete(wsm.chatRooms[chatID], identifier)
			if len(wsm.chatRooms[chatID]) == 0 {
				delete(wsm.chatRooms, chatID)
			}
		}
	}
	wsm.lock.Unlock()

	if exists {
		client.SafeClose()
	}
}
