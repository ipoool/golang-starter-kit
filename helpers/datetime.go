package helpers

import (
	"fmt"
	"strings"
	"time"
)

const (
	dateTimeLayout       string = "2006-01-02 15:04:05"
	timestampLayout      string = "20060102150405"
	dateTimeLayoutHeader string = "Mon, _2 Jan 2006 15:04:05 MST"
	timeLayout           string = "15:04:05"
	dateLayout           string = "2006-01-02"
)

// IndonesianMount - define default Indonesian month name
var IndonesianMount = map[int]string{
	1:  "Januari",
	2:  "Februari",
	3:  "Maret",
	4:  "April",
	5:  "Mei",
	6:  "Juni",
	7:  "Juli",
	8:  "Agustus",
	9:  "September",
	10: "Oktober",
	11: "November",
	12: "Desember",
}

func timeLoadLocation() *time.Location {
	time, _ := time.LoadLocation("Asia/Jakarta")
	return time
}

// GetCurrentDateTime - get current date and time
func GetCurrentDateTime() time.Time {
	currentDateTime := time.Now().In(timeLoadLocation())
	return currentDateTime
}

// GetGrabDatetime - Get current date and time with RFC7231 layout as grab requirement (output: Wed, 10 Jan 2018 04:54:22 GMT)
func GetGrabDatetime() string {
	var rfc7231Layout = "Mon, 02 Jan 2006 15:04:05 GMT"
	currentDateTime := time.Now().UTC()
	return currentDateTime.Format(rfc7231Layout)
}

// GetCurrentDate - get today date only
func GetCurrentDate() string {
	currentDate := time.Now().Local()
	return currentDate.Format(dateLayout)
}

// GetCurrentTime - get today date only
func GetCurrentTime() time.Time {
	currentTime := time.Now().Local()
	return currentTime
}

// TimeToString - convert from time.Time into string
func TimeToString(dateInput time.Time) string {
	return dateInput.Format(dateTimeLayout)
}

// StringToTime - convert from date string into time.Time
func StringToTime(dateInput string) (time.Time, error) {
	result, err := time.ParseInLocation(dateTimeLayout, dateInput, timeLoadLocation())
	return result, err
}

// GetCurrentTimestamp - get current timestamp (string) with layout
func GetCurrentTimestamp() string {
	t := time.Now()
	return t.Format(timestampLayout)
}

// DateTimeClearFromDB - Clear the datetime fro DB
func DateTimeClearFromDB(dateTimeInput string) (time.Time, error) {
	dateTimeClear := strings.Replace(dateTimeInput, "T", " ", -1)
	dateTimeString := strings.Join(strings.Split(dateTimeClear, "+07:00"), " ")
	dateTimeString = strings.Replace(dateTimeString, "Z", "", -1)
	dateTime, err := StringToTime(strings.Trim(dateTimeString, " "))

	if err != nil {
		fmt.Println(err.Error())
		return GetCurrentDateTime(), err
	}

	return dateTime, nil
}
