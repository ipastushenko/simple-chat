package routes

import (
    "fmt"
    "github.com/gorilla/mux"
    "github.com/urfave/negroni"
    "github.com/ipastushenko/simple-chat/middleware"
    "github.com/ipastushenko/simple-chat/settings"
)

func Router() *mux.Router {
    config := settings.GetInstance()
    apiPath := fmt.Sprintf("/api/%v",config.ApiVersion)

    router := mux.NewRouter()
    authRouter := mux.NewRouter().PathPrefix(apiPath).Subrouter()
    anonymousRouter := router.PathPrefix(apiPath).Subrouter()

    appendAuthAuthRouter(authRouter)
    appendAnonymousAuthRouter(anonymousRouter)

    router.PathPrefix(apiPath).Handler(
        negroni.New(
            negroni.HandlerFunc(middleware.RequireAuth),
            negroni.Wrap(authRouter),
        ),
    )

    return router
}
