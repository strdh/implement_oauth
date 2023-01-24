package utils

import (
    "time"
    "github.com/golang-jwt/jwt"
)

const KEY = "what-zit-tooya"

type Claims struct {
    Id int `json:"id"`
    Username string `json:"username"`
    jwt.StandardClaims
}

func GenerateToken(id int, username string) string {
    claims := Claims{
        Id: id,
        Username: username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
            Issuer: "what-zit-tooya",
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    result, err := token.SignedString([]byte(KEY))
    PanicIfError(err)
    return result
}

func VerifyJWT(jwtToken string) bool {
    token, err := jwt.ParseWithClaims(jwtToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(KEY), nil
    })

    if err != nil || !token.Valid {
        return false
    }

    return true
}