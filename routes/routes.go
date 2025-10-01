package routes

import (
    "github.com/gin-gonic/gin"
    "user-service/controllers"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.POST("/signup", controllers.Signup)
    r.POST("/login", controllers.Login)

    auth := r.Group("/dashboard")
    auth.Use(controllers.AuthMiddleware())
    auth.GET("/", controllers.Dashboard)

    return r
}

