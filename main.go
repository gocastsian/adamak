package main

import (
	"github.com/gocastsian/adamak/protocol"
	"github.com/gocastsian/adamak/services"

	"github.com/labstack/echo/v4"
)

func main() {
	var config = protocol.ServiceInitializeConfig{
		StorageEngine: "mysql", // TODO::: get from OS flags
	}

	switch config.StorageEngine {
	case "mysql":
		dsn := "adamak_user:adamak_pass@tcp(127.0.0.1:3306)/adamak?charset=utf8mb4&parseTime=True&loc=Local"
		config.MySQL.Init(dsn)
	default:
		panic("[ADAMAK] Don't support requested storage engine: " + config.StorageEngine)
	}

	// setup http server and router
	config.Echo = echo.New()

	services.Init(&config)

	config.Echo.Logger.Fatal(config.Echo.Start(":8080"))
}
