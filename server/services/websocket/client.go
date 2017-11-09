package websocket

import (
    "github.com/gorilla/websocket"
    "log"
    "time"
    "encoding/json"
)

const (
    writeTimeout = 10 * time.Second
    pongTimeout = 30 * time.Second
    pingPeriod = (pongTimeout * 9) / 10
)

var events map[string]func()IMessageData

events = {
    "id": func() IMessageData { return new(Data2) },
    "name": func() IMessageData { return new(Data1) },
}

type IMessageData interface {}

type Data1 struct {
    Name string `json:"name"`
}

type Data2 struct {
    Id int `json:"id"`
}

type Message struct {
    Event string `json:"event"`
    Data IMessageData `json:"data"`
}

type RawMessage struct {
    Event string `json:"event"`
    Data json.RawMessage `json:"data"`
}

func (message *Message) UnmarshalJSON(b []byte) error {
    rawMessage := RawMessage{}
    err := json.Unmarshal(b, &rawMessage)
    if err != nil {
        return err
    }
    data := events[rawMessage.Event]()
    err = json.Unmarshal(rawMessage.Data, &data)
    if err != nil {
        return err
    }

    message.Event = rawMessage.Event
    message.Data = data

    return nil
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
