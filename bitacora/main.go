package main

import (
	"bitacora/config"
	"bitacora/routes"

	_ "bitacora/config"
)

func main() {

	defer config.Database.Close()

	r := routes.SetupRoutes()

	r.Run("127.0.0.1:8080")
}
