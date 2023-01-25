package main 

import (
    "os"
    "net/http"
    // "github.com/joho/godotenv"
    "exercise/gooauth/routes"
    "exercise/gooauth/utils"
)

func main() {
    router := routes.NewRouter()
    server := http.Server{
        Addr: os.Getenv("ADDRESS"),
        Handler: router,
    }

    err := server.ListenAndServe()
    utils.PanicIfError(err)
}