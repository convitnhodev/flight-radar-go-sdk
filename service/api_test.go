package service

import (
	"encoding/json"
	"errors"
	"testing"
)

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

type loginTest struct {
	user           string
	password       string
	expectedStatus string
}

type getZonesTest struct {
	expectedVersion float64
}

type getCountryFlagTest struct {
	country string
	urlFlag string
	err     error
}

var loginTests = []loginTest{
	{"random", "random", "fail"},
	// {"", "", "success"}, // TODO: add user and password to test login
}

var getzonesTests = []getZonesTest{
	{4},
}

var getCountryFlagTests = []getCountryFlagTest{
	{"Korea", "https://www.flightradar24.com/static/images/data/flags-small/korea.svg", nil},
	{"China", "https://www.flightradar24.com/static/images/data/flags-small/china.svg", nil},
	{"xxx", "", errors.New("invalid country flag")},
}

var getAirport = []string{
	"BQL",
	"BYK",
	"UAR",
	"BLD",
	"WBU",
	"KBTF",
}

var getFlight = []string{
	"30c38acf",
	"30c3b690",
	"30c41921",
	"30c4355f",
	"306cca05",
	"30c13531",
}

var test_api = NewFlightRadar24API()

func TestLogin(t *testing.T) {
	for _, test := range loginTests {
		if content, err := test_api.Login(test.user, test.password); err != nil {
			if status, ok := content["status"]; !ok || status != test.expectedStatus {
				t.Errorf("Login(%s, %s) = %s; expected %s (%s)", test.user, test.password, content["status"], test.expectedStatus, content["message"])
			}
		} else {
			if content["status"] != test.expectedStatus {
				t.Errorf("Login(%s, %s) = %s; expected %s (%s)", test.user, test.password, content["status"], test.expectedStatus, content["message"])
			}
		}
	}
}

func TestGetAirlines(t *testing.T) {
	if content, err := test_api.GetAirlines(); err != nil {
		t.Errorf("GetAirlines() = %s", err.Error())
	} else {
		if len(content) == 0 {
			t.Errorf("GetAirlines() = %s", "empty")
		}
		t.Logf("GetAirlines() = %v", prettyPrint(content)) // TOD
	}
}

func TestGetAirports(t *testing.T) {
	if content, err := test_api.GetAirports(); err != nil {
		t.Errorf("GetAirports() = %s", err.Error())
	} else {
		if len(content) == 0 {
			t.Errorf("GetAirports() = %s", "empty")
		}
		t.Logf("GetAirports() = %v", prettyPrint(content)) // TOD
	}
}

func TestGetAirport(t *testing.T) {
	for _, test := range getAirport {
		if content, err := test_api.GetAirport(test); err != nil {
			t.Errorf("GetAirport() = %s", err.Error())
		} else {
			if len(content) == 0 {
				t.Errorf("GetAirport() = %s", "empty")
			}
			t.Logf("GetAirport() = %v", prettyPrint(content)) // TOD
		}
	}

}

func TestGetZones(t *testing.T) {
	for _, test := range getzonesTests {
		if content, err := test_api.GetZones(); err != nil {
			if status, ok := content["version"]; !ok && status != test.expectedVersion {
				t.Errorf("Getzones version = %v; expected %v ", content["version"], test.expectedVersion)
			}
		} else {
			if content["version"] != test.expectedVersion {
				t.Errorf("Getzones version = %v; expected %v ", content["version"], test.expectedVersion)
			}
		}
	}
}

func TestGetCountryFlag(t *testing.T) {
	for _, test := range getCountryFlagTests {
		if result, err := test_api.GetCountryFlag(test.country); err != nil {
			if err.Error() != test.err.Error() {
				t.Errorf("GetCountryFlag return error: %v; expected %v", err, test.err)
			}
		} else {
			if result != test.urlFlag {
				t.Errorf("GetCountryFlag return value: %v, expected %v", result, test.urlFlag)
			}
		}
	}
}

func TestGetFlights(t *testing.T) {
	if content, err := test_api.GetFlights(nil, nil, nil, nil); err != nil {
		t.Errorf("GetFlights() = %s", err.Error())
	} else {
		if len(content) == 0 {
			t.Errorf("GetFlights() = %s", "empty")
		}
		t.Logf("GetFlights() = %v", prettyPrint(content)) // TODO: remove
	}
}

func TestGetFlight(t *testing.T) {
	for _, test := range getFlight {
		if content, err := test_api.GetDetailFlight(test); err != nil {
			t.Errorf("GetDetailFlight() = %s", err.Error())
		} else {
			if content == nil {
				t.Errorf("GetDetailFlight() = %s", "empty")
			}
			t.Logf("GetDetailFlight() = %v", prettyPrint(content)) // TOD
		}
	}
}
