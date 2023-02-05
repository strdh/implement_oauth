package main

import (
    // "os"
    "fmt"
    "crypto/rand"
    "crypto/aes"
    "crypto/cipher"
    "encoding/base64"
    // "golang.org/x/crypto/ed25519"
    "exercise/gooauth/utils"
    "github.com/joho/godotenv"
)

func enc(plaintext []byte, key []byte) []byte {
    block, err := aes.NewCipher(key)
    utils.PanicIfError(err)

    aesGCM, err := cipher.NewGCM(block)
    utils.PanicIfError(err)

    nonce := make([]byte, aesGCM.NonceSize())
    _, err = rand.Read(nonce)
    utils.PanicIfError(err)

    ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
    return ciphertext
}

func dec(ciphertext []byte, key []byte) []byte {
    block, err := aes.NewCipher(key)
    utils.PanicIfError(err)

    aesGCM, err := cipher.NewGCM(block)
    utils.PanicIfError(err)

    nonceSize := aesGCM.NonceSize()
    nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

    plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
    utils.PanicIfError(err)

    return plaintext
}


func main() {
    err1 := godotenv.Load(".env")
    utils.PanicIfError(err1)
    bytes := make([]byte, 32)
    _, err := rand.Read(bytes)
    utils.PanicIfError(err)
    fmt.Println("key: ", base64.StdEncoding.EncodeToString(bytes))
    fmt.Println("key: ", bytes)

    plaintext := []byte("Anselma Hanadya Putri")
    ciphertext := enc(plaintext, bytes)
    fmt.Println("plaintext: ", base64.StdEncoding.EncodeToString(plaintext))
    fmt.Println("ciphertext: ", base64.StdEncoding.EncodeToString(ciphertext))
    fmt.Println("ciphertext: ", ciphertext)

    fmt.Println("decrypted: ", base64.StdEncoding.EncodeToString(dec(ciphertext, bytes)))

    fmt.Println("Test", utils.GenerateKeyToken())
}