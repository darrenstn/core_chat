// infra/adapter/gorillaConn.go
package adapter

import (
	"github.com/gorilla/websocket"
)

type GorillaConn struct {
	*websocket.Conn
}

func (gc *GorillaConn) ReadMessage() (int, []byte, error) {
	return gc.Conn.ReadMessage()
}

func (gc *GorillaConn) WriteMessage(messageType int, data []byte) error {
	return gc.Conn.WriteMessage(messageType, data)
}

func (gc *GorillaConn) Close() error {
	return gc.Conn.Close()
}
