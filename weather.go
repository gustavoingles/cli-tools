package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func getWeatherInfo() {
	resp, err := http.Get("https://api.weatherapi.com/v1/current.json?q=Maceio&key=5d9163f72c2a45749e0202217251605")
	if err != nil {
		log.Fatalf("Failed to fetch the api data: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read the response body: %v", err)
	}

	type Location struct {
		Name string `json:"name"`
		Region string `json:"region"`
		Country string `json:"country"`
		TimezoneId string `json:"tz_id"`
		Localtime string `json:"localtime"`
	}

	type Condition struct {
		Text string `json:"text"`
	}

	type Current struct {
		Temperature float64 `json:"temp_c"`
		Condition Condition `json:"condition"`
	}


	type weatherInfo struct {
		Location Location `json:"location"`
		Current Current `json:"current"`
	}

	weather := weatherInfo{}
	json.Unmarshal(respBody, &weather)

	weatherJSON, err := json.MarshalIndent(weather, "", " ")
	if err != nil {
		log.Fatalf("Failed marshaling into JSON: %v", err)
	}
	
	fmt.Println(string(weatherJSON))
}
