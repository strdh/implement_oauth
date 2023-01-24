package handlers

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "github.com/go-playground/validator/v10"
    "exercise/gooauth/app/web/requests"
    "exercise/gooauth/app/web/responses"
    "exercise/gooauth/app/models"
    "exercise/gooauth/utils"
)

type AuthHandler struct {
    UserModel *models.UserModel
    Validate *validator.Validate
}

func NewAuthHandler(userModel *models.UserModel, validator *validator.Validate) *AuthHandler {
    return &AuthHandler{
        UserModel: userModel,
        Validate: validator,
    }
}

func (handler *AuthHandler) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
    userCrateRequest := requests.UserCreateRequest{}
    utils.ReadFromRequestBody(request, &userCrateRequest)

    err := handler.Validate.Struct(userCrateRequest)
    utils.PanicIfError(err)
    user := models.User{
        FullName: userCrateRequest.FullName,
        Email: userCrateRequest.Email,
        Username: userCrateRequest.Username,
        Password: utils.HashPassword(userCrateRequest.Password),
    }

    userResponse := handler.UserModel.Create(request.Context(), user)
    webResponse := responses.WebResponse{
        Code: http.StatusOK,
        Status: "User created successfully",
        Data: userResponse,
    }

    utils.WriteToResponseBody(writer, webResponse)
}

func (handler *AuthHandler) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
    userLoginRequest := requests.UserLoginRequest{}
    utils.ReadFromRequestBody(request, &userLoginRequest)

    err := handler.Validate.Struct(userLoginRequest)
    utils.PanicIfError(err)

    userResponse, err := handler.UserModel.FindByUsername(request.Context(), userLoginRequest.Username)
    utils.PanicIfError(err)

    if err != nil || !result {
        panic(exception.NewNotFoundError(err.Error()))
    }

    if !utils.VerifyPassword(userLoginRequest.Password, userResponse.Password) {
        panic(exception.NewNotFoundError(err.Error()))
    }

    dataResponse := map[string]interface{}{
        "username": userResponse.Username,
        "token": utils.GenerateToken(userResponse.Username),
    }

    webResponse := responses.WebResponse{
        Code: http.StatusOK,
        Status: "User logged in successfully",
        Data: dataResponse,
    }

    utils.WriteToResponseBody(writer, webResponse)
}