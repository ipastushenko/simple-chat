package controllers

import (
    "net/http"
    "encoding/json"
)

type signInResponse struct {
    AuthToken string `json:"auth_token"`
}

type signOutResponse struct {
    Success bool `json:"success"`
}

func SignIn(rw http.ResponseWriter, req *http.Request) {
    response := signInResponse{AuthToken: "TEST_TOKEN"}
    json.NewEncoder(rw).Encode(response)
}

func SignOut(rw http.ResponseWriter, req *http.Request) {
    response := signOutResponse{Success: true}
    json.NewEncoder(rw).Encode(response)
}
