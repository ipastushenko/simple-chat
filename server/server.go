package main

import (
    "log"
    "github.com/ipastushenko/simple-chat/settings"
    "github.com/ipastushenko/simple-chat/routes"
    "github.com/ipastushenko/simple-chat/middleware"
    "github.com/urfave/negroni"
    "net/http"
    "fmt"
)

func main () {
    config := settings.GetInstance()
    router := routes.Router()
    serverMiddleware := negroni.New()
    serverMiddleware.Use(negroni.NewRecovery())
    serverMiddleware.Use(negroni.NewLogger())
    serverMiddleware.Use(negroni.HandlerFunc(middleware.JsonResponse))
    serverMiddleware.UseHandler(router)
    log.Println(config.Env)
    log.Println(config.Server.Port)
    http.ListenAndServe(fmt.Sprintf(":%v", config.Server.Port), serverMiddleware)
}
