package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"encoding/json"
)

func makeRequest(url string) []map[string]interface{}{
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject map[string]interface{}
	json.Unmarshal([]byte(responseData), &responseObject)
	if err := json.Unmarshal(responseData, &responseObject); err != nil {
		log.Fatal("Error unmarshalling JSON:", err)
	}
	return []map[string]interface{}{responseObject}
}

func main(){
	url := `https://api.wheretheiss.at/v1/satellites/25544/`
	data := makeRequest(url)[0]

	latitude, latOk := data["latitude"]
	longitude, lonOk := data["longitude"]
	if latOk && lonOk {
		fmt.Printf("Latitude: %f, Longitude: %f\n", latitude, longitude)
	}
	url = "https://api.wheretheiss.at/v1/coordinates/"+fmt.Sprint(latitude)+","+fmt.Sprint(longitude)
	data = makeRequest(url)[0]
	timezone_id := data["timezone_id"]
	country_code := data["country_code"]
	fmt.Println("Time Zone: ", timezone_id, "Coutnry: ", country_code)
}