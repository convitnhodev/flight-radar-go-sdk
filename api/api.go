package api

import (
	"errors"
	"fmt"
	"github.com/convitnhodev/flight-radar-go-sdk/models"
	_package "github.com/convitnhodev/flight-radar-go-sdk/package"
	"io"
	"net/url"
	"strconv"
	"strings"

	"github.com/convitnhodev/flight-radar-go-sdk/core"
	"github.com/convitnhodev/flight-radar-go-sdk/request"
)

type FlightRadar24API struct {
	realTimeFlightTrackerConfig map[string]string
}

func NewFlightRadar24API() *FlightRadar24API {
	return &FlightRadar24API{
		realTimeFlightTrackerConfig: map[string]string{
			"faa":       "1",
			"satellite": "1",
			"mlat":      "1",
			"flarm":     "1",
			"adsb":      "1",
			"gnd":       "1",
			"air":       "1",
			"vehicles":  "1",
			"estimated": "1",
			"maxage":    "14400",
			"gliders":   "1",
			"stats":     "1",
			"limit":     "5000",
		},
	}
}

func NewCustomFlightRadar24API(config map[string]string) *FlightRadar24API {
	return &FlightRadar24API{
		realTimeFlightTrackerConfig: config,
	}
}

func (api *FlightRadar24API) Login(user, password string) (map[string]interface{}, error) {
	payload := map[string]io.Reader{
		"email":    strings.NewReader(user),
		"password": strings.NewReader(password),
		"remember": strings.NewReader("true"),
		"type":     strings.NewReader("web"),
	}

	req, err := request.NewAPIRequest(core.UserLoginURL, nil, core.JSONHeaders, payload).SendRequest()
	if err != nil {
		return nil, fmt.Errorf("cannot send request: %s", err.Error())
	}

	content, err := req.GetContent()
	if err != nil {
		return nil, fmt.Errorf("cannot get content: %s", err.Error())
	}

	result, ok := content.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected content type")
	}

	if status, ok := result["status"]; !ok || status != "success" {
		return result, fmt.Errorf("login failed: %s", result["message"])
	}

	cookie, err := req.GetCookie("_frPl")
	if err != nil {
		return result, fmt.Errorf("cannot get cookie: %s", err.Error())
	}
	api.realTimeFlightTrackerConfig["enc"] = cookie.Value

	return result, nil
}

// Get the data from Flightradar24.
func (api *FlightRadar24API) GetAirlines() ([]map[string]interface{}, error) {
	req, err := request.NewAPIRequest(core.AirlinesDataURL, nil, core.JSONHeaders, nil).SendRequest()
	if err != nil {
		return nil, fmt.Errorf("cannot send request: %s", err.Error())
	}

	content, err := req.GetContent()
	if err != nil {
		return nil, fmt.Errorf("cannot get content: %s", err.Error())
	}

	result, ok := content.(map[string]interface{})["rows"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected content type")
	}

	rows := make([]map[string]interface{}, len(result))
	for i, row := range result {
		rows[i] = row.(map[string]interface{})
	}

	return rows, nil
}

func (api *FlightRadar24API) GetZones() (map[string]interface{}, error) {
	req, err := request.NewAPIRequest(core.ZonesDataURL, nil, core.JSONHeaders, nil).SendRequest()
	if err != nil {
		return nil, fmt.Errorf("cannot send request: %s", err.Error())
	}

	content, err := req.GetContent()
	if err != nil {
		return nil, fmt.Errorf("cannot get content: %s", err.Error())
	}

	result, ok := content.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected content type")
	}

	return result, nil
}

func (api *FlightRadar24API) GetRealTimeFlightTrackerConfig() map[string]string {
	result := api.realTimeFlightTrackerConfig
	return result
}

func (api *FlightRadar24API) GetCountryFlag(country string) (string, error) {
	flagUrl := core.CountryFlagURL
	formattedCountry := strings.ReplaceAll(strings.ToLower(country), " ", "-")
	formattedFlagUrl := strings.ReplaceAll(flagUrl, "{}", formattedCountry)

	headers := core.ImageHeaders
	if _, ok := headers["Origin"]; ok {
		headers.Del("Origin")
	}

	req, err := request.NewAPIRequest(formattedFlagUrl, nil, headers, nil).SendRequest()
	if err != nil {
		return "", err
	}

	statusCode := req.GetStatusCode()
	if !strings.HasPrefix(strconv.Itoa(statusCode), "4") {
		return formattedFlagUrl, nil
	}
	return "", errors.New("invalid country flag")

}

func (api *FlightRadar24API) GetDetailFlight(flightId string) (map[string]interface{}, error) {
	detailFlightUrl := core.FlightDataURL
	formattedDetailFlightUrl := strings.ReplaceAll(detailFlightUrl, "{}", flightId)
	req, err := request.NewAPIRequest(formattedDetailFlightUrl, nil, core.JSONHeaders, nil).SendRequest()
	if err != nil {
		return nil, err
	}

	content, err := req.GetContent()
	if err != nil {
		return nil, fmt.Errorf("cannot get content: %s", err.Error())
	}

	result, ok := content.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected content type")
	}

	return result, err
}

func (api *FlightRadar24API) SetRealTimeFlightTrackerConfig(config map[string]string) {
	for key, value := range config {
		if _, ok := api.realTimeFlightTrackerConfig[key]; ok && _package.IsNumeric(value) {
			api.realTimeFlightTrackerConfig[key] = value
		}
	}
}

func (api *FlightRadar24API) GetFlights(airline *string, bounds *string, registration *string, aircraftType *string) ([]*models.Flight, error) {
	requestParams := api.realTimeFlightTrackerConfig

	if airline != nil {
		requestParams["airline"] = *airline
	}

	if bounds != nil {
		requestParams["bounds"] = strings.Replace(*bounds, ",", "%2C", -1)
	}

	if registration != nil {
		requestParams["reg"] = *registration
	}

	if aircraftType != nil {
		requestParams["type"] = *aircraftType
	}

	values := url.Values{}
	for key, value := range requestParams {
		values.Add(key, value)
	}

	req, err := request.NewAPIRequest(core.RealTimeFlightTrackerDataURL, values, core.JSONHeaders, nil).SendRequest()
	if err != nil {
		return nil, err
	}

	content, err := req.GetContent()
	if err != nil {
		return nil, fmt.Errorf("cannot get content: %s", err.Error())
	}

	result, ok := content.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected content type")
	}

	flights := make([]*models.Flight, 0)

	for id, info := range result {
		if _package.IsNumeric(string(id[0])) {
			value := info.([]interface{})
			flights = append(flights, models.NewFlight(id, value))
		}
	}

	fmt.Println(result)

	return flights, nil
}
