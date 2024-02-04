package utils

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func ConvertStrToDateTime(date string) time.Time {
	now := time.Now()
	if date != "" {
		log.Printf("supplied time >>>> %s", date)
		timeString := date
		theTime, err := time.Parse("2006-01-02 03:04:05", timeString)

		if err != nil {
			fmt.Println("Could not parse datetime:", err)
		}
		log.Printf("TIME >>> %s", theTime)
		return theTime
	} else {
		return now
	}
}

func ConvertEnToIDDateTime(date time.Time) string {
	locale := strings.NewReplacer(
		"January", "Januari",
		"February", "Febuari",
		"March", "Maret",
		"April", "April",
		"May", "Mei",
		"June", "Juni",
		"July", "Juli",
		"August", "Agustus",
		"September", "September",
		"October", "Oktober",
		"November", "November",
		"December", "Desember")

	currentDate := date.Format("2 January 2006")
	outputString := locale.Replace(currentDate)
	return outputString
}
