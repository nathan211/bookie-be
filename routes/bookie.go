package routes

import (
    "bookie-be/handlers"
    "github.com/gin-gonic/gin"
)

func BookieRoutes(r *gin.Engine) {
    authorized := r.Group("/")
    authorized.Use(handlers.Authenticate())
    {
        authorized.POST("/bookies", handlers.CreateBookie)
        authorized.GET("/bookies", handlers.GetBookies)
        authorized.PUT("/bookies", handlers.UpdateBookie)
        authorized.DELETE("/bookies", handlers.DeleteBookie)
    }
}