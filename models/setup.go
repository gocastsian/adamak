package models
// source: https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "adamak_user:adamak_pass@tcp(127.0.0.1:3306)/adamak?charset=utf8mb4&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	if aErr := database.AutoMigrate(&User{}); aErr != nil {
		panic("Failed to auto migrate database!")
	}

	DB = database
}
