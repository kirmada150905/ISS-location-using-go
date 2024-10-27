package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"encoding/json"
	"time"
	"flag"
)

func makeRequest(url string, option int) []map[string]interface{}{
	response, err := http.Get(url)
	if err != nil{
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)	
	if err != nil{
		log.Fatal(err)
	}


	if option == 1 || option == 2{
		var responseObject map[string]interface{}
		// json.Unmarshal([]byte(responseData), &responseObject)
		if err := json.Unmarshal(responseData, &responseObject); err != nil {
			log.Fatal("Error unmarshalling JSON:", err)
		}
		return []map[string]interface{}{responseObject}
	}else{
		var responseObject []interface{}
		// json.Unmarshal([]byte(responseData), &responseObject)
		if err := json.Unmarshal(responseData, &responseObject); err != nil {
			log.Fatal("Error unmarshalling JSON:", err)
		}
		if len(responseObject) == 0 {
			log.Fatal("No data returned from the API")
		}
		data := responseObject[0].(map[string]interface{})
		return []map[string]interface{}{data}
	}

}


func latAngLong() []float64{
	var res []float64
	url := "https://api.wheretheiss.at/v1/satellites/25544"
	data := makeRequest(url, 1)[0]
	
	latitude, latOk := data["latitude"].(float64)
	longitude, lonOk := data["longitude"].(float64)
	if latOk && lonOk {
		fmt.Printf("Latitude: %f, Longitude: %f\n", latitude, longitude)
	} else {
		fmt.Println("Error extracting latitude and longitude")
	}
	res = append(res, latitude)
	res = append(res, longitude)
	return res
}

func timeAndCountry(){
	coordinates := latAngLong()
	latitude, longitude := coordinates[0], coordinates[1]
	url := "https://api.wheretheiss.at/v1/coordinates/"+fmt.Sprint(latitude)+","+fmt.Sprint(longitude)
	// url := "https://api.wheretheiss.at/v1/coordinates/37.795517,-122.393693"
	data := makeRequest(url, 2)[0]
	timezone_id := data["timezone_id"]
	country_code := data["country_code"]
	if country_code == "??"{
		country_code = "Not Found"
	}
	fmt.Println("Time Zone: ", timezone_id, "Country: ", country_code)
}

func dateAndTime(){
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
	data := makeRequest(url, 3)[0]
	latitude, latOk := data["latitude"]
	longitude, lonOk := data["longitude"]
	if latOk && lonOk {
		fmt.Printf("Latitude: %f, Longitude: %f\n", latitude, longitude)
	}
}

// func printMenu(){
// 	fmt.Print("Menu:\n" +
// 	"1 - current Latitude and Longitude of ISS\n" +
// 	"2 - current Time Zone and Country code of ISS\n" +
// 	"3 - Latitude and Longitude based on Date and Time\n" +
// 	"5 - Exit\n\n")
// }

// func main(){
// 	var option int
// 	printMenu()
// 	for{
// 		fmt.Print("Choose an Option (enter 4 to see Menu): ")
// 		fmt.Scan(&option)
// 		if(option == 1){
// 			latAngLong()
// 		}else if(option == 2){
// 			timeAndCountry()
// 		}else if(option == 3){
// 			dateAndTime()
// 		}else if(option == 4){
// 			printMenu()
// 		}else if(option == 5){
// 			fmt.Println("Exit")
// 			break
// 		}
// 		fmt.Print("\n")
// 	}
// }

func main() {
	option := flag.Int("option", 0, "Choose an option: 1 for Latitude and Longitude, 2 for Time and Country, 3 for Date and Time, 4 to show menu, 5 to exit.")
	flag.Parse()

	// if *option == 0 {
	// 	printMenu()
	// 	os.Exit(1)
	// }

	switch *option {
	case 1:
		latAngLong()
	case 2:
		timeAndCountry()
	case 3:
		dateAndTime()
	default:
		fmt.Println("Invalid option. Use -h for help.")
		os.Exit(1)
	}
}
