package routes

import (
    "github.com/gorilla/mux"
    "github.com/ipastushenko/simple-chat/controllers/websocket"
)

func appendAnonymousWebSocketRouter(router *mux.Router) {
    router.Handle("/websocket", websocket.NewWebSocketHandler())
}
