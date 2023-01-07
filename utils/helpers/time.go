package helpers

import (
	"time"
)

func ParseDateTime(dateTime time.Time) string {
	return dateTime.Format("2006-01-02T15:04:05")
}
func ParseTimeFromStringToTime(timeInString string) (time.Time, error) {
	timeInFormat, err := time.Parse("2006-01-02T15:04:05Z", timeInString)
	return timeInFormat, err
}
