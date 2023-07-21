package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CatImage struct {
	ID       string `json:"id"`
	URL      string `json:"url"`
	Category string `json:"category"`
}

const (
	baseURL = "https://api.thecatapi.com/v1/images/search"
	apiKey  = "live_2qbvPlCSToNpUrR0CS6dLYhkfPBuO84brSz4qRybTtkX4iWw8jGn0lB0hia2M0wk"
)

func fetchCatImage() {
	url := fmt.Sprintf("%s?limit=10&breed_ids=beng&api_key=%v", baseURL, apiKey)
	response, err := http.Get(url)
	if err != nil {
		fmt.Sprintf("Error fetching cat image:", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Sprintf("Error reading response body:", err)
		return
	}
	var catImages []CatImage
	err = json.Unmarshal(body, &catImages)
	if err != nil {
		fmt.Sprintf("Error reading JSON response:", err)
		return
	}
	// fmt.Println(string(body))
	if len(catImages) > 0 {
		catImage := catImages[0]
		fmt.Println("Cat Image ID:", catImage.ID)
		fmt.Println("Cat Image URL:", catImage.URL)
		fmt.Println("Cat Image Category:", catImage.Category)
	} else {
		fmt.Println("No cat images found.")
	}

}
func main() {
	fetchCatImage()
}
