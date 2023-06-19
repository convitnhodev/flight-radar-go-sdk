package _package

import (
	"fmt"
	"golang.org/x/net/html"
	"strconv"
	"time"
)

func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func ConvertToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	default:
		return fmt.Sprintf("%v", v)
	}
}

func ConvertStringToTime(timeString string) (*time.Time, error) {
	layout := "02 Jan 2006"
	t, err := time.Parse(layout, timeString)
	if err != nil {
		fmt.Println("Error when connvert time:", err)
		return nil, err
	}
	return &t, nil
}

func ConvertTimeToTimeWithAmAndPm(timeString string, date time.Time) (*time.Time, error) {
	dateString := date.Format("2006-01-02")
	layout := "03:04 PM"
	fullTimeString := dateString + " " + timeString

	t, err := time.Parse(layout, fullTimeString)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func extractText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}

	var result string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result += extractText(c)
	}

	return result
}
