package middleware

import (
    "os"
    "strconv"
    "encoding/base64"
    "net/http"
    "github.com/julienschmidt/httprouter"
    "exercise/gooauth/app/web/responses"
    "exercise/gooauth/app/models"
    "exercise/gooauth/utils"
    "exercise/gooauth/app/exception"
)

func AuthMiddleware(next httprouter.Handle, userModel *models.UserModel) httprouter.Handle {
    return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
        userToken := request.Header.Get("JWT-TOKEN")
        userId := request.Header.Get("USER-ID")
        finalUserId, err := strconv.Atoi(userId)
        utils.PanicIfError(err)

        userKey, err := userModel.FindUserKey(request.Context(), finalUserId)
        if err != nil {
            panic(exception.NewNotFoundError(err.Error()))
        }

        AESKEY := []byte(os.Getenv("AES_KEY"))
        decodedKey, err := base64.StdEncoding.DecodeString(userKey)
        finalKey, err := utils.AesDecrypt(decodedKey, AESKEY)
        utils.PanicIfError(err)

        result := utils.VerifyJWT(userToken, base64.StdEncoding.EncodeToString(finalKey))

        if !result || userToken == "" {
            writer.Header().Set("Content-Type", "application/json")
            writer.WriteHeader(http.StatusUnauthorized)

            response := responses.WebResponse{
                Code: http.StatusUnauthorized,
                Status: "Unauthorized",
                Data: "You dont have permission to access this resource",
            }

            utils.WriteToResponseBody(writer, response)
            return
        }

        next(writer, request, params)
    }
}

