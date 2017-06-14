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

type Data struct {
    Id int `json:"id"`
    Test interface{} `json:"test"`
}

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
    data := Data{}
    for {
        err := connection.ReadJSON(&data)

        if err != nil {
            log.Println("read: ", err.Error())
            break
        }

        log.Printf("recv: ", data)
        err = connection.WriteJSON(data)
        if err != nil {
            log.Println("write: ", err.Error())
            break
        }
    }
}
