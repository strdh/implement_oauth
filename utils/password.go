package utils

import (
    "encoding/base64"
    "golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
    salt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    PanicIfError(err)
    return base64.StdEncoding.EncodeToString(salt)
}

func VerifyPassword(hashedPassword, password string) bool {
    decodeHash, err := base64.StdEncoding.DecodeString(hashedPassword)
    PanicIfError(err)

    err = bcrypt.CompareHashAndPassword(decodeHash, []byte(password))
    if err != nil {
        return false
    }

    return true
}