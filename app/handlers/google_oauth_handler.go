package handlers

import (
    "fmt"
    "net/http"
    "golang.org/x/oauth2"
	// "golang.org/x/oauth2/google"
    "io/ioutil"
    "exercise/gooauth/config"
)

var googleOauthConfig = config.GoogleOauthInit()
var oauthStateString = "random"

func GoogleLogin(writer http.ResponseWriter, request *http.Request) {
    url := googleOauthConfig.Config.AuthCodeURL(oauthStateString)
    http.Redirect(writer, request, url, http.StatusTemporaryRedirect)
}

func GoogleCallback(writer http.ResponseWriter, request *http.Request) {
    content, err :=  getUserInfo(request.FormValue("state"), request.FormValue("code"))

    if err != nil {
        fmt.Println(err.Error())
        http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
        return
    }

    fmt.Fprintf(writer, "Content: %s\n", content)
}

func getUserInfo(state string, code string) ([]byte, error) {
   if state != oauthStateString {
       return nil, fmt.Errorf("invalid oauth state")
   }

    token, err := googleOauthConfig.Config.Exchange(oauth2.NoContext, code)
    if err != nil {
        return nil, fmt.Errorf("Code exchange failed with '%s'", err)
    }

    response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
    if err != nil {
        return nil, fmt.Errorf("Failed getting user info with '%s'", err)
    }

    defer response.Body.Close()
    contents, err := ioutil.ReadAll(response.Body)
    if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	return contents, nil
}