package api

import (
	"fmt"
	"io"
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
	core := core.NewCore()
	payload := map[string]io.Reader{
		"email":    strings.NewReader(user),
		"password": strings.NewReader(password),
		"remember": strings.NewReader("true"),
		"type":     strings.NewReader("web"),
	}

	req, err := request.NewAPIRequest(core.UserLoginURL, nil, core.JSONHeaders, payload).SendRequest()
	if err != nil {
		return nil, err
	}

	content, err := req.GetContent()
	if err != nil {
		return nil, err
	}

	result, ok := content.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected content type")
	}

	if result["status"] != "success" {
		return result, fmt.Errorf("login failed: %s", result["message"])
	}

	cookie, err := req.GetCookie("_frPl")
	if err != nil {
		return nil, fmt.Errorf("cannot get cookie")
	}
	api.realTimeFlightTrackerConfig["enc"] = cookie.Value

	return result, nil
}

func (api *FlightRadar24API) GetZones() (map[string]interface{}, error) {
	core := core.NewCore()

	req, err := request.NewAPIRequest(core.ZonesDataURL, nil, core.JSONHeaders, nil).SendRequest()
	if err != nil {
		return nil, err
	}

	content, err := req.GetContent()
	if err != nil {
		return nil, err
	}

	result, ok := content.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected content type")
	}

	return result, nil

}
