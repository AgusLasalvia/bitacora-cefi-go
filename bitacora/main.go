package main

import (
	"bitacora/routes"

	_ "bitacora/config"
)

func main() {

	r := routes.SetupRoutes()

	r.Run("0.0.0.0:8080")
}
