package utils

import (
	"fmt"
	"time"
)

func ConvertStrToDateTime(date string) time.Time {
	now := time.Now()
	if date != "" {
		timeString := date
		theTime, err := time.Parse("2006-01-02 03:04:05", timeString)

		if err != nil {
			fmt.Println("Could not parse datetime:", err)
		}
		return theTime
	}
	return now
}
