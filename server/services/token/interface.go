package token

import "net/http"

type ITokenService interface {
    GenerateToken(interface {}) (string, error)
    ParseToken(string) (interface{}, bool)
    UpdateRequestContext(interface{}, *http.Request) *http.Request
    RefreshToken(string) (string, error)
    RevokeToken(string) error
    RevokeOtherTokens(string) error
}
