package routes

import (
    "github.com/gorilla/mux"
    "github.com/ipastushenko/simple-chat/controllers"
)

func AuthRouter(router *mux.Router) {
    router.HandleFunc("/auth/sign_in", controllers.SignIn).Methods("POST")
    router.HandleFunc("/auth/sign_out", controllers.SignOut).Methods("GET")
}
