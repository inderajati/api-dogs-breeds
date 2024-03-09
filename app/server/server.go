package server

import (
	"app/controllers/fetchDogBreeds/app/controllers"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// DogBreedsAPI structur
type DogBreedsAPI struct {
	e *echo.Echo
}

// NewServer Instance of Echo
func NewServer() *DogBreedsAPI {
	return &DogBreedsAPI{
		e: echo.New(),
	}
}

// Start server functionality
func (s *DogBreedsAPI) Start(port string) {
	// get all breeds endpoint
	s.e.GET("/breeds/list/all", controllers.GetListBreeds)
	s.e.GET("/breed/:sub-breed/list", controllers.GetListSubBreeds)
	s.e.GET("/breed/:breed/images", controllers.GetListBreedImages)

	// endpoint for dog services
	s.e.GET("/dogs", controllers.GetAllDogs)
	s.e.POST("/dog", controllers.CreateDog)
	s.e.PUT("/dogs/:id", controllers.UpdateDog)
	s.e.DELETE("/dogs/:id", controllers.DeleteDog)

	// swagger
	s.e.GET("/swagger/*", echoSwagger.WrapHandler)
	// Start Server
	s.e.Logger.Fatal(s.e.Start(port))
}

// Close server functionality
func (s *DogBreedsAPI) Close() {
	s.e.Close()
}