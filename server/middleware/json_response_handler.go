package middleware

import "net/http"

type JsonResponseHandler struct {}

func NewJsonResponseHandler() *JsonResponseHandler {
    return &JsonResponseHandler{}
}

func (handler *JsonResponseHandler) ServeHTTP(
    responseWriter http.ResponseWriter,
    request *http.Request,
    next http.HandlerFunc,
) {
    responseWriter.Header().Set("Content-Type", "application/json")
    next(responseWriter, request)
}
