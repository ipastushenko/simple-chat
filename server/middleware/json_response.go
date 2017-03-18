package middleware

import "net/http"

func JsonResponse(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
    rw.Header().Set("Content-Type", "application/json")
    next(rw, req)
}
