package main

import "github.com/gocastsian/adamak/models"

func main() {

	// Connect to database and auto migrate
	models.ConnectDatabase()
}
