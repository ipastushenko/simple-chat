package auth

import (
    "time"
    "crypto/rsa"
    "golang.org/x/crypto/bcrypt"
    jwt "github.com/dgrijalva/jwt-go"
    "github.com/ipastushenko/simple-chat/models"
    "github.com/ipastushenko/simple-chat/settings"
)

type jwtTokenSecret struct {
    secretKey *rsa.PrivateKey
    verifyKey *rsa.PublicKey
}

type TokenClaims struct {
    StandardClaims *jwt.StandardClaims
    UserId int `json:"user_id"`
}

func (claims TokenClaims) Valid () error {
    return claims.StandardClaims.Valid()
}

const (
    UserIdContextName string = "user_id"
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

func Authenticate(user *models.User) (string, bool) {
    comparePasswords := bcrypt.CompareHashAndPassword(
        []byte(testUser.Password),
        []byte(user.Password),
    )

    if user.Username == testUser.Username && comparePasswords == nil {
        token, err := generateToken(testUser)
        if err != nil {
            return "", false
        }
        return token, true
    }

    return "", false
}

func VerifyToken(tokenString string) (*jwt.Token, bool) {
    token, err := jwt.ParseWithClaims(
        tokenString,
        &TokenClaims {},
        func (*jwt.Token) (interface{}, error) {
            return secret.verifyKey, nil
        },
    )
    if err == nil && token.Valid {
        return token, true
    }

    return nil, false
}

func generateToken(user *models.User) (string, error) {
    config := settings.GetInstance()
    token := jwt.New(jwt.SigningMethodRS256)
    timeNow := time.Now()
    exp := timeNow.Add(
            time.Duration(config.Server.TokenExpiration) * time.Minute,
        ).Unix()
    iat := timeNow.Unix()
    token.Claims = &TokenClaims{
        StandardClaims: &jwt.StandardClaims{
            ExpiresAt: exp,
            IssuedAt: iat,
        },
        UserId: user.Id,
    }
    tokenString, err := token.SignedString(secret.secretKey)

    if err != nil {
        panic(err)
        return "", err
    }

    return tokenString, nil
}
