package routers

import (
	"bioskop-app/controllers"

	"github.com/gin-gonic/gin"
)

func StartSever() *gin.Engine {
	router := gin.Default()

	router.POST("/bioskop", controllers.CreateBioskop)

	return router
}
