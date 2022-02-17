package user

import (
	"github.com/gocastsian/adamak/protocol"
)

func Init(config *protocol.ServiceInitializeConfig) {
	switch config.StorageEngine {
	case "mysql":
		var storage_MySQL storage_MySQL
		storage_MySQL.init(config)
		storage = &storage_MySQL
	default:
		panic("[User Microservice] Don't support requested storage engine: "+config.StorageEngine)
	}

	// add routes
	config.Echo.GET("/users", findUsersEcho)
	config.Echo.POST("/users", createUserEcho)
	config.Echo.GET("/users/:id", findUserEcho)
	config.Echo.PATCH("/users/:id", updateUserEcho)
	config.Echo.DELETE("/users/:id", deleteUserEcho)
}
