package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	AuthRoutes(r)

	BusinessTypeRoutes(r)
	BusinessRoutes(r)

	return r
}
