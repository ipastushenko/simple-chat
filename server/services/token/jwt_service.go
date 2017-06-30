package token

import (
    "net/http"
    "crypto/rsa"
    "sync"
    "time"
    "io/ioutil"
    "encoding/pem"
    "crypto/x509"
    "context"
    jwt "github.com/dgrijalva/jwt-go"
    "github.com/ipastushenko/simple-chat/server/settings"
)

type JWTClaims struct {
    jwt.StandardClaims
    Info interface{}
}

type JWTService struct {
    secret *jwtSecret
}

type jwtSecret struct {
    secretKey *rsa.PrivateKey
    verifyKey *rsa.PublicKey
}

const (
    userIdContextName string = "info"
)

var (
    service *JWTService
    once sync.Once
)

func NewJWTService() ITokenService {
    return getInstance()
}

func (service *JWTService) GenerateToken(info interface{}) (string, error) {
    config := settings.GetInstance()
    token := jwt.New(jwt.SigningMethodRS256)

    timeNow := time.Now()
    expirationTime := time.Duration(config.Server.TokenExpiration)
    exp := timeNow.Add(expirationTime * time.Minute).Unix()
    iat := timeNow.Unix()
    token.Claims = &JWTClaims{
        Info: info,
        StandardClaims: jwt.StandardClaims{IssuedAt: iat, ExpiresAt: exp},
    }

    return token.SignedString(service.secret.secretKey)
}

func (service *JWTService) ParseToken(tokenString string) (interface{}, bool) {
    token, err := jwt.ParseWithClaims(
        tokenString,
        &JWTClaims{},
        func (*jwt.Token) (interface{}, error) {
            return service.secret.verifyKey, nil
        },
    )
    if err == nil && token.Valid {
        return token, true
    }

    return nil, false
}

func (service *JWTService) UpdateRequestContext(
    token interface{},
    request *http.Request,
) *http.Request {
    jwtToken, ok := token.(*jwt.Token)
    if !ok {
        return request
    }
    claims, ok := jwtToken.Claims.(*JWTClaims)
    if !ok {
        return request
    }

    currentContext := request.Context()
    newContext := context.WithValue(
        currentContext,
        userIdContextName,
        claims.Info,
    )

    return request.WithContext(newContext)
}

func (service *JWTService) GetRequestContextInfo(
    request *http.Request,
) interface{} {
    return request.Context().Value(userIdContextName)
}

//TODO: need to implement
func (service *JWTService) RefreshToken(string) (string, error) {
    return "", nil
}

//TODO: need to implement
func (service *JWTService) RevokeToken(string) error {
    return nil
}

//TODO: need to implement
func (service *JWTService) RevokeOtherTokens(string) error {
    return nil
}

func getInstance() ITokenService {
    once.Do(func() {
        service = &JWTService{secret: &jwtSecret{}}
        service.loadSecretKey()
        service.loadVerifyKey()
    })

    return service
}

func (service *JWTService) loadSecretKey() {
    config := settings.GetInstance()
    secretKey, err := ioutil.ReadFile(config.Server.SecretKeyPath)
    if err != nil {
        panic(err)
    }
    secretData, _ := pem.Decode(secretKey)
    importedSecretKey, err := x509.ParsePKCS1PrivateKey(secretData.Bytes)
    if err != nil {
        panic(err)
    }

    service.secret.secretKey = importedSecretKey
}

func (service *JWTService) loadVerifyKey() {
    config := settings.GetInstance()
    verifyKey, err := ioutil.ReadFile(config.Server.VerifyKeyPath)
    if err != nil {
        panic(err)
    }
    verifyData, _ := pem.Decode(verifyKey)
    rawVerifyKey, err := x509.ParsePKIXPublicKey(verifyData.Bytes)
    if err != nil {
        panic(err)
    }
    importedVerifyKey, ok := rawVerifyKey.(*rsa.PublicKey)
    if !ok {
        panic("rsa public key cast")
    }

    service.secret.verifyKey = importedVerifyKey
}
