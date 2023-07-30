package main

import (
	"log"
	"todo/configurations"
	"todo/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	var migrationError error = configurations.RunMigrations()

	if migrationError != nil {
		log.Fatal(migrationError)
	}
	
	var e *echo.Echo = echo.New()
	routes.Routes(e)
	e.Logger.Fatal(e.Start(":8080"))

}