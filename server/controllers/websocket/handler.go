package websocket

import(
    "net/http"
    "log"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

type WebSocketHandler struct {}

func NewWebSocketHandler() *WebSocketHandler {
    return &WebSocketHandler{}
}

func (handler *WebSocketHandler) ServeHTTP(
    responseWriter http.ResponseWriter,
    request *http.Request,
) {
    connection, err := upgrader.Upgrade(responseWriter, request, nil)

    if err != nil {
        log.Println(err.Error())
        return
    }

    defer connection.Close()
    for {
        messageType, message, err := connection.ReadMessage()

        if err != nil {
            log.Println("read: ", err.Error())
            break
        }

        log.Printf("recv: %s", message)
        err = connection.WriteMessage(messageType, message)
        if err != nil {
            log.Println("write: ", err)
            break
        }
    }
}
