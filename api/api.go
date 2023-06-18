package api

import (
	"fmt"
	"net/url"
	"sdkflight/core"
	request_ "sdkflight/request"
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

func (api *FlightRadar24API) Login(user, password string) (map[string]string, error) {
	data := url.Values{
		"email":    {user},
		"password": {password},
		"remember": {"true"},
		"type":     {"web"},
	}
	core := core.NewCore()

	request, err := request_.NewAPIRequest(core.UserLoginURL, nil, core.JSONHeaders, data)
	if err != nil {
		return nil, err
	}
	cookie, err := request.GetCookie("_frPl")
	if err != nil {
		return nil, err
	}
	api.realTimeFlightTrackerConfig["enc"] = cookie.Value

	content, err := request.GetContent()
	if err != nil {
		return nil, err
	}

	if result, ok := content.(map[string]string); ok {
		return result, nil
	}

	return nil, fmt.Errorf("unexpected content type")
}
