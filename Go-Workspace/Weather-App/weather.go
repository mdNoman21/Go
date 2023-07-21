package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	baseURL = "http://api.weatherstack.com"
	apiKey  = "8130f1d58819127465e88b75449042f0"
)

type WeatherResponse struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current struct {
		Temperature          float64  `json:"temperature"`
		Weather_Descriptions []string `json:weather_descriptions`
	} `json:"current"`
}

func fetchWeather(location string) {
	url := fmt.Sprintf("%s/current?access_key=%s&query=%s", baseURL, apiKey, location)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching weather data:", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	// fmt.Println("JSON response:", string(body)) // Debug statement to check the JSON response

	var weatherResponse WeatherResponse
	err = json.Unmarshal(body, &weatherResponse)
	if err != nil {
		fmt.Println("Error parsing JSON response:", err)
		return
	}

	fmt.Println("Location:", weatherResponse.Location.Name)
	fmt.Println("Temperature:", weatherResponse.Current.Temperature)
	fmt.Println("Condition:", weatherResponse.Current.Weather_Descriptions)
}

func main() {
	location := "Muzaffarpur" // Replace with your preferred location
	fetchWeather(location)
}
