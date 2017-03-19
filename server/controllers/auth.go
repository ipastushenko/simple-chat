package controllers

import (
    "net/http"
    "encoding/json"
    "log"
    "github.com/ipastushenko/simple-chat/models"
    "github.com/ipastushenko/simple-chat/services/auth"
)

type signInResponse struct {
    AuthToken string `json:"auth_token"`
}

type signOutResponse struct {
    UserId int `json:"user_id"`
}

func SignIn(rw http.ResponseWriter, req *http.Request) {
    decoder := json.NewDecoder(req.Body)
    defer req.Body.Close()
    user := &models.User{}
    err := decoder.Decode(user)
    if err != nil {
        log.Println(err.Error())
        rw.WriteHeader(http.StatusUnauthorized)
        return
    }
    token, ok := auth.Authenticate(user)
    if ok {
        json.NewEncoder(rw).Encode(signInResponse{AuthToken: token})
    } else {
        rw.WriteHeader(http.StatusUnauthorized)
    }
}

func SignOut(rw http.ResponseWriter, req *http.Request) {
    response := signOutResponse{req.Context().Value("user_id").(int)}
    json.NewEncoder(rw).Encode(response)
}
