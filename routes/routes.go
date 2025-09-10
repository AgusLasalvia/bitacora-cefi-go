package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// allow cors && specific methods
	r.Use(cors.New(
		cors.Config{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"POST", "GET"},
			AllowHeaders: []string{"Content-Type"},
		},
	))

	// inport && init routes
	RecordRouter(r)

	return r
}
