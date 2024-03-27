//go:build ex0.go
//without Concurrency
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const APIkey = "022d82153e5544a79dec4c79c484bf87"

func fetchweather(city string) interface{} {
	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, APIkey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error Fetching weather for %s: %s\n", city, err)
		return data.Main.Temp
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Printf("Error Decoding weather data for %s: %s", city, err)
		return data.Main.Temp
	}
	return data.Main.Temp
}

func main() {
	begintime := time.Now()

	cities := []string{"Surat", "Kolkata", "Kharagpur", "Mumbai"}

	for _, city := range cities {
		data := fetchweather(city)
		fmt.Printf("Weather for %s is %v\n", city ,data)
	}
	endtime := time.Now()

	fmt.Printf("Execution time: %s\n", endtime.Sub(begintime))
}
