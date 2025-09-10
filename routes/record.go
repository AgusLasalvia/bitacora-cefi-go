package routes

import (
	"bitacora/controllers"

	"github.com/gin-gonic/gin"
)

func RecordRouter(router *gin.Engine) {
	recordController := controllers.NewRecordController()
	route := router.Group("/records")
	{
		route.GET("/record", recordController.GetRecordByID)
		route.GET("/mahine", recordController.GetRecordByMachine)
		route.POST("/add", recordController.AddRecord)
	}
}
