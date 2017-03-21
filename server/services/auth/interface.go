package auth

type IUserCredentials interface {
    GetUsername() string
    GetPassword() string
}

type IAuthService interface {
    Authenticate(IUserCredentials) (interface{}, bool)
}
