package flight

import (
	"fmt"
)

type Flight struct {
	__defaultText                         string
	id                                    string
	icao24bit                             string
	latitude                              string
	longitude                             string
	heading                               string
	altitude                              string
	groundSpeed                           string
	squawk                                string
	aircraftCode                          string
	registration                          string
	time                                  string
	originAirportIATA                     string
	destinationAirportIATA                string
	number                                string
	airlineIATA                           string
	onGround                              string
	verticalSpeed                         string
	callsign                              string
	airlineICAO                           string
	aircraftAge                           string
	aircraftCountryID                     string
	aircraftHistory                       []interface{}
	aircraftImages                        []interface{}
	aircraftModel                         string
	airlineName                           string
	airlineShortName                      string
	destinationAirportAltitude            string
	destinationAirportCountryCode         string
	destinationAirportCountryName         string
	destinationAirportLatitude            string
	destinationAirportLongitude           string
	destinationAirportICAO                string
	destinationAirportBaggage             string
	destinationAirportGate                string
	destinationAirportName                string
	destinationAirportTerminal            string
	destinationAirportVisible             string
	destinationAirportWebsite             string
	destinationAirportTimezoneAbbr        string
	destinationAirportTimezoneAbbrName    string
	destinationAirportTimezoneName        string
	destinationAirportTimezoneOffset      string
	destinationAirportTimezoneOffsetHours string
	originAirportAltitude                 string
	originAirportCountryCode              string
	originAirportCountryName              string
	originAirportLatitude                 string
	originAirportLongitude                string
	originAirportICAO                     string
	originAirportBaggage                  string
	originAirportGate                     string
	originAirportName                     string
	originAirportTerminal                 string
	originAirportVisible                  string
	originAirportWebsite                  string
	originAirportTimezoneAbbr             string
	originAirportTimezoneAbbrName         string
	originAirportTimezoneName             string
	originAirportTimezoneOffset           string
	originAirportTimezoneOffsetHours      string
	statusIcon                            string
	statusText                            string
	timeDetails                           map[string]interface{}
	trail                                 []interface{}
}

func NewFlight(flightID string, info []interface{}) *Flight {
	flight := &Flight{
		__defaultText:                         "N/A",
		id:                                    flightID,
		icao24bit:                             getFlightInfo(info[0]),
		latitude:                              getFlightInfo(info[1]),
		longitude:                             getFlightInfo(info[2]),
		heading:                               getFlightInfo(info[3]),
		altitude:                              getFlightInfo(info[4]),
		groundSpeed:                           getFlightInfo(info[5]),
		squawk:                                getFlightInfo(info[6]),
		aircraftCode:                          getFlightInfo(info[8]),
		registration:                          getFlightInfo(info[9]),
		time:                                  getFlightInfo(info[10]),
		originAirportIATA:                     getFlightInfo(info[11]),
		destinationAirportIATA:                getFlightInfo(info[12]),
		number:                                getFlightInfo(info[13].(string)[:2]),
		airlineIATA:                           getFlightInfo(info[13].(string)[:2]),
		onGround:                              getFlightInfo(info[14]),
		verticalSpeed:                         getFlightInfo(info[15]),
		callsign:                              getFlightInfo(info[16]),
		airlineICAO:                           getFlightInfo(info[17]),
		aircraftAge:                           getFlightInfo(info[18]),
		aircraftCountryID:                     getFlightInfo(info[19]),
		aircraftHistory:                       info[20].([]interface{}),
		aircraftImages:                        info[21].([]interface{}),
		aircraftModel:                         getFlightInfo(info[22]),
		airlineName:                           getFlightInfo(info[23]),
		airlineShortName:                      getFlightInfo(info[24]),
		destinationAirportAltitude:            getFlightInfo(info[25]),
		destinationAirportCountryCode:         getFlightInfo(info[26]),
		destinationAirportCountryName:         getFlightInfo(info[27]),
		destinationAirportLatitude:            getFlightInfo(info[28]),
		destinationAirportLongitude:           getFlightInfo(info[29]),
		destinationAirportICAO:                getFlightInfo(info[30]),
		destinationAirportBaggage:             getFlightInfo(info[31]),
		destinationAirportGate:                getFlightInfo(info[32]),
		destinationAirportName:                getFlightInfo(info[33]),
		destinationAirportTerminal:            getFlightInfo(info[34]),
		destinationAirportVisible:             getFlightInfo(info[35]),
		destinationAirportWebsite:             getFlightInfo(info[36]),
		destinationAirportTimezoneAbbr:        getFlightInfo(info[37]),
		destinationAirportTimezoneAbbrName:    getFlightInfo(info[38]),
		destinationAirportTimezoneName:        getFlightInfo(info[39]),
		destinationAirportTimezoneOffset:      getFlightInfo(info[40]),
		destinationAirportTimezoneOffsetHours: getFlightInfo(info[41]),
		originAirportAltitude:                 getFlightInfo(info[42]),
		originAirportCountryCode:              getFlightInfo(info[43]),
		originAirportCountryName:              getFlightInfo(info[44]),
		originAirportLatitude:                 getFlightInfo(info[45]),
		originAirportLongitude:                getFlightInfo(info[46]),
		originAirportICAO:                     getFlightInfo(info[47]),
		originAirportBaggage:                  getFlightInfo(info[48]),
		originAirportGate:                     getFlightInfo(info[49]),
		originAirportName:                     getFlightInfo(info[50]),
		originAirportTerminal:                 getFlightInfo(info[51]),
		originAirportVisible:                  getFlightInfo(info[52]),
		originAirportWebsite:                  getFlightInfo(info[53]),
		originAirportTimezoneAbbr:             getFlightInfo(info[54]),
		originAirportTimezoneAbbrName:         getFlightInfo(info[55]),
		originAirportTimezoneName:             getFlightInfo(info[56]),
		originAirportTimezoneOffset:           getFlightInfo(info[57]),
		originAirportTimezoneOffsetHours:      getFlightInfo(info[58]),
		statusIcon:                            getFlightInfo(info[59]),
		statusText:                            getFlightInfo(info[60]),
		timeDetails:                           info[61].(map[string]interface{}),
		trail:                                 info[62].([]interface{}),
	}
	return flight
}

func getFlightInfo(info interface{}) string {
	if val, ok := info.(string); ok && (val != "" && val != "0") {
		return val
	}
	return "__defaultText"
}

func (f *Flight) String() string {
	template := "<(%s) %s - Altitude: %s - Ground Speed: %s - Heading: %s>"
	return fmt.Sprintf(template, f.aircraftCode, f.registration, f.altitude, f.groundSpeed, f.heading)
}

func (f *Flight) GetFlightID() string {
	return f.id
}

//func main() {
//	info := []interface{}{"ABC123", "ICAO123", "51.5074", "-0.1278", "123", "250", "1234", "XYZ", "123", "2023-06-17", "LHR", "JFK", "ABC123", "AA", "0", "100", "CSN", "ICAO123", "5", "US", []interface{}{}, []interface{}{}, "Boeing 747", "American Airlines", "AA", "100", "US", "UK", "51.5074", "-0.1278", "ICAO123", "N/A", "N/A", "London Heathrow Airport", "Terminal 2", "1", "https://www.heathrow.com/", "BST", "British Summer Time", "Europe/London", "1", "0", "1234", "US", "United States", "40.7128", "-74.0060", "ICAO456", "N/A", "N/A", "John F. Kennedy International Airport", "Terminal 4", "1", "https://www.jfkairport.com/", "EDT", "Eastern Daylight Time", "America/New_York", "-4", "-4", "✈️", "Scheduled", map[string]interface{}{}, []interface{}{}}
//
//	flight := NewFlight("123456", info)
//	fmt.Println(flight)
//}
