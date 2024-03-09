package main

import (
	"app/controllers/fetchDogBreeds/app/config"
	"app/controllers/fetchDogBreeds/app/server"

	_ "app/server/build/docs"
)

// @title API for Dog Breeds
// @version 1.0
// @description This is a simple API service for Dog Breeds consume from the Dog API.
// @host localhost:3002
func main() {
	// Connect To Database
	config.DatabaseInit()
	gorm := config.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	dbGorm.Ping()

	server := server.NewServer()
	server.Start(":3002")
}