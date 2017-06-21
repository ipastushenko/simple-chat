package websocket

import (
    "github.com/gorilla/websocket"
)

type IWebSocketService interface {
    InitConnection(*websocket.Conn, interface {})
}
