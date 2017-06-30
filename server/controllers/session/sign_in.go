package session

import(
    "net/http"
    "encoding/json"
    "log"
    "github.com/ipastushenko/simple-chat/server/models"
    "github.com/ipastushenko/simple-chat/server/services/session"
)

type SignInHandler struct {
    sessionService session.ISessionService
}

type signInResponse struct {
    AuthToken string `json:"auth_token"`
}

func NewSignInHandler() *SignInHandler {
    return &SignInHandler{
        sessionService: session.NewSessionService(),
    }
}

func (handler *SignInHandler) ServeHTTP(
    responseWriter http.ResponseWriter,
    request *http.Request,
) {
    decoder := json.NewDecoder(request.Body)
    defer request.Body.Close()
    user := &models.User{}
    err := decoder.Decode(user)
    if err != nil {
        log.Println(err.Error())
        responseWriter.WriteHeader(http.StatusUnauthorized)
        return
    }
    token, ok := handler.sessionService.SignIn(user)
    if !ok {
        responseWriter.WriteHeader(http.StatusUnauthorized)
        return
    }
    tokenString, ok := token.(string)
    if !ok {
        responseWriter.WriteHeader(http.StatusUnauthorized)
        return
    }

    response := &signInResponse{AuthToken: tokenString}
    json.NewEncoder(responseWriter).Encode(response)
}
