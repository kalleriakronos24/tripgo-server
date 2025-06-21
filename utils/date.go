package utils

import (
	"fmt"
	"time"
)

func ConvertStrToDateTime(date string) time.Time {
	now := time.Now()
	if date != "" {
		timeString := date
		theTime, err := time.Parse(time.RFC3339, timeString)
		if err != nil {
			fmt.Println("Could not parse datetime:", err)
		}
		return theTime
	} else {
		return now
	}
}

func ConvertEnToIDDateTime(date time.Time) string {
	currentDate := date.Format("2 January 2006 09:00 AM")
	return currentDate
}

func CurrentDateTime() string {
	return time.Now().Format(time.RFC3339)
}
