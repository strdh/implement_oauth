package tests

import (
    "testing"
    "time"
    "github.com/golang-jwt/jwt"
    "github.com/stretchr/testify/assert"
    "exercise/gooauth/utils"
)

const KEY = "what-zit-tooya"

type Claims struct {
    Id int `json:"id"`
    Username string `json:"username"`
    jwt.StandardClaims
}

func TestGenerateToken(t *testing.T) {
    id := 1
    username := "testuser"
    tokenString := utils.GenerateToken(id, username)

    // Assert that the token string is not empty
    assert.NotEmpty(t, tokenString, "Token string should not be empty")

    // Verify the token
     token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(KEY), nil
    })

    assert.Nil(t, err, "Error should be nil")
    assert.True(t, token.Valid, "Token should be valid")

    // Assert that the claims are correct
    claims, ok := token.Claims.(*Claims)
    assert.True(t, ok, "Claims should be of type Claims")
    assert.Equal(t, id, claims.Id, "Id in claims should match input id")
    assert.Equal(t, username, claims.Username, "Username in claims should match input username")
    assert.True(t, claims.ExpiresAt == time.Now().Add(time.Hour*24).Unix(), "ExpiresAt in claims should be in the future")
    assert.Equal(t, "what-zit-tooya", claims.Issuer, "Issuer in claims should match expected value")
}

func TestVerifyJWT(t *testing.T) {
    id := 1
    username := "testuser"

    stringToken := utils.GenerateToken(id, username)
    assert.True(t, utils.VerifyJWT(stringToken), "Token should be valid")
    invalidToken := stringToken + "invalid"
    assert.False(t, utils.VerifyJWT(invalidToken), "Token should be invalid")
}