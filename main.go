package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const weatherAPIURL = "https://api.openweathermap.org/data/2.5/weather?q=Toronto&appid=f59d1d1e04e09cfba52671eb99b77335"

type WeatherInfo struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		Humidity  int     `json:"humidity"`
		WindSpeed float64 `json:"wind_speed"`
	} `json:"main"`
}

func getTorontoWeather() (*WeatherInfo, error) {
	resp, err := http.Get(weatherAPIURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weatherInfo WeatherInfo
	err = json.Unmarshal(body, &weatherInfo)
	if err != nil {
		return nil, err
	}

	return &weatherInfo, nil
}

func torontoWeatherHandler(w http.ResponseWriter, r *http.Request) {
	weatherInfo, err := getTorontoWeather()
	if err != nil {
		http.Error(w, "Error fetching weather data", http.StatusInternalServerError)
		return
	}

	temperature := weatherInfo.Main.Temp
	feelsLike := weatherInfo.Main.FeelsLike
	humidity := weatherInfo.Main.Humidity
	windSpeed := weatherInfo.Main.WindSpeed
	description := weatherInfo.Weather[0].Description

	fmt.Fprintf(w, "Weather in Toronto:\n")
	fmt.Fprintf(w, "Temperature: %.2f°C\n", temperature)
	fmt.Fprintf(w, "Feels like: %.2f°C\n", feelsLike)
	fmt.Fprintf(w, "Humidity: %d%%\n", humidity)
	fmt.Fprintf(w, "Wind Speed: %.2f m/s\n", windSpeed)
	fmt.Fprintf(w, "Description: %s\n", description)

	fmt.Fprintf(w, "Temperature in Toronto: %.2f°C", temperature)
}

func main() {
	http.HandleFunc("/weather", torontoWeatherHandler)
	fmt.Println("Server is listening on port 8088")
	log.Fatal(http.ListenAndServe(":8088", nil))
}
