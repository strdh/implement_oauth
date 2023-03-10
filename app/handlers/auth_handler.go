package handlers

import (
    // "fmt"
    "os"
    "net/http"
    "encoding/base64"
    "github.com/julienschmidt/httprouter"
    "github.com/go-playground/validator/v10"
    "exercise/gooauth/app/web/requests"
    "exercise/gooauth/app/web/responses"
    "exercise/gooauth/app/models"
    "exercise/gooauth/utils"
    "exercise/gooauth/app/exception"
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
    userCreateRequest := requests.UserCreateRequest{}
    utils.ReadFromRequestBody(request, &userCreateRequest)

    err := handler.Validate.Struct(userCreateRequest)
    utils.PanicIfError(err)
    
    user := models.User{
        FullName: userCreateRequest.FullName,
        Email: userCreateRequest.Email,
        Username: userCreateRequest.Username,
        Password: utils.HashPassword(userCreateRequest.Password),
    }

    userKey := utils.GenerateKeyToken()

    userResponse := handler.UserModel.Create(request.Context(), user, userKey)
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
    if err != nil {
        panic(exception.NewNotFoundError(err.Error()))
    }

    result, err := utils.VerifyPassword(userResponse.Password, userLoginRequest.Password)
    if err != nil || !result {
        panic(exception.NewNotFoundError(err.Error()))
    }

    userKey, err := handler.UserModel.FindUserKey(request.Context(), userResponse.Id)
    if err != nil {
        panic(exception.NewNotFoundError(err.Error()))
    }

    AESKEY := []byte(os.Getenv("AES_KEY"))
    decodedKey, err := base64.StdEncoding.DecodeString(userKey)
    finalKey, err := utils.AesDecrypt(decodedKey, AESKEY)
    utils.PanicIfError(err)

    dataResponse := map[string]interface{}{
        "username": userResponse.Username,
        "token": utils.GenerateToken(userResponse.Id, userResponse.Username, base64.StdEncoding.EncodeToString(finalKey)),
    }

    webResponse := responses.WebResponse{
        Code: http.StatusOK,
        Status: "User logged in successfully",
        Data: dataResponse,
    }

    utils.WriteToResponseBody(writer, webResponse)
}