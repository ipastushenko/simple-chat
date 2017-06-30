package websocket

import(
    "net/http"
    "log"
    "github.com/gorilla/websocket"
    wss "github.com/ipastushenko/simple-chat/server/services/websocket"
    "github.com/ipastushenko/simple-chat/server/services/token"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

type WebSocketHandler struct {
    webSocketService wss.IWebSocketService
    tokenService token.ITokenService
}

func NewWebSocketHandler() *WebSocketHandler {
    return &WebSocketHandler{
        webSocketService: wss.NewWebSocketService(),
        tokenService: token.NewJWTService(),
    }
}

func (handler *WebSocketHandler) ServeHTTP(
    responseWriter http.ResponseWriter,
    request *http.Request,
) {
    connection, err := upgrader.Upgrade(responseWriter, request, nil)
    if err != nil {
        log.Println(err.Error())
        responseWriter.WriteHeader(http.StatusInternalServerError)
        return
    }
    info := handler.tokenService.GetRequestContextInfo(request)
    handler.webSocketService.InitConnection(connection, info)
}
