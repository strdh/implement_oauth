package handlers

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "exercise/gooauth/app/web/responses"
    "exercise/gooauth/utils"
)

type MainHandler struct {}

func NewMainHandler() *MainHandler {
    return &MainHandler{}
}

func (handler *MainHandler) Index(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
    webResponse := responses.WebResponse{
        Code: http.StatusOK,
        Status: "Welcome to Gooauth",
        Data: nil,
    }

    utils.WriteToResponseBody(writer, webResponse)
}