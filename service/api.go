package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/convitnhodev/flight-radar-go-sdk/package/html_package"

	"github.com/convitnhodev/flight-radar-go-sdk/models"
	_package "github.com/convitnhodev/flight-radar-go-sdk/package"

	"github.com/convitnhodev/flight-radar-go-sdk/component"
	"github.com/convitnhodev/flight-radar-go-sdk/transport"
)

type FlightRadar24API struct {
	realTimeFlightTrackerConfig map[string]string
	cookies                     []*http.Cookie
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
			"limit":     "5000", // TODO: change limit to 5000
		},
	}
}

func NewCustomFlightRadar24API(config map[string]string) *FlightRadar24API {
	return &FlightRadar24API{
		realTimeFlightTrackerConfig: config,
	}
}

func (api *FlightRadar24API) GetCookies() []*http.Cookie {
	return api.cookies
}

func (api *FlightRadar24API) SetCokkies(ck []*http.Cookie) {
	api.cookies = ck
	for _, c := range ck {
		if c.Name == "_frPl" {
			api.realTimeFlightTrackerConfig["enc"] = c.Value
			break
		}
	}
}

func (api *FlightRadar24API) Login(user, password string) (map[string]interface{}, error) {
	payload := map[string]io.Reader{
		"email":    strings.NewReader(user),
		"password": strings.NewReader(password),
		"remember": strings.NewReader("true"),
		"type":     strings.NewReader("web"),
	}

	req, err := transport.NewAPIRequest(component.UserLoginURL, nil, component.JSONHeaders, payload).SendRequest()
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

	cookies := req.GetCookies()
	api.cookies = cookies

	return result, nil
}

// Get the data from Flightradar24.
func (api *FlightRadar24API) GetAirlines() ([]map[string]interface{}, error) {
	req, err := transport.NewAPIRequest(component.AirlinesDataURL, nil, component.JSONHeaders, nil).SendRequest()
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

func (api *FlightRadar24API) GetAirlineLogo(iata string, icao string) (string, error) {
	flagUrl := component.AirlineLogoURL
	replacedPath := strings.ReplaceAll(flagUrl, "{}_{}", fmt.Sprintf("%s_%s", iata, icao))

	// Get the first airline logo URL.
	req, err := transport.NewAPIRequest(replacedPath, nil, component.ImageHeaders, nil).SendRequest()
	if err != nil {
		return "", err
	}

	statusCode := req.GetStatusCode()
	if !strings.HasPrefix(strconv.Itoa(statusCode), "4") {
		return replacedPath, nil
	}

	// Get the second airline logo URL.
	secondReplacedPath := strings.ReplaceAll(component.AlternativeAirlineLogoURL, "{}", icao)
	secondReq, err := transport.NewAPIRequest(secondReplacedPath, nil, component.ImageHeaders, nil).SendRequest()
	if err != nil {
		return "", err
	}

	statusCode = secondReq.GetStatusCode()
	if !strings.HasPrefix(strconv.Itoa(statusCode), "4") {
		return secondReplacedPath, nil
	}
	return "", nil
}

func (api *FlightRadar24API) GetAirports() ([]map[string]interface{}, error) {
	req, err := transport.NewAPIRequest(component.AirportsDataURL, nil, component.JSONHeaders, nil).SendRequest()
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

func (api *FlightRadar24API) GetAirport(code string) (map[string]interface{}, error) {

	detailAirportURL := component.AirportDataURL
	formattedDetailAirportUrl := strings.ReplaceAll(detailAirportURL, "{}", code)
	req, err := transport.NewAPIRequest(formattedDetailAirportUrl, nil, component.JSONHeaders, nil).SendRequest()
	if err != nil {
		return nil, err
	}

	content, err := req.GetContent()
	if err != nil {
		return nil, fmt.Errorf("cannot get content: %s", err.Error())
	}

	result, ok := content.(map[string]interface{})["details"].((map[string]interface{}))
	if !ok {
		return nil, fmt.Errorf("unexpected content type")
	}

	return result, nil

}

func (api *FlightRadar24API) GetZones() (map[string]interface{}, error) {
	req, err := transport.NewAPIRequest(component.ZonesDataURL, nil, component.JSONHeaders, nil).SendRequest()
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
	flagUrl := component.CountryFlagURL
	formattedCountry := strings.ReplaceAll(strings.ToLower(country), " ", "-")
	formattedFlagUrl := strings.ReplaceAll(flagUrl, "{}", formattedCountry)

	headers := component.ImageHeaders
	if _, ok := headers["Origin"]; ok {
		headers.Del("Origin")
	}

	req, err := transport.NewAPIRequest(formattedFlagUrl, nil, headers, nil).SendRequest()
	if err != nil {
		return "", err
	}

	statusCode := req.GetStatusCode()
	if !strings.HasPrefix(strconv.Itoa(statusCode), "4") {
		return formattedFlagUrl, nil
	}
	return "", errors.New("invalid country flag")

}

func (api *FlightRadar24API) GetDetailFlight(flightId string) (*models.DetailFlight, error) {
	detailFlightUrl := component.FlightDataURL
	formattedDetailFlightUrl := strings.ReplaceAll(detailFlightUrl, "{}", flightId)
	req, err := transport.NewAPIRequest(formattedDetailFlightUrl, nil, component.JSONHeaders, nil).SendRequest()
	if err != nil {
		return nil, err
	}

	content, err := req.GetContent()
	if err != nil {
		return nil, fmt.Errorf("cannot get content: %s", err.Error())
	}

	jsonData, err := json.Marshal(content)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	var result models.DetailFlight

	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, fmt.Errorf("unexpected content type")
	}

	return &result, err
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

	req, err := transport.NewAPIRequest(component.RealTimeFlightTrackerDataURL, values, component.JSONHeaders, nil).SendRequest()
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

	return flights, nil
}

func (api *FlightRadar24API) GetFlightsByFlightNo(flightNo string, limit int) (interface{}, error) {
	url := component.FlightSearchURL
	url = strings.ReplaceAll(url, "{}", flightNo)
	url = strings.ReplaceAll(url, "{2}", _package.ConvertToString(limit))

	req, err := transport.NewAPIRequest(url, nil, component.JSONHeaders, nil).SendRequest()

	if err != nil {
		return nil, err
	}

	content, err := req.GetContent()
	if err != nil {
		return nil, fmt.Errorf("cannot get content: %s", err.Error())
	}

	return content, nil

}

func (api *FlightRadar24API) GetAllFlightWithKey(keySearch string) ([]models.CoreFlight, error) {
	url := component.FlightSearchALLUrl
	url = strings.ReplaceAll(url, "{}", keySearch)
	req, err := transport.NewAPIRequest(url, nil, component.Headers, nil).AddCookies(api.cookies).SendRequest()
	if err != nil {
		return nil, err
	}

	content, err := req.GetContent()
	if err != nil {
		return nil, fmt.Errorf("cannot get content: %s", err.Error())
	}

	result := string(content.([]uint8))
	a := html_package.ConvertArrayRawModelToModel(html_package.GetHTML(result))
	return a, nil

}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
