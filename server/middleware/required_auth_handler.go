package middleware

import (
    "net/http"
    "github.com/ipastushenko/simple-chat/services/token"
)

const (
    authTokenParamName string = "auth_token"
)

type RequiredAuthHandler struct {
    TokenService token.ITokenService
}

func NewRequiredAuthHandler() *RequiredAuthHandler {
    return &RequiredAuthHandler{
        TokenService: token.NewJWTService(),
    }
}

func (handler *RequiredAuthHandler) ServeHTTP (
    responseWriter http.ResponseWriter,
    request *http.Request,
    next http.HandlerFunc,
) {
    authToken := request.Header.Get(authTokenParamName)
    if authToken == "" {
        query := request.URL.Query()
        authToken = query.Get(authTokenParamName)
    }
    if authToken == "" {
        responseWriter.WriteHeader(http.StatusUnauthorized)
        return
    }
    if token, ok := handler.TokenService.ParseToken(authToken); ok {
        newRequest := handler.TokenService.UpdateRequestContext(token, request)
        next(responseWriter, newRequest)
    } else {
        responseWriter.WriteHeader(http.StatusUnauthorized)
    }
}
