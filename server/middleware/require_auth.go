package middleware

import (
    "net/http"
    "context"
    "github.com/ipastushenko/simple-chat/services/auth"
    jwt "github.com/dgrijalva/jwt-go"
)

func RequireAuth(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
    query := req.URL.Query()
    auth_token := query.Get("auth_token")
    if auth_token != "" {
        token, ok := auth.VerifyToken(auth_token)
        if ok {
            claims, _ := token.Claims.(jwt.MapClaims)
            ctx := req.Context()
            newContext := context.WithValue(ctx, "user_id", int(claims["user_id"].(float64)))
            next(rw, req.WithContext(newContext))
        } else {
            rw.WriteHeader(http.StatusUnauthorized)
        }
    } else {
        rw.WriteHeader(http.StatusUnauthorized)
    }
}
