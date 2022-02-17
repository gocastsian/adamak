package external

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL struct {
	DB *gorm.DB
}

func (m *MySQL) Init(dsn string) {
	var database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("[External Dependency] Failed to connect to mysql database!")
	}
	m.DB = database
}
