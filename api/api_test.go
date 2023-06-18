package api

import (
	"testing"
)

type loginTest struct {
	user           string
	password       string
	expectedStatus string
}

type getZonesTest struct {
	expectedVersion float64
}

var loginTests = []loginTest{
	{"random", "random", "fail"},
	// {"", "", "success"}, // TODO: add user and password to test login
}

var getzonesTests = []getZonesTest{
	{4},
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
