package routes

import (
    "github.com/gorilla/mux"
    "github.com/ipastushenko/simple-chat/server/controllers/websocket"
)

func appendAuthWebSocketRouter(router *mux.Router) {
    router.Handle("/websocket", websocket.NewWebSocketHandler()).Methods("GET")
}
