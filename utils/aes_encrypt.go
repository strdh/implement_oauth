package utils

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
)

func AesEncrypt(plaintext []byte, key []byte) []byte {
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

func AesDecrypt(ciphertext []byte, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    PanicIfError(err)

    aesGCM, err := cipher.NewGCM(block)
    PanicIfError(err)

    nonceSize := aesGCM.NonceSize()
    nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

    plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return nil, err
    }

    return plaintext, nil
}