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

//TODO: test user
var password, _ = bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
var testUser = &models.User{
    Username: "test",
    Password: string(password),
    Id: 1,
}

func Authenticate(user *models.User) (string, bool) {
    if user.Username == testUser.Username {
        if bcrypt.CompareHashAndPassword([]byte(testUser.Password), []byte(user.Password)) == nil {
            token, err := generateToken(testUser)
            if err != nil {
                return "", false
            }
            return token, true
        }
    }

    return "", false
}

func VerifyToken(tokenString string) (*jwt.Token, bool) {
    token, err := jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error) {
        return secret.verifyKey, nil
    })
    if err == nil && token.Valid {
        return token, true
    }

    return nil, false
}

func generateToken(user *models.User) (string, error) {
    config := settings.GetInstance()
    token := jwt.New(jwt.SigningMethodRS256)
    iat := time.Now()
    exp := iat.Add(time.Duration(config.Server.TokenExpiration) * time.Minute).Unix()
    claims := make(jwt.MapClaims)
    claims["exp"] = exp
    claims["iat"] = iat.Unix()
    claims["user_id"] = user.Id
    token.Claims = claims
    tokenString, err := token.SignedString(secret.secretKey)

    if err != nil {
        panic(err)
        return "", err
    }

    return tokenString, nil
}
