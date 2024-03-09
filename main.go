package main

import (
	"app/controllers/app/config"
	"app/controllers/app/controllers"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "app/controllers/docs"

	"github.com/labstack/echo/v4"
)

// @title API for Dog Breeds
// @version 1.0
// @description This is a simple API service for Dog Breeds consume from the Dog API.
// @host localhost:3002
// @BasePath /
func main() {
	
	// Connect To Database
	config.DatabaseInit()
	gorm := config.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	dbGorm.Ping()
	e := echo.New()

	// get all breeds endpoint
	e.GET("/breeds/list/all", controllers.GetListBreeds)
	e.GET("/breed/:sub-breed/list", controllers.GetListSubBreeds)
	e.GET("/breed/:breed/images", controllers.GetListBreedImages)

	// endpoint for dog services
	e.GET("/dogs", controllers.GetAllDogs)
	e.POST("/dog", controllers.CreateDog)
	e.PUT("/dogs/:id", controllers.UpdateDog)
	e.DELETE("/dogs/:id", controllers.DeleteDog)

	// swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// Start Server
	e.Logger.Fatal(e.Start(":3002"))
}