package routes

import (
    "github.com/gorilla/mux"
    "github.com/ipastushenko/simple-chat/settings"
    "fmt"
)

func Router() *mux.Router {
    config := settings.GetInstance()
    router := mux.NewRouter()
    router = router.PathPrefix(fmt.Sprintf("/api/%v",config.ApiVersion)).Subrouter()

    AuthRouter(router)

    return router
}
