package main

import (
    "log"
    "fmt"
    "net/http"
    "github.com/ipastushenko/simple-chat/settings"
    "github.com/ipastushenko/simple-chat/routes"
    "github.com/ipastushenko/simple-chat/middleware"
    "github.com/ipastushenko/simple-chat/services/auth"
)

func main () {
    config := settings.GetInstance()
    auth.LoadSecretKeys()
    router := routes.Router()
    serverMiddleware := middleware.Middleware()
    serverMiddleware.UseHandler(router)
    log.Printf(
        "%v server has been started on %v port",
        config.Env,
        config.Server.Port,
    )
    log.Fatal(
        http.ListenAndServe(
            fmt.Sprintf(":%v", config.Server.Port),
            serverMiddleware,
        ),
    )
}
