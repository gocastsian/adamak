package main

import (
	"github.com/gocastsian/adamak/controllers"
	"github.com/gocastsian/adamak/models"
	"github.com/gocastsian/adamak/validations"
	"github.com/labstack/echo/v4"
)

func main() {

	// connect to database and auto migrate
	models.ConnectDatabase()

	// setup http server and router
	e := echo.New()

	// add custom validator
	e.Validator = validations.New()

	// add routes
	e.POST("/users", controllers.CreateUser)

	e.Logger.Fatal(e.Start(":8080"))
}
