package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"encoding/json"
	"time"
)

func makeRequest(url string) []map[string]interface{}{
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()

	
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	
	var responseObject []interface{}
	if err := json.Unmarshal(responseData, &responseObject); err != nil {
		log.Fatal("Error unmarshalling JSON:", err)
	}
	if len(responseObject) == 0 {
		log.Fatal("No data returned from the API")
	}
	data := responseObject[0].(map[string]interface{})
	return []map[string]interface{}{data}
}

func main(){
	url := "https://api.wheretheiss.at/v1/satellites/25544/positions?timestamps="
	var dateTime string
	// dateTime = fmt.Sprint(time.Now().Format("2006-01-02 15:04:05"))
	var date string
	var tme string
	fmt.Print("Enter Date and Time in format, YYYY-MM-DD HR:MIN:SEC -> ")
	fmt.Scan(&date, &tme)
	dateTime = date + " " + tme
	layout := "2006-01-02 15:04:05"
	location, _ := time.LoadLocation("Asia/Kolkata")
	t, err := time.ParseInLocation(layout, dateTime, location)
	for err != nil {
		fmt.Println("Error parsing date and time:", err)
		fmt.Println("Try again")
		fmt.Print("Enter Date and Time in format, YYYY-MM-DD HR:MIN:SEC -> ")
		fmt.Scan(&date, &tme)
		dateTime = date + " " + tme
		t, err = time.ParseInLocation(layout, dateTime, location)
	}
	timestamp := t.Unix()
	// fmt.Printf("Timestamp: %d\n", timestamp)
	url = url + fmt.Sprint(timestamp)
	data := makeRequest(url)[0]
	latitude, latOk := data["latitude"]
	longitude, lonOk := data["longitude"]
	if latOk && lonOk {
		fmt.Printf("Latitude: %f, Longitude: %f\n", latitude, longitude)
	}
}