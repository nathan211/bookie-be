package routes

import (
	"bookie-be/handlers"

	"github.com/gin-gonic/gin"
)

func BusinessTypeRoutes(r *gin.Engine) {
	authorized := r.Group("/")
	authorized.Use(handlers.Authenticate())

	{
		// BusinessType routes
		authorized.POST("/business_types", handlers.CreateBusinessType)
		authorized.GET("/business_types", handlers.GetBusinessTypes)
		authorized.PUT("/business_types/:id", handlers.UpdateBusinessType)
		authorized.DELETE("/business_types/:id", handlers.DeleteBusinessType)
	}
}
