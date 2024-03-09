package controllers

import (
	"app/controllers/app/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type DogBreeds struct {
	Message map[string]interface{} `json:"message"`
}

type DogBreadImages struct {
	Message []string `json:"message"`
}

// Handle Get List Breeds
// @Summary Get all breeds
// @Description get all breeds
// @ID get-list-breeds
// @Accept  json
// @Produce  json
// @Success 200 {object} DogBreeds
// @Router /breeds/list/all [get]
func GetListBreeds(c echo.Context) (err error) {
	var bodyBytes []byte
	// request body
	if c.Request().Body != nil {
		// Read the Body content
		bodyBytes, _ = io.ReadAll(c.Request().Body)
	}
	// Restore the io.ReadCloser to its original state
	c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	// Continue to use the Body, like Binding it to a struct:
	u := new(DogBreeds)
	// bind the model with the context body
	er := c.Bind(u)
	// panic!
	if er != nil {
		return er
	}
	// crawl Dog Breeds API
	r, err := utils.GetListBreeds()
	if err != nil {
		return err
	}
	// Unmarshal json
	if err = json.Unmarshal([]byte(*r), &u); err != nil {
		return err
	}

	fmt.Println("Response: ", u)
	return c.JSON(http.StatusOK, u)
}

// Handle Get List Breed Images
// @Summary Get all breed images
// @Description get all breed images
// @ID get-list-breed-images
// @Accept  json
// @Produce  json
// @Success 200 {object} DogBreadImages
// @Router /breed/{breed}/images [get]
func GetListBreedImages(c echo.Context) (err error) {
	var bodyBytes []byte
	breedParam := c.Param("breed")

	// request body
	if c.Request().Body != nil {
		// Read the Body content
		bodyBytes, _ = io.ReadAll(c.Request().Body)
	}
	// Restore the io.ReadCloser to its original state
	c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	// Continue to use the Body, like Binding it to a struct:
	u := new(DogBreadImages)
	// bind the model with the context body
	er := c.Bind(u)
	// panic!
	if er != nil {
		return er
	}
	// crawl Dog Breeds API
	r, err := utils.GetListBreedImages(breedParam)
	if err != nil {
		return err
	}
	// Unmarshal json
	if err = json.Unmarshal([]byte(*r), &u); err != nil {
		return err
	}

	// Handle case of get list images of shiba breed
	if breedParam == "shiba" {
		var shibaImages []string
		for i := 0; i < len(u.Message); i++ {
			// Retrieve image number	
			img := strings.Split(u.Message[i], "/")
			title := img[5]
			// Check image string number
			numArray := regexp.MustCompile(`\d`).FindAllString(title, -1)
			// Convert string number to int
			num := strings.Join(numArray, "")
			numInt, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println("Error: ", err)
			}
			// Check if image number is odd
			if numInt % 2 != 0 {
				shibaImages = append(shibaImages, u.Message[i])
			}
		}
		u.Message = shibaImages
	}

	fmt.Println("Response: ", u)
	return c.JSON(http.StatusOK, u)
}