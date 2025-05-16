package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func getWeatherInfo() {
	resp, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=-9.6658&longitude=-35.7353&hourly=temperature_2m,precipitation_probability,relative_humidity_2m&timezone=auto")
	if err != nil {
		log.Fatalf("Failed to fetch the api data: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read the response body: %v", err)
	}

	var respBodyJSON map[string]any

	err = json.Unmarshal(respBody, &respBodyJSON)
	if err != nil {
		log.Fatalf("Failed unmarshaling the response body: %v", err)
	}
	
	prettyJSON, err := json.MarshalIndent(respBodyJSON, "", " ")
	if err != nil {
		log.Fatalf("Failed marshaling with indetation: %v", err)
	}
	
}