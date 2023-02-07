package routes

import (
    "github.com/julienschmidt/httprouter"
    "github.com/go-playground/validator/v10"
     _ "github.com/go-sql-driver/mysql"
    "exercise/gooauth/app/handlers"
    "exercise/gooauth/app/models"
    "exercise/gooauth/config"
    "exercise/gooauth/app/exception"
    "exercise/gooauth/middleware"
)

func NewRouter() *httprouter.Router {
    db := config.NewDB()
    validator := validator.New()

    userModel := models.NewUserModel(db)
    authHandler := handlers.NewAuthHandler(userModel, validator)

    mainHandler := handlers.NewMainHandler()

    router := httprouter.New()
    router.POST("/api/register", authHandler.Register)
    router.POST("/api/login", authHandler.Login)

    router.GET("/api/main", middleware.AuthMiddleware(mainHandler.Index, userModel))

    router.PanicHandler = exception.ErrorHandler

    return router
}