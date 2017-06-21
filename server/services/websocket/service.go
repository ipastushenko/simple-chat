package websocket

import (
    "sync"
    "log"
    "github.com/gorilla/websocket"
)

type WebSocketService struct {}

//TODO: temp implementation
type Data struct {
    Id int `json:"id"`
}

var (
    service *WebSocketService
    once sync.Once
)

func NewWebSocketService() IWebSocketService {
    once.Do(func () {
        service = &WebSocketService{}
    })

    return service
}

func (service *WebSocketService) InitConnection(
    connection *websocket.Conn,
    userData interface{},
) {
    if data, ok := userData.(map[string]interface{}); ok {
        log.Println(data["user_id"])
    }

    defer connection.Close()
    data := Data{}
    //TODO: temp implementation
    for {
        err := connection.ReadJSON(&data)
        if err != nil {
            log.Println("read: ", err.Error())
            break
        }

        err = connection.WriteJSON(data)
        if err != nil {
            log.Println("write: ", err.Error())
            break
        }
    }
}
