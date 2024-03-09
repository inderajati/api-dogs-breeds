package controllers

import (
	"app/controllers/fetchDogBreeds/app/config"
	"app/controllers/fetchDogBreeds/app/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handle get all dogs
// @Summary Show dog list
// @Description get dog list
// @ID get-dog-list
// @Accept  json
// @Produce  json
// @Success 200 {object} Dogs
// @Router /dogs [get]
func GetAllDogs(c echo.Context) error {
	db := config.DB()
	var dogs []models.Dogs

	if res := db.Find(&dogs); res.Error != nil {
		data := map[string]interface{}{
			"message": res.Error.Error(),
		}
		return c.JSON(http.StatusOK, data)
	}
	response := map[string]interface{}{
		"data": dogs,
	}
	return c.JSON(http.StatusOK, response)
}

// Handle create dog
// @Summary Create a new dog
// @Description create a new dog
// @ID create-dog
// @Accept  json
// @Produce  json
// @Success 200 {object} Message "Dog created successfully"
// @Router /dog [post]
func CreateDog(c echo.Context) error {
	d := new(models.Dogs)
	db := config.DB()

	// Binding data
	if err := c.Bind(d); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	dog := &models.Dogs{
		Name:     d.Name,
		Breed:    d.Breed,
		SubBreed: d.SubBreed,
	}

	if err := db.Create(&dog).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Dog created successfully",
	}

	return c.JSON(http.StatusOK, response)
}

// Handle update dog
// @Summary Update dog
// @Description update dog by ID
// @ID update-dog-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Dog ID"
// @Success 200 {object} Message "Dog updated successfully"
// @Router /dogs/{id} [put]
func UpdateDog(c echo.Context) error {
	id := c.Param("id")
	d := new(models.Dogs)
	db := config.DB()

	// Binding data
	if err := c.Bind(d); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	existingDog := new(models.Dogs)

	if err := db.First(&existingDog, id).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusNotFound, data)
	}

	existingDog.Name = d.Name
	existingDog.Breed = d.Breed
	existingDog.SubBreed = d.SubBreed
	if err := db.Save(&existingDog).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Dog updated successfully",
	}

	return c.JSON(http.StatusOK, response)
}

// Handle delete dog
// @Summary Delete a dog
// @Description delete a dog
// @ID delete-dog
// @Accept  json
// @Produce  json
// @Param id path int true "Dog ID"
// @Success 200 {object} Message "a dog has been deleted"
// @Router /dogs/{id} [delete]
func DeleteDog(c echo.Context) error {
	id := c.Param("id")
	db := config.DB()

	dog := new(models.Dogs)

	err := db.Delete(&dog, id).Error
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "a dog has been deleted",
	}
	return c.JSON(http.StatusOK, response)
}
