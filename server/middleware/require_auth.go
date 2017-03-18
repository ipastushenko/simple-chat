package middleware

import "net/http"

func RequireAuth(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
    //TODO: add correct auth token
    if req.Header.Get("Authorization") == "TEST_TOKEN" {
        next(rw, req)
    } else {
        rw.WriteHeader(http.StatusUnauthorized)
    }
}
