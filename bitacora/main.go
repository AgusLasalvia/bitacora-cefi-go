package main

import (
	"bitacora/routes"

	_ "bitacora/config"
)

func main() {

	r := routes.SetupRoutes()

	r.Run("127.0.0.1:8080")
}
