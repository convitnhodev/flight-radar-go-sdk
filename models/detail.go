package models

import "time"

type CoreFlight struct {
	Date                     time.Time
	From                     string
	To                       string
	AirCraft                 string
	FlightTime               string
	ScheduledTimeOfDeparture time.Time
	ActualTimeOfDeparture    time.Time
	ScheduledTimeOfArrival   time.Time
	Status                   string
}
