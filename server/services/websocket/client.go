package websocket

import (
    "github.com/gorilla/websocket"
    "log"
    "time"
)

const (
    writeTimeout = 10 * time.Second
    pongTimeout = 30 * time.Second
    pingPeriod = (pongTimeout * 9) / 10
)

type Message struct{
    Id int `json:"id"`
    Event string `json:"event"`
}

type Client struct {
    wss *WebSocketService
    conn *websocket.Conn
    Send chan interface{}
    UserId int
}

func NewClient(wss *WebSocketService, conn *websocket.Conn, userId int) *Client {
    return &Client{
        wss: wss,
        conn: conn,
        Send: make(chan interface{}),
        UserId: userId,
    }
}

func (client *Client) InitReadHandler() {
    defer func() {
        client.wss.unregister <- client
        client.conn.Close()
        close(client.Send)
    }()

    client.conn.SetReadDeadline(time.Now().Add(pongTimeout))
    client.conn.SetPongHandler(
        func(string) error {
            client.conn.SetReadDeadline(time.Now().Add(pongTimeout))
            log.Println("pong")
            return nil
        },
    )

    message := Message{}
    for {
        err := client.conn.ReadJSON(&message)
        if err != nil {
            if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
                log.Printf("error: %v", err)
            }
            break
        }
        log.Println(message)
        client.wss.Broadcast(&BroadcastInfo{message: message})
    }
}

func (client *Client) InitWriteHandler() {
    ticker := time.NewTicker(pingPeriod)
    defer func() {
        ticker.Stop()
        client.conn.Close()
    }()

    for {
        select {
        case message, ok := <- client.Send:
            client.conn.SetWriteDeadline(time.Now().Add(writeTimeout))
            if !ok {
                client.conn.WriteMessage(websocket.CloseMessage, []byte{})
                return
            }
            err := client.conn.WriteJSON(message)
            if err != nil {
                return
            }
        case <- ticker.C:
            client.conn.SetWriteDeadline(time.Now().Add(writeTimeout))
            err := client.conn.WriteMessage(websocket.PingMessage, []byte{});
            log.Println("ping")
            if err != nil {
                return
            }
        }
    }
}
