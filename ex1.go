//go:build ex01.go
//With Concurrency
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const APIkey = "022d82153e5544a79dec4c79c484bf87"

func fetchweather(city string,ch chan<- string, wg *sync.WaitGroup){
	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		}`json:"main"`
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s",city,APIkey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error Fetching weather for %s: %s\n",city,err)
		defer wg.Done()
	}

	defer resp.Body.Close()

	if err:=json.NewDecoder(resp.Body).Decode(&data); err!=nil {
		fmt.Printf("Error Decoding weather data for %s: %s",city,err)
		defer wg.Done()
	}
	defer wg.Done()

	ch <- fmt.Sprintf("Weather for %s: %.2f°C",city,data.Main.Temp)
}

func main(){
	begintime := time.Now()
	ch := make(chan string)
	var wg sync.WaitGroup
	cities := []string {"Surat", "Kolkata", "Kharagpur", "Mumbai"}

	for _,city :=range cities{
		wg.Add(1)
		go fetchweather(city, ch, &wg)
	}
	go func(){
		wg.Wait()
		close(ch)
	}()

	for data := range ch{
		fmt.Println(data)
	}
	endtime :=time.Now()

	fmt.Printf("Execution time: %s\n",endtime.Sub(begintime))
}