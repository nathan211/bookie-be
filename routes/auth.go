package routes

import (
    "bookie-be/handlers"
    "github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
    r.POST("/register", handlers.Register)
    r.POST("/login", handlers.Login)
}