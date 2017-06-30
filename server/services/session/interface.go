package session

import "github.com/ipastushenko/simple-chat/server/services/auth"

type ISessionService interface {
    SignIn(auth.IUserCredentials) (interface{}, bool)
    SignOut(interface{}) error
}
