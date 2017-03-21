package routes

import (
    "github.com/gorilla/mux"
    "github.com/ipastushenko/simple-chat/controllers/session"
)

func appendAuthAuthRouter(router *mux.Router) {
    router.Handle("/auth/sign_out", session.NewSignOutHandler()).Methods("GET")
}

func appendAnonymousAuthRouter(router *mux.Router) {
    router.Handle("/auth/sign_in", session.NewSignInHandler()).Methods("POST")
}
