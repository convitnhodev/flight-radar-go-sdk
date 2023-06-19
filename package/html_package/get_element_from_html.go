package html_package

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/convitnhodev/flight-radar-go-sdk/models"
	"strconv"
	"strings"
	"time"
)

type ConvertFlight struct {
	DateTime                       string
	DateOffset                     string
	From                           string
	To                             string
	AirCraft                       string
	FlightTime                     string
	ScheduledTimeOfDepartureTime   string
	ScheduledTimeOfDepartureOffset string
	ActualTimeOfDepartureTime      string
	ActualTimeOfDepartureOffset    string
	ScheduledTimeOfArrivalTime     string
	ScheduledTimeOfArrivalOffset   string
	Status                         string
}

func GetHTML(htmlString string) []ConvertFlight {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlString))
	if err != nil {
		return nil
	}

	arrayResult := make([]ConvertFlight, 0)

	source := `section#cnt-data-content table tbody tr`
	trr := doc.Find(source)
	size := trr.Length()
	for i := 0; i < size; i++ {
		var tmp string
		var ok bool
		var result ConvertFlight
		// subSource := source + " td"
		tr := trr.Eq(i)
		tr = tr.Find("td")
		tmp, ok = tr.Eq(2).Attr("data-timestamp")
		if !ok {
			result.DateTime = ""
		} else {
			result.DateTime = tmp
		}

		tmp, ok = tr.Eq(2).Attr("data-offset")
		if !ok {
			result.DateOffset = ""
		} else {
			result.DateOffset = tmp
		}

		result.From = tr.Eq(3).Find("a").Text()
		result.To = tr.Eq(4).Find("a").Text()
		result.AirCraft = tr.Eq(5).Text()
		result.FlightTime = tr.Eq(6).Text()
		tmp, ok = tr.Eq(7).Attr("data-timestamp")
		if !ok {
			result.ScheduledTimeOfDepartureTime = ""
		} else {
			result.ScheduledTimeOfDepartureTime = tmp
		}

		tmp, ok = tr.Eq(7).Attr("data-offset")
		if !ok {
			result.ScheduledTimeOfDepartureOffset = ""
		} else {
			result.ScheduledTimeOfDepartureOffset = tmp
		}

		tmp, ok = tr.Eq(8).Attr("data-timestamp")
		if !ok {
			result.ActualTimeOfDepartureTime = ""
		} else {
			result.ActualTimeOfDepartureTime = tmp
		}

		tmp, ok = tr.Eq(8).Attr("data-offset")
		if !ok {
			result.ActualTimeOfDepartureOffset = ""
		} else {
			result.ActualTimeOfDepartureOffset = tmp
		}

		tmp, ok = tr.Eq(9).Attr("data-timestamp")
		if !ok {
			result.ScheduledTimeOfArrivalTime = ""
		} else {
			result.ScheduledTimeOfArrivalTime = tmp
		}

		tmp, ok = tr.Eq(9).Attr("data-offset")
		if !ok {
			result.ScheduledTimeOfArrivalOffset = ""
		} else {
			result.ScheduledTimeOfArrivalOffset = tmp
		}

		tmp, ok = tr.Eq(11).Attr("data-prefix")
		if !ok {
			result.Status = ""
		} else {
			result.Status = tmp
		}

		arrayResult = append(arrayResult, result)

	}

	return arrayResult
}

func ConvertArrayRawModelToModel(data []ConvertFlight) []models.CoreFlight {
	results := make([]models.CoreFlight, 0)
	for _, v := range data {
		result := ConvertRawModelToModel(v)
		results = append(results, result)
	}

	return results
}

func ConvertRawModelToModel(data ConvertFlight) models.CoreFlight {
	datetimeInt, _ := strconv.ParseInt(data.DateTime, 10, 64)
	dateOffsetInt, _ := strconv.ParseInt(data.DateOffset, 10, 64)
	resultDate := time.Unix(int64(datetimeInt)-int64(dateOffsetInt), 0)

	scheduledTimeOfDepartureTimeInt, _ := strconv.ParseInt(data.ScheduledTimeOfDepartureTime, 10, 64)
	ScheduledTimeOfDepartureOffsetInt, _ := strconv.ParseInt(data.ScheduledTimeOfDepartureOffset, 10, 64)
	resultScheduledTimeOfDeparture := time.Unix(int64(scheduledTimeOfDepartureTimeInt)-int64(ScheduledTimeOfDepartureOffsetInt), 0)

	actualTimeOfDepartureTimeInt, _ := strconv.ParseInt(data.ActualTimeOfDepartureTime, 10, 64)
	actualTimeOfDepartureOffsetInt, _ := strconv.ParseInt(data.ActualTimeOfDepartureOffset, 10, 64)
	resultActualTimeOfDeparture := time.Unix(int64(actualTimeOfDepartureTimeInt)-int64(actualTimeOfDepartureOffsetInt), 0)

	scheduledTimeOfArrivalTimeInt, _ := strconv.ParseInt(data.ScheduledTimeOfArrivalTime, 10, 64)
	scheduledTimeOfArrivalTimeOffsetInt, _ := strconv.ParseInt(data.ScheduledTimeOfArrivalOffset, 10, 64)
	resultScheduledTimeOfArrival := time.Unix(int64(scheduledTimeOfArrivalTimeInt)-int64(scheduledTimeOfArrivalTimeOffsetInt), 0)

	if data.From != "" {
		data.From = strings.Trim(data.From, "()")
	}

	if data.To != "" {
		data.To = strings.Trim(data.To, "()")
	}

	return models.CoreFlight{
		Date:                     resultDate,
		From:                     data.From,
		To:                       data.To,
		AirCraft:                 data.AirCraft,
		FlightTime:               data.FlightTime,
		ScheduledTimeOfDeparture: resultScheduledTimeOfDeparture,
		ActualTimeOfDeparture:    resultActualTimeOfDeparture,
		ScheduledTimeOfArrival:   resultScheduledTimeOfArrival,
		Status:                   data.Status,
	}

}
