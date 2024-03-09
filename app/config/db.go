package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB
var e error

func DatabaseInit() {
	host := "127.0.0.1"
	username := "root"
	password := "root"
	dbName := "dogs_db"
	port := "8889"

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8&parseTime=true&loc=Local"
	database, e = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if e != nil {
		panic(e)
	}
}

func DB() *gorm.DB {
	return database
}