package routes

import (
    "github.com/julienschmidt/httprouter"
    "github.com/go-playground/validator/v10"
     _ "github.com/go-sql-driver/mysql"
    "exercise/gooauth/app/handlers"
    "exercise/gooauth/app/models"
    "exercise/gooauth/config"
    "exercise/gooauth/app/exception"
)

func NewRouter() *httprouter.Router {
    db := config.NewDB()
    validator := validator.New()

    userModel := models.NewUserModel(db)
    authHandler := handlers.NewAuthHandler(userModel, validator)

    router := httprouter.New()
    router.POST("/api/register", authHandler.Register)
    router.POST("/api/login", authHandler.Login)

    router.PanicHandler = exception.ErrorHandler

    return router
}