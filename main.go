package main 

import (
    "net/http"
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