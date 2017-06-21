package session

import(
    "net/http"
    "encoding/json"
    "github.com/ipastushenko/simple-chat/services/session"
    "github.com/ipastushenko/simple-chat/services/token"
)

//TODO: temp signout handler response
type SignOutHandler struct {
    sessionService session.ISessionService
    tokenService token.ITokenService
}

func NewSignOutHandler() *SignOutHandler {
    return &SignOutHandler{
        sessionService: session.NewSessionService(),
        tokenService: token.NewJWTService(),
    }
}

func (handler *SignOutHandler) ServeHTTP(
    responseWriter http.ResponseWriter,
    request *http.Request,
) {
    info := handler.tokenService.GetRequestContextInfo(request)
    json.NewEncoder(responseWriter).Encode(info)
}
