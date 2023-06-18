package models

import (
	"fmt"

	_package "github.com/convitnhodev/flight-radar-go-sdk/package"
)

const (
	defaultText string = "N/A"
)

type Flight struct {
	Id                                    string                 `json:"id,omitempty"`
	Icao24bit                             string                 `json:"icao24bit,omitempty"`
	Latitude                              string                 `json:"latitude,omitempty"`
	Longitude                             string                 `json:"longitude,omitempty"`
	Heading                               string                 `json:"heading,omitempty"`
	Altitude                              string                 `json:"altitude,omitempty"`
	GroundSpeed                           string                 `json:"ground_speed,omitempty"`
	Squawk                                string                 `json:"squawk,omitempty"`
	AircraftCode                          string                 `json:"aircraft_code,omitempty"`
	Registration                          string                 `json:"registration,omitempty"`
	Time                                  string                 `json:"time,omitempty"`
	OriginAirportIATA                     string                 `json:"origin_airport_iata,omitempty"`
	DestinationAirportIATA                string                 `json:"destination_airport_iata,omitempty"`
	Number                                string                 `json:"number,omitempty"`
	AirlineIATA                           string                 `json:"airline_iata,omitempty"`
	OnGround                              string                 `json:"on_ground,omitempty"`
	VerticalSpeed                         string                 `json:"vertical_speed,omitempty"`
	CallSign                              string                 `json:"call_sign,omitempty"`
	AirlineICAO                           string                 `json:"airline_icao,omitempty"`
	AircraftAge                           string                 `json:"aircraft_age,omitempty"`
	AircraftCountryID                     string                 `json:"aircraft_country_id,omitempty"`
	AircraftHistory                       []interface{}          `json:"aircraft_history,omitempty"`
	AircraftImages                        []interface{}          `json:"aircraft_images,omitempty"`
	AircraftModel                         string                 `json:"aircraft_model,omitempty"`
	AirlineName                           string                 `json:"airline_name,omitempty"`
	AirlineShortName                      string                 `json:"airline_short_name,omitempty"`
	DestinationAirportAltitude            string                 `json:"destination_airport_altitude,omitempty"`
	DestinationAirportCountryCode         string                 `json:"destination_airport_country_code,omitempty"`
	DestinationAirportCountryName         string                 `json:"destination_airport_country_name,omitempty"`
	DestinationAirportLatitude            string                 `json:"destination_airport_latitude,omitempty"`
	DestinationAirportLongitude           string                 `json:"destination_airport_longitude,omitempty"`
	DestinationAirportICAO                string                 `json:"destination_airport_icao,omitempty"`
	DestinationAirportBaggage             string                 `json:"destination_airport_baggage,omitempty"`
	DestinationAirportGate                string                 `json:"destination_airport_gate,omitempty"`
	DestinationAirportName                string                 `json:"destination_airport_name,omitempty"`
	DestinationAirportTerminal            string                 `json:"destination_airport_terminal,omitempty"`
	DestinationAirportVisible             string                 `json:"destination_airport_visible,omitempty"`
	DestinationAirportWebsite             string                 `json:"destination_airport_website,omitempty"`
	DestinationAirportTimezoneAbbr        string                 `json:"destination_airport_timezone_abbr,omitempty"`
	DestinationAirportTimezoneAbbrName    string                 `json:"destination_airport_timezone_abbr_name,omitempty"`
	DestinationAirportTimezoneName        string                 `json:"destination_airport_timezone_name,omitempty"`
	DestinationAirportTimezoneOffset      string                 `json:"destination_airport_timezone_offset,omitempty"`
	DestinationAirportTimezoneOffsetHours string                 `json:"destination_airport_timezone_offset_hours,omitempty"`
	OriginAirportAltitude                 string                 `json:"origin_airport_altitude,omitempty"`
	OriginAirportCountryCode              string                 `json:"origin_airport_country_code,omitempty"`
	OriginAirportCountryName              string                 `json:"origin_airport_country_name,omitempty"`
	OriginAirportLatitude                 string                 `json:"origin_airport_latitude,omitempty"`
	OriginAirportLongitude                string                 `json:"origin_airport_longitude,omitempty"`
	OriginAirportICAO                     string                 `json:"origin_airport_icao,omitempty"`
	OriginAirportBaggage                  string                 `json:"origin_airport_baggage,omitempty"`
	OriginAirportGate                     string                 `json:"origin_airport_gate,omitempty"`
	OriginAirportName                     string                 `json:"origin_airport_name,omitempty"`
	OriginAirportTerminal                 string                 `json:"origin_airport_terminal,omitempty"`
	OriginAirportVisible                  string                 `json:"origin_airport_visible,omitempty"`
	OriginAirportWebsite                  string                 `json:"origin_airport_website,omitempty"`
	OriginAirportTimezoneAbbr             string                 `json:"origin_airport_timezone_abbr,omitempty"`
	OriginAirportTimezoneAbbrName         string                 `json:"origin_airport_timezone_abbr_name,omitempty"`
	OriginAirportTimezoneName             string                 `json:"origin_airport_timezone_name,omitempty"`
	OriginAirportTimezoneOffset           string                 `json:"origin_airport_timezone_offset,omitempty"`
	OriginAirportTimezoneOffsetHours      string                 `json:"origin_airport_timezone_offset_hours,omitempty"`
	StatusIcon                            string                 `json:"status_icon,omitempty"`
	StatusText                            string                 `json:"status_text,omitempty"`
	TimeDetails                           map[string]interface{} `json:"time_details,omitempty"`
	Trail                                 []interface{}          `json:"trail,omitempty"`
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
		AirlineIATA:            getAirlineIATA(info[13]),
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
		AirlineIATA:                           getAirlineIATA(info[13]),
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

func getAirlineIATA(info interface{}) string {
	if info, ok := info.(string); ok {
		if len(info) > 2 {
			return info[:2]
		}
		return info
	}
	return defaultText
}

func (f *Flight) String() string {
	template := "<(%s) %s - Altitude: %s - Ground Speed: %s - Heading: %s>"
	return fmt.Sprintf(template, f.AircraftCode, f.Registration, f.Altitude, f.GroundSpeed, f.Heading)
}

func (f *Flight) GetFlightID() string {
	return f.Id
}
