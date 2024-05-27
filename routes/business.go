package routes

import (
	"bookie-be/handlers"

	"github.com/gin-gonic/gin"
)

func BusinessRoutes(r *gin.Engine) {
	authorized := r.Group("/")
	authorized.Use(handlers.Authenticate())

	{
		// Business routes
		authorized.POST("/businesses", handlers.CreateBusiness)
		authorized.GET("/businesses", handlers.GetBusinesses)
		authorized.PUT("/businesses/:id", handlers.UpdateBusiness)
		authorized.DELETE("/businesses/:id", handlers.DeleteBusiness)
	}
}
