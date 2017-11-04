package routes

import (
    "fmt"
    "github.com/gorilla/mux"
    "github.com/urfave/negroni"
    "github.com/ipastushenko/simple-chat/server/middleware"
    "github.com/ipastushenko/simple-chat/server/settings"
)

func Router() *mux.Router {
    config := settings.GetInstance()
    apiPath := fmt.Sprintf("/api/%v", config.GetString("api_version"))

    router := mux.NewRouter()
    authRouter := mux.NewRouter().PathPrefix(apiPath).Subrouter()
    anonymousRouter := router.PathPrefix(apiPath).Subrouter()

    appendAuthAuthRouter(authRouter)
    appendAuthWebSocketRouter(authRouter)

    appendAnonymousAuthRouter(anonymousRouter)

    router.PathPrefix(apiPath).Handler(
        negroni.New(
            middleware.NewRequiredAuthHandler(),
            negroni.Wrap(authRouter),
        ),
    )

    return router
}
