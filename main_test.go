package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetWeather(w *testing.T) {
	//Create a req
	req, err := http.NewRequest("GET", "/weather", nil)
	if err != nil {
		w.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(torontoWeatherHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		w.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var response WeatherInfo
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	fmt.Println("Status code is", rr.Code)
	_ = response.Main.Temp
}
