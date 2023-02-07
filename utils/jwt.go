package utils

import (
    "os"
    "time"
    "crypto/rand"
    "encoding/base64"
    "github.com/golang-jwt/jwt"
)

type Claims struct {
    Id int `json:"id"`
    Username string `json:"username"`
    jwt.StandardClaims
}

func GenerateKeyToken() string {
    AESKEY := []byte(os.Getenv("AES_KEY"))

    bytes := make([]byte, 32)
    _, err := rand.Read(bytes)
    PanicIfError(err)

    ciphertext := AesEncrypt(bytes, AESKEY)
    return base64.StdEncoding.EncodeToString(ciphertext)
}

func GenerateToken(id int, username string, key string) string {
    claims := Claims{
        Id: id,
        Username: username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
            Issuer: "what-zit-tooya",
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    result, err := token.SignedString([]byte(key))
    PanicIfError(err)
    return result
}

func VerifyJWT(jwtToken string, key string) bool {
    token, err := jwt.ParseWithClaims(jwtToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(key), nil
    })

    if err != nil || !token.Valid {
        return false
    }

    return true
}