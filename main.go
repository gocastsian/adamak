package main

import (
	"github.com/gocastsian/adamak/adapter/store"
	v1 "github.com/gocastsian/adamak/delivery/http/v1"
	"github.com/gocastsian/adamak/validator"

	"github.com/labstack/echo/v4"
)

func main() {
	dsn := "adamak_user:adamak_pass@tcp(127.0.0.1:3306)/adamak?charset=utf8mb4&parseTime=True&loc=Local"
	// connect to database and auto migrate
	mysqlStore := store.New(dsn)

	// setup http server and router
	e := echo.New()

	// add routes
	e.GET("/users", v1.FindUsers(mysqlStore))
	e.POST("/users", v1.CreateUser(mysqlStore,
		validator.ValidateCreateUser))
	e.GET("/users/:id", v1.FindUser(mysqlStore,
		validator.ValidateFindUser(mysqlStore)))
	e.PATCH("/users/:id", v1.UpdateUser(mysqlStore,
		validator.ValidateUpdateUser(mysqlStore)))
	e.DELETE("/users/:id", v1.DeleteUser(mysqlStore,
		validator.ValidateDeleteUser(mysqlStore)))

	e.Logger.Fatal(e.Start(":8080"))
}
