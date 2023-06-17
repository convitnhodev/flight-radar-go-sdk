package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

//	func main() {
//		api := api.NewFlightRadar24API()
//		api.Login(`admin@tenxtenx.com`, `e9Lc5r7_tCUS!U5`)
//	}
func main() {
	url := "https://www.flightradar24.com/user/login"

	// Prepare the payload
	payload := map[string]interface{}{
		"email":    "admin@tenxtenx.com",
		"password": "e9Lc5r7_tCUS!U5",
		"remember": true,
		"type":     "web",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshaling payload:", err)
		return
	}

	// Prepare the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the request headers
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,vi;q=0.8")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Referer", "https://www.flightradar24.com/52.50,13.31/4")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Process the response
	// ...
	// Here you can handle the response body and status code
	// For example, you can read the response body using resp.Body and decode it if it's in JSON format
	// You can also check the response status code using resp.StatusCode

	// Print the response status code
	fmt.Println("Response Status Code:", resp.StatusCode)
}
