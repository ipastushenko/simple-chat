package middleware

import (
    "net/http"
    "context"
    "github.com/ipastushenko/simple-chat/services/auth"
)

const (
    authTokenParamName string = "auth_token"
)

func RequireAuth(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
    query := req.URL.Query()
    auth_token := query.Get(authTokenParamName)
    if auth_token == "" {
        rw.WriteHeader(http.StatusUnauthorized)
        return
    }
    if token, ok := auth.VerifyToken(auth_token); ok {
        newReq := req
        if claims, ok := token.Claims.(*auth.TokenClaims); ok {
            ctx := req.Context()
            newContext := context.WithValue(
                ctx,
                auth.UserIdContextName,
                claims.UserId,
            )
            newReq = req.WithContext(newContext)
        }
        next(rw, newReq)
    } else {
        rw.WriteHeader(http.StatusUnauthorized)
    }
}
