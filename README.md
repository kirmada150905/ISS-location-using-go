# ISS Position and Time API Client

This Go application fetches the current latitude and longitude of the International Space Station (ISS) and provides information about its time zone and country based on geographic coordinates. Additionally, the application allows you to retrieve the ISS's latitude and longitude for a specified date and time.

## Usage

To run the application, use the following commands:

```bash
go run four.go -h         # Display help information
go run four.go -option 1  # Option 1: Fetch current ISS location
go run four.go -option 2  # Option 2: Fetch current ISS time zone and country code
go run four.go -option 3  # Option 3: Fetch ISS location on a particular date and time
