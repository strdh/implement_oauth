package exception

import (
    "net/http"
    "github.com/go-playground/validator/v10"
    "exercise/gooauth/utils"
    "exercise/gooauth/app/web/responses"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
    if notFoundError(writer, request, err) {
        return
    }

    if validationErrors(writer, request, err) {
        return
    }

    internalServerError(writer, request, err)
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
    exception, ok := err.(validator.ValidationErrors)

    if ok {
        writer.Header().Set("Content-Type", "application/json")
        writer.WriteHeader(http.StatusBadRequest)

        webResponse := responses.WebResponse{
            Code: http.StatusBadRequest,
            Status: "Validation Error",
            Data: exception.Error()
        }

        utils.WriteToResponseBody(writer, webResponse)
        return true
    } else {
        return false
    }
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
    exception, ok := err.(NotFoundError)

    if ok {
        writer.Header().Set("Content-Type", "application/json")
        writer.WriteHeader(http.StatusNotFound)

        webResponse := responses.WebResponse{
            Code: http.StatusNotFound,
            Status: "Not Found",
            Data: exception.Error,
        }

        utils.WriteToResponseBody(writer, webResponse)
        return true
    } else {
        return false
    }
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
    writer.Header().Set("Content-Type", "application/json")
    writer.WriteHeader(http.StatusInternalServerError)

    webResponse := responses.WebResponse{
        Code: http.StatusInternalServerError,
        Status: "Internal Server Error",
        Data: err,
    }

    helper.WriteToResponseBody(writer, webResponse)
}