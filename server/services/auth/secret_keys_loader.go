package auth

import (
    "sync"
    "io/ioutil"
    "crypto/rsa"
    "encoding/pem"
    "crypto/x509"
    "github.com/ipastushenko/simple-chat/settings"
)

var (
    secret *jwtTokenSecret
    once sync.Once
)

func LoadSecretKeys() *jwtTokenSecret {
    once.Do(func() {
        config := settings.GetInstance()
        secretKey, err := ioutil.ReadFile(config.Server.SecretKeyPath)
        if err != nil {
            panic(err)
        }
        verifyKey, err := ioutil.ReadFile(config.Server.VerifyKeyPath)
        if err != nil {
            panic(err)
        }
        secretData, _ := pem.Decode(secretKey)
        verifyData, _ := pem.Decode(verifyKey)
        importedSecretKey, err := x509.ParsePKCS1PrivateKey(secretData.Bytes)
        if err != nil {
            panic(err)
        }
        rawVerifyKey, err := x509.ParsePKIXPublicKey(verifyData.Bytes)
        if err != nil {
            panic(err)
        }
        importedVerifyKey, ok := rawVerifyKey.(*rsa.PublicKey)
        if !ok {
            panic("rsa public key cast")
        }
        secret = &jwtTokenSecret{importedSecretKey, importedVerifyKey}
    })

    return secret
}
