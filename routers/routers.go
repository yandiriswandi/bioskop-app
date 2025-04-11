package routers

import (
	"bioskop-app/controllers"

	"github.com/gin-gonic/gin"
)

func StartSever() *gin.Engine {
	router := gin.Default()

	router.POST("/bioskop", controllers.CreateBioskop)
	router.PUT("/bioskop/:id", controllers.UpdateBioskop)
	router.DELETE("/bioskop/:id", controllers.DeleteBioskop)
	router.GET("/bioskop", controllers.GetBioskopList)
	router.GET("/bioskop/:id", controllers.GetBioskopByID)

	return router
}
