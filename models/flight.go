package models

import (
	"fmt"
	_package "github.com/convitnhodev/flight-radar-go-sdk/package"
)

const (
	defaultText = "N/A"
)

type Flight struct {
	Id                                    string
	Icao24bit                             string
	Latitude                              string
	Longitude                             string
	Heading                               string
	Altitude                              string
	GroundSpeed                           string
	Squawk                                string
	AircraftCode                          string
	Registration                          string
	Time                                  string
	OriginAirportIATA                     string
	DestinationAirportIATA                string
	Number                                string
	AirlineIATA                           string
	OnGround                              string
	VerticalSpeed                         string
	CallSign                              string
	AirlineICAO                           string
	AircraftAge                           string
	AircraftCountryID                     string
	AircraftHistory                       []interface{}
	AircraftImages                        []interface{}
	AircraftModel                         string
	AirlineName                           string
	AirlineShortName                      string
	DestinationAirportAltitude            string
	DestinationAirportCountryCode         string
	DestinationAirportCountryName         string
	DestinationAirportLatitude            string
	DestinationAirportLongitude           string
	DestinationAirportICAO                string
	DestinationAirportBaggage             string
	DestinationAirportGate                string
	DestinationAirportName                string
	DestinationAirportTerminal            string
	DestinationAirportVisible             string
	DestinationAirportWebsite             string
	DestinationAirportTimezoneAbbr        string
	DestinationAirportTimezoneAbbrName    string
	DestinationAirportTimezoneName        string
	DestinationAirportTimezoneOffset      string
	DestinationAirportTimezoneOffsetHours string
	OriginAirportAltitude                 string
	OriginAirportCountryCode              string
	OriginAirportCountryName              string
	OriginAirportLatitude                 string
	OriginAirportLongitude                string
	OriginAirportICAO                     string
	OriginAirportBaggage                  string
	OriginAirportGate                     string
	OriginAirportName                     string
	OriginAirportTerminal                 string
	OriginAirportVisible                  string
	OriginAirportWebsite                  string
	OriginAirportTimezoneAbbr             string
	OriginAirportTimezoneAbbrName         string
	OriginAirportTimezoneName             string
	OriginAirportTimezoneOffset           string
	OriginAirportTimezoneOffsetHours      string
	StatusIcon                            string
	StatusText                            string
	TimeDetails                           map[string]interface{}
	Trail                                 []interface{}
}

func NewFlight(flightID string, info []interface{}) *Flight {
	flight := &Flight{
		Id:                     flightID,
		Icao24bit:              getFlightInfo(info[0]),
		Latitude:               getFlightInfo(info[1]),
		Longitude:              getFlightInfo(info[2]),
		Heading:                getFlightInfo(info[3]),
		Altitude:               getFlightInfo(info[4]),
		GroundSpeed:            getFlightInfo(info[5]),
		Squawk:                 getFlightInfo(info[6]),
		AircraftCode:           getFlightInfo(info[8]),
		Registration:           getFlightInfo(info[9]),
		Time:                   getFlightInfo(info[10]),
		OriginAirportIATA:      getFlightInfo(info[11]),
		DestinationAirportIATA: getFlightInfo(info[12]),
		Number:                 getFlightInfo(info[13]),
		AirlineIATA:            getFlightInfo(info[13].(string)[:2]),
		OnGround:               getFlightInfo(info[14]),
		VerticalSpeed:          getFlightInfo(info[15]),
		CallSign:               getFlightInfo(info[16]),
		AirlineICAO:            getFlightInfo(info[18]),
	}
	return flight
}

func NewDetailFlight(flightID string, info []interface{}) *Flight {
	flight := &Flight{
		Id:                                    flightID,
		Icao24bit:                             getFlightInfo(info[0]),
		Latitude:                              getFlightInfo(info[1]),
		Longitude:                             getFlightInfo(info[2]),
		Heading:                               getFlightInfo(info[3]),
		Altitude:                              getFlightInfo(info[4]),
		GroundSpeed:                           getFlightInfo(info[5]),
		Squawk:                                getFlightInfo(info[6]),
		AircraftCode:                          getFlightInfo(info[8]),
		Registration:                          getFlightInfo(info[9]),
		Time:                                  getFlightInfo(info[10]),
		OriginAirportIATA:                     getFlightInfo(info[11]),
		DestinationAirportIATA:                getFlightInfo(info[12]),
		Number:                                getFlightInfo(info[13]),
		AirlineIATA:                           getFlightInfo(info[13].(string)[:2]),
		OnGround:                              getFlightInfo(info[14]),
		VerticalSpeed:                         getFlightInfo(info[15]),
		CallSign:                              getFlightInfo(info[16]),
		AirlineICAO:                           getFlightInfo(info[17]),
		AircraftAge:                           getFlightInfo(info[18]),
		AircraftCountryID:                     getFlightInfo(info[19]),
		AircraftHistory:                       info[20].([]interface{}),
		AircraftImages:                        info[21].([]interface{}),
		AircraftModel:                         getFlightInfo(info[22]),
		AirlineName:                           getFlightInfo(info[23]),
		AirlineShortName:                      getFlightInfo(info[24]),
		DestinationAirportAltitude:            getFlightInfo(info[25]),
		DestinationAirportCountryCode:         getFlightInfo(info[26]),
		DestinationAirportCountryName:         getFlightInfo(info[27]),
		DestinationAirportLatitude:            getFlightInfo(info[28]),
		DestinationAirportLongitude:           getFlightInfo(info[29]),
		DestinationAirportICAO:                getFlightInfo(info[30]),
		DestinationAirportBaggage:             getFlightInfo(info[31]),
		DestinationAirportGate:                getFlightInfo(info[32]),
		DestinationAirportName:                getFlightInfo(info[33]),
		DestinationAirportTerminal:            getFlightInfo(info[34]),
		DestinationAirportVisible:             getFlightInfo(info[35]),
		DestinationAirportWebsite:             getFlightInfo(info[36]),
		DestinationAirportTimezoneAbbr:        getFlightInfo(info[37]),
		DestinationAirportTimezoneAbbrName:    getFlightInfo(info[38]),
		DestinationAirportTimezoneName:        getFlightInfo(info[39]),
		DestinationAirportTimezoneOffset:      getFlightInfo(info[40]),
		DestinationAirportTimezoneOffsetHours: getFlightInfo(info[41]),
		OriginAirportAltitude:                 getFlightInfo(info[42]),
		OriginAirportCountryCode:              getFlightInfo(info[43]),
		OriginAirportCountryName:              getFlightInfo(info[44]),
		OriginAirportLatitude:                 getFlightInfo(info[45]),
		OriginAirportLongitude:                getFlightInfo(info[46]),
		OriginAirportICAO:                     getFlightInfo(info[47]),
		OriginAirportBaggage:                  getFlightInfo(info[48]),
		OriginAirportGate:                     getFlightInfo(info[49]),
		OriginAirportName:                     getFlightInfo(info[50]),
		OriginAirportTerminal:                 getFlightInfo(info[51]),
		OriginAirportVisible:                  getFlightInfo(info[52]),
		OriginAirportWebsite:                  getFlightInfo(info[53]),
		OriginAirportTimezoneAbbr:             getFlightInfo(info[54]),
		OriginAirportTimezoneAbbrName:         getFlightInfo(info[55]),
		OriginAirportTimezoneName:             getFlightInfo(info[56]),
		OriginAirportTimezoneOffset:           getFlightInfo(info[57]),
		OriginAirportTimezoneOffsetHours:      getFlightInfo(info[58]),
		StatusIcon:                            getFlightInfo(info[59]),
		StatusText:                            getFlightInfo(info[60]),
		TimeDetails:                           info[61].(map[string]interface{}),
		Trail:                                 info[62].([]interface{}),
	}
	return flight
}

func getFlightInfo(info interface{}) string {
	if (info != nil || info == 0) && info != defaultText {
		return _package.ConvertToString(info)
	} else {
		return defaultText
	}
}

func (f *Flight) String() string {
	template := "<(%s) %s - Altitude: %s - Ground Speed: %s - Heading: %s>"
	return fmt.Sprintf(template, f.AircraftCode, f.Registration, f.Altitude, f.GroundSpeed, f.Heading)
}

func (f *Flight) GetFlightID() string {
	return f.Id
}

// func main() {
// 	info := []interface{}{"ABC123", "ICAO123", "51.5074", "-0.1278", "123", "250", "1234", "XYZ", "123", "2023-06-17", "LHR", "JFK", "ABC123", "AA", "0", "100", "CSN", "ICAO123", "5", "US", []interface{}{}, []interface{}{}, "Boeing 747", "American Airlines", "AA", "100", "US", "UK", "51.5074", "-0.1278", "ICAO123", "N/A", "N/A", "London Heathrow Airport", "Terminal 2", "1", "https://www.heathrow.com/", "BST", "British Summer Time", "Europe/London", "1", "0", "1234", "US", "United States", "40.7128", "-74.0060", "ICAO456", "N/A", "N/A", "John F. Kennedy International Airport", "Terminal 4", "1", "https://www.jfkairport.com/", "EDT", "Eastern Daylight Time", "America/New_York", "-4", "-4", "✈️", "Scheduled", map[string]interface{}{}, []interface{}{}}

// 	flight := NewFlight("123456", info)
// 	fmt.Println(flight)
// }
