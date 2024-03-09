package utils

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

var baseUrl = "https://dog.ceo/api"

func GetListBreeds() (*string, error) {
	// Build the request
	req, err := http.NewRequest("GET", baseUrl+"/breeds/list/all", nil)
	if err != nil {
		fmt.Println("Error when getting list: ", err)
		return nil, err
	}

	// Send the request set timeout 5 second
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	// Do the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error when sending request: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error when reading response: ", err)
		return nil, err
	}

	// Print the response
	res := string(data)
	fmt.Println("Result: ", res)

	return &res, nil
}

func GetListSubBreeds(subBreed string) (*string, error) {
	// Build the request
	req, err := http.NewRequest("GET", baseUrl+"/breed/"+subBreed+"/list", nil)
	if err != nil {
		fmt.Println("Error when getting list: ", err)
		return nil, err
	}

	// Send the request set timeout 2 second
	client := &http.Client{
		Timeout: 2 * time.Second,
	}
	// Do the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error when sending request: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error when reading response: ", err)
		return nil, err
	}

	// Print the response
	res := string(data)
	fmt.Println("Result: ", res)

	return &res, nil
}

func GetListBreedImages(breed string) (*string, error) {
	// Build the request
	req, err := http.NewRequest("GET", baseUrl+"/breed/"+breed+"/images", nil)
	if err != nil {
		fmt.Println("Error when getting list: ", err)
		return nil, err
	}

	// Send the request
	client := &http.Client{}
	// Do the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error when sending request: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error when reading response: ", err)
		return nil, err
	}

	// Print the response
	res := string(data)
	fmt.Println("Result: ", res)

	return &res, nil
}

func GetListSubBreedImages(breed string, subBreed string) (*string, error) {
	// Build the request
	req, err := http.NewRequest("GET", baseUrl+"/breed/"+breed+"/"+subBreed+"/images", nil)
	if err != nil {
		fmt.Println("Error when getting list: ", err)
		return nil, err
	}

	// Send the request
	client := &http.Client{}
	// Do the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error when sending request: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error when reading response: ", err)
		return nil, err
	}

	// Print the response
	res := string(data)
	fmt.Println("Result: ", res)

	return &res, nil
}