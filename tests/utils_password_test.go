package tests

import (
    "testing"
    "exercise/gooauth/utils"
    "encoding/base64"
    "github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
    password := "Secret-123-**&&^^"
    hashedPassword := utils.HashPassword(password)

    assert.NotEqual(t, password, hashedPassword, "Hashed password should not be equal to the original password")

    _, err := base64.StdEncoding.DecodeString(hashedPassword)
    assert.Nil(t, err, "Hashed password should be a valid base64 encoded string")
}

func TestVerifyPassword(t *testing.T) {
    password := "Secret-123-**&&^^"
    hashedPassword := utils.HashPassword(password)

    assert.True(t, utils.VerifyPassword(hashedPassword, password), "Password should be verified")
    assert.False(t, utils.VerifyPassword(hashedPassword, "wrongpassword"), "Password should not be verified")
}
