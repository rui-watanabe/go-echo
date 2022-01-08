package database

import (
	"go-echo/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := config.Config.DBUser + ":" + config.Config.DBPass + "@tcp(" + config.Config.DBHost + ")/" + config.Config.DBName + "?charset=utf8&parseTime=True&loc=Local"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	DB = connection
}
