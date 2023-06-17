package core

import (
	"net/http"
)

type Core struct {
	// Base URLs.
	CDNFlightradarBaseURL string
	FlightradarBaseURL    string
	DataLiveBaseURL       string
	DataCloudBaseURL      string

	// User login URL.
	UserLoginURL string

	// Flights data URLs.
	RealTimeFlightTrackerDataURL string
	FlightDataURL                string

	// Airports data URLs.
	AirportDataURL  string
	AirportsDataURL string

	// Airlines data URL.
	AirlinesDataURL string

	// Zones data URL.
	ZonesDataURL string

	// Country flag image URL.
	CountryFlagURL string

	// Airline logo image URL.
	AirlineLogoURL            string
	AlternativeAirlineLogoURL string

	Headers      http.Header
	JSONHeaders  http.Header
	ImageHeaders http.Header
}

func NewCore() *Core {
	core := &Core{
		CDNFlightradarBaseURL:        "https://cdn.flightradar24.com",
		FlightradarBaseURL:           "https://www.flightradar24.com",
		DataLiveBaseURL:              "https://data-live.flightradar24.com",
		DataCloudBaseURL:             "https://data-cloud.flightradar24.com",
		UserLoginURL:                 "https://www.flightradar24.com/user/login",
		RealTimeFlightTrackerDataURL: "",
		FlightDataURL:                "",
		AirportDataURL:               "",
		AirportsDataURL:              "",
		AirlinesDataURL:              "",
		ZonesDataURL:                 "",
		CountryFlagURL:               "",
		AirlineLogoURL:               "",
		AlternativeAirlineLogoURL:    "",
		Headers: http.Header{
			"accept-encoding": []string{"gzip, br"},
			"accept-language": []string{"pt-BR,pt;q=0.9,en-US;q=0.8,en;q=0.7"},
			"cache-control":   []string{"max-age=0"},
			"origin":          []string{"https://www.flightradar24.com"},
			"referer":         []string{"https://www.flightradar24.com/"},
			"sec-fetch-dest":  []string{"empty"},
			"sec-fetch-mode":  []string{"cors"},
			"sec-fetch-site":  []string{"same-site"},
			"user-agent":      []string{"Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36"},
		},
		JSONHeaders: http.Header{
			"Accept-Encoding": []string{"gzip, deflate, br"},
			"Accept-Language": []string{"en-US,en;q=0.9,vi;q=0.8"},
			"Cache-Control":   []string{"max-age=0"},
			"Origin":          []string{"https://www.flightradar24.com"},
			"Referer":         []string{"https://www.flightradar24.com/"},
			"Sec-Fetch-Dest":  []string{"empty"},
			"Sec-Fetch-Mode":  []string{"cors"},
			"Sec-Fetch-Site":  []string{"same-origin"},
			"User-Agent":      []string{"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"},
			"Accept":          []string{"application/json"},
		},
		ImageHeaders: http.Header{
			"accept-encoding": []string{"gzip, br"},
			"accept-language": []string{"pt-BR,pt;q=0.9,en-US;q=0.8,en;q=0.7"},
			"cache-control":   []string{"max-age=0"},
			"origin":          []string{"https://www.flightradar24.com"},
			"referer":         []string{"https://www.flightradar24.com/"},
			"sec-fetch-dest":  []string{"empty"},
			"sec-fetch-mode":  []string{"cors"},
			"sec-fetch-site":  []string{"same-site"},
			"user-agent":      []string{"Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36"},
			"accept":          []string{"image/gif, image/jpg, image/jpeg, image/png"},
		},
	}
	return core
}

//func main() {
//	// Create a new instance of Core
//	core := NewCore()
//
//	// Access the properties of the Core instance
//	println("CDNFlightradarBaseURL:", core.CDNFlightradarBaseURL)
//	println("FlightradarBaseURL:", core.FlightradarBaseURL)
//	println("DataLiveBaseURL:", core.DataLiveBaseURL)
//	println("DataCloudBaseURL:", core.DataCloudBaseURL)
//	// ...
//}
