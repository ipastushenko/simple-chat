package session

import (
    "sync"
    "github.com/ipastushenko/simple-chat/models"
    "github.com/ipastushenko/simple-chat/services/auth"
    "github.com/ipastushenko/simple-chat/services/token"
)

type SessionService struct {
    authService auth.IAuthService
    tokenService token.ITokenService
}

var (
    service *SessionService
    once sync.Once
)

func NewSessionService() ISessionService {
    once.Do(func() {
        service = &SessionService{
            authService: auth.NewAuthService(),
            tokenService: token.NewJWTService(),
        }
    })

    return service
}

func (service *SessionService) SignIn(
    credentials auth.IUserCredentials,
) (interface{}, bool) {
    rawUser, ok := service.authService.Authenticate(credentials)
    if !ok {
        return nil, false
    }
    user, ok := rawUser.(*models.User)
    if !ok {
        return nil, false
    }
    claims := &token.JWTClaims{UserId: user.Id}
    token, err := service.tokenService.GenerateToken(claims)
    if err != nil {
        return nil, false
    }

    return token, true
}

func (service *SessionService) SignOut(token interface{}) error {
    return nil
}
