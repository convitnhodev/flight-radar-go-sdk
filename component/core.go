package component

import (
	"net/http"
)

var (
	// Base URLs.
	CDNFlightradarBaseURL string = "https://cdn.flightradar24.com"
	FlightradarBaseURL    string = "https://www.flightradar24.com"
	DataLiveBaseURL       string = "https://data-live.flightradar24.com"
	DataCloudBaseURL      string = "https://data-cloud.flightradar24.com"

	// User login URL.
	UserLoginURL string = FlightradarBaseURL + "/user/login"

	// Flights data URLs.
	RealTimeFlightTrackerDataURL string = DataCloudBaseURL + "/zones/fcgi/feed.js"
	FlightDataURL                string = DataLiveBaseURL + "/clickhandler/?flight={}"

	// Airports data URLs.
	AirportDataURL  string = FlightradarBaseURL + "/airports/traffic-stats/?airport={}"
	AirportsDataURL string = FlightradarBaseURL + "/_json/airports.php"

	// Airlines data URL.
	AirlinesDataURL string = FlightradarBaseURL + "/_json/airlines.php"

	// Zones data URL.
	ZonesDataURL string = FlightradarBaseURL + "/js/zones.js.php"

	// Country flag image URL.
	CountryFlagURL string = FlightradarBaseURL + "/static/images/data/flags-small/{}.svg"

	// Airline logo image URL.
	AirlineLogoURL            string = CDNFlightradarBaseURL + "/assets/airlines/logotypes/{}_{}.png"
	AlternativeAirlineLogoURL string = FlightradarBaseURL + "/static/images/data/operators/{}_logo0.png"

	// Search
	FlightSearchURL string = FlightradarBaseURL + "/v1/search/web/find?query={}&limit={2}"

	Headers http.Header = http.Header{
		"Accept-Encoding": []string{"gzip, br"},
		"Accept-Language": []string{"pt-BR,pt;q=0.9,en-US;q=0.8,en;q=0.7"},
		"Cache-Control":   []string{"max-age=0"},
		"Origin":          []string{"https://www.flightradar24.com"},
		"Referer":         []string{"https://www.flightradar24.com/"},
		"Sec-Fetch-Dest":  []string{"empty"},
		"Sec-Fetch-Mode":  []string{"cors"},
		"Sec-Fetch-Site":  []string{"same-site"},
		"User-Agent":      []string{"Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36"},
	}

	JSONHeaders http.Header = http.Header{
		"Accept-Encoding": []string{"gzip, br"},
		"Accept-Language": []string{"pt-BR,pt;q=0.9,en-US;q=0.8,en;q=0.7"},
		"Cache-Control":   []string{"max-age=0"},
		"Origin":          []string{"https://www.flightradar24.com"},
		"Referer":         []string{"https://www.flightradar24.com/"},
		"Sec-Fetch-Dest":  []string{"empty"},
		"Sec-Fetch-Mode":  []string{"cors"},
		"Sec-Fetch-Site":  []string{"same-site"},
		"User-Agent":      []string{"Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36"},
		"Accept":          []string{"application/json"},
	}

	ImageHeaders http.Header = http.Header{
		"Accept-Encoding": []string{"gzip, br"},
		"Accept-Language": []string{"pt-BR,pt;q=0.9,en-US;q=0.8,en;q=0.7"},
		"Cache-Control":   []string{"max-age=0"},
		"Origin":          []string{"https://www.flightradar24.com"},
		"Referer":         []string{"https://www.flightradar24.com/"},
		"Sec-Fetch-Dest":  []string{"empty"},
		"Sec-Fetch-Mode":  []string{"cors"},
		"Sec-Fetch-Site":  []string{"same-site"},
		"User-Agent":      []string{"Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36"},
		"Accept":          []string{"image/gif, image/jpg, image/jpeg, image/png"},
	}
)
