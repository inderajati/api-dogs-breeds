package controllers

import (
	"app/controllers/app/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ResponseMessage struct {
	Message []string `json:"message"`
}

// Handle Get List Sub Breeds
// @Summary Get all sub breeds
// @Description get all sub breeds
// @ID get-list-sub-breeds
// @Accept  json
// @Produce  json
// @Success 200 {object} ResponseMessage
// @Router /breed/{sub-breed}/list [get]
func GetListSubBreeds(c echo.Context) (err error) {
	var bodyBytes []byte
	subBreedParam := c.Param("sub-breed")
	// request body
	if c.Request().Body != nil {
		// Read the Body content
		bodyBytes, _ = io.ReadAll(c.Request().Body)
	}
	// Restore the io.ReadCloser to its original state
	c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	// Continue to use the Body, like Binding it to a struct:
	u := new(ResponseMessage)
	// bind the model with the context body
	er := c.Bind(u)
	if er != nil {
		return er
	}
	// crawl Dog Breeds API
	r, err := utils.GetListSubBreeds(subBreedParam)
	if err != nil {
		return err
	}

	fmt.Println("Get sub breeds: ", *r)

	if err = json.Unmarshal([]byte(*r), &u); err != nil {
		return err
	}

	// Handle case of get list sub-breeds of sheepdog and terrier breed
	if subBreedParam == "sheepdog" || subBreedParam == "terrier" {
		subBreedMap := make(map[string][]string)
		var subBreedImages = make(chan []string)

		// Loop data sub-breeds
		for i := 0; i < len(u.Message); i++ {
			if subBreedParam == "sheepdog" {
				u.Message[i] = subBreedParam + "-" + u.Message[i]
				subBreedMap[u.Message[i]] = []string{}
			} else {
				p := new(ResponseMessage)
				er := c.Bind(p)
				if er != nil {
					return er
				}
				// Call API get list sub-breed images
				img, err := utils.GetListSubBreedImages(subBreedParam, u.Message[i])
				if err != nil {
					return err
				}
				// Unmarshal json result images
				if err = json.Unmarshal([]byte(*img), &p); err != nil {
					return err
				}
				// Channeling data images
				go func(images []string) {
					var data = images
					subBreedImages <- data
				}(p.Message)
				// Append data sub-breed images
				u.Message[i] = subBreedParam + "-" + u.Message[i]
				subBreedMap[u.Message[i]] = <-subBreedImages
			}
		}

		fmt.Println("Map: ", subBreedMap)
		return c.JSON(http.StatusOK, subBreedMap)
	}

	fmt.Println("Response: ", u)
	return c.JSON(http.StatusOK, u)
}

func GetListSubBreedImages(c echo.Context) (err error) {

	var bodyBytes []byte
	breedParam := c.Param("breed")
	subBreedParam := c.Param("sub-breed")
	// request body
	if c.Request().Body != nil {
		// Read the Body content
		bodyBytes, _ = io.ReadAll(c.Request().Body)
	}
	// Restore the io.ReadCloser to its original state
	c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	// Continue to use the Body, like Binding it to a struct:
	u := new(ResponseMessage)
	// bind the model with the context body
	er := c.Bind(u)
	if er != nil {
		return er
	}
	// crawl Dog Breeds API
	r, err := utils.GetListSubBreedImages(breedParam, subBreedParam)
	if err != nil {
		return err
	}

	fmt.Println("Get sub breeds: ", *r)

	if err = json.Unmarshal([]byte(*r), &u); err != nil {
		return err
	}

	fmt.Println("Response: ", u)
	return c.JSON(http.StatusOK, u)

}
