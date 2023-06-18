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
	{"", "", "success"}, // TODO: add user and password to test login
	{"random", "random", "fail"},
}

var getzonesTests = []getZonesTest{
	{4},
}

func TestLogin(t *testing.T) {
	api := NewFlightRadar24API()

	for _, test := range loginTests {
		if content, err := api.Login(test.user, test.password); err != nil {
			if status, ok := content["status"]; !ok && status != test.expectedStatus {
				t.Errorf("Login(%s, %s) = %s; expected %s (%s)", test.user, test.password, content["status"], test.expectedStatus, content["message"])
			}
		} else {
			if content["status"] != test.expectedStatus {
				t.Errorf("Login(%s, %s) = %s; expected %s (%s)", test.user, test.password, content["status"], test.expectedStatus, content["message"])
			}
		}
	}
}

func TestGetZones(t *testing.T) {
	api := NewFlightRadar24API()

	for _, test := range getzonesTests {
		if content, err := api.GetZones(); err != nil {
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
