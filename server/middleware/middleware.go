package middleware

import "github.com/urfave/negroni"

func Middleware() *negroni.Negroni {
    serverMiddleware := negroni.New()
    serverMiddleware.Use(negroni.NewRecovery())
    serverMiddleware.Use(negroni.NewLogger())
    serverMiddleware.Use(negroni.HandlerFunc(JsonResponse))

    return serverMiddleware
}
