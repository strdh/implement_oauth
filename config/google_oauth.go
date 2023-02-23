package config

import (
    "os"
    "golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleOauthConfig struct {
    Config *oauth2.Config
}

func GoogleOauthInit() *GoogleOauthConfig {
    return &GoogleOauthConfig{
        Config: &oauth2.Config{
            RedirectURL: "http://localhost:8080/callback",
            ClientID: os.Getenv("GOOGLE_CLIENT_ID"),
            ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
            Scopes: []string{"https://www.googleapis.com/auth/userinfo.email"},
            Endpoint: google.Endpoint,
        },
    }
}
