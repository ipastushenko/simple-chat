package main

import (
    "log"
    "fmt"
    "net/http"
    "github.com/ipastushenko/simple-chat/server/settings"
    "github.com/ipastushenko/simple-chat/server/routes"
    "github.com/ipastushenko/simple-chat/server/middleware"
)

func main () {
    config := settings.GetInstance()
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
