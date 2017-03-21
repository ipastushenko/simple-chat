package session

import(
    "net/http"
    "encoding/json"
    "github.com/ipastushenko/simple-chat/services/session"
)

//TODO: temp signout handler response
type SignOutHandler struct {
    sessionService session.ISessionService
}

type signOutResponse struct {
    UserId int `json:"user_id"`
}

func NewSignOutHandler() *SignOutHandler {
    return &SignOutHandler{
        sessionService: session.NewSessionService(),
    }
}

func (handler *SignOutHandler) ServeHTTP(
    responseWriter http.ResponseWriter,
    request *http.Request,
) {
    userId, ok := request.Context().Value("user_id").(int);
    if !ok {
        responseWriter.WriteHeader(http.StatusUnauthorized)
        return
    }

    response := &signOutResponse{UserId: userId}
    json.NewEncoder(responseWriter).Encode(response)
}
