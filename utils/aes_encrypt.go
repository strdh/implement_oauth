package utils

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
)

func aesEncrypt(plaintext []byte, key []byte) []byte {
    block, err := aes.NewCipher(key)
    PanicIfError(err)

    aesGCM, err := cipher.NewGCM(block)
    PanicIfError(err)

    nonce := make([]byte, aesGCM.NonceSize())
    _, err = rand.Read(nonce)
    PanicIfError(err)

    ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
    return ciphertext
}

func aesDecrypt(ciphertext []byte, key []byte) []byte {
    block, err := aes.NewCipher(key)
    PanicIfError(err)

    aesGCM, err := cipher.NewGCM(block)
    PanicIfError(err)

    nonceSize := aesGCM.NonceSize()
    nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

    plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
    PanicIfError(err)

    return plaintext
}