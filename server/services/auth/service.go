package auth

import (
    "sync"
    "golang.org/x/crypto/bcrypt"
    "github.com/ipastushenko/simple-chat/server/models"
)

//TODO: test user
var password, _ = bcrypt.GenerateFromPassword(
    []byte("password"),
    bcrypt.DefaultCost,
)
var testUser = &models.User{
    Username: "test",
    Password: string(password),
    Id: 1,
}

var (
    service *AuthService
    once sync.Once
)

type AuthService struct {}

func NewAuthService() IAuthService {
    once.Do(func() {
        service = &AuthService{}
    })

    return service
}

func (service *AuthService) Authenticate(
    credentials IUserCredentials,
) (interface{}, bool) {
    comparePasswords := bcrypt.CompareHashAndPassword(
        []byte(testUser.Password),
        []byte(credentials.GetPassword()),
    )
    isCompareUsername := credentials.GetUsername() == testUser.Username
    if isCompareUsername && comparePasswords == nil {
        return testUser, true
    }

    return nil, false
}
