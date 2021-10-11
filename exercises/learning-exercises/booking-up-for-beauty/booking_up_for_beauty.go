package booking

import (
	"fmt"
	"time"
)

const (
	ScheduleLayout    = "1/2/2006 15:04:05"
	HasPassedLayout   = "January 2, 2006 15:04:05"
	IsAfternoonLayout = "Monday, January 2, 2006 15:04:05"
	DescriptionLayout = "Monday, January 2, 2006, at 15:04."
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	time, err := time.Parse(ScheduleLayout, date)
	if err != nil {
		panic(err)
	}
	return time
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	parsedDate, err := time.Parse(HasPassedLayout, date)
	if err != nil {
		panic(err)
	}
	return time.Now().After(parsedDate)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	parsedDate, err := time.Parse(IsAfternoonLayout, date)
	if err != nil {
		panic(err)
	}
	hour, _, _ := parsedDate.Clock()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	return fmt.Sprintf("You have an appointment on %s", Schedule(date).Format(DescriptionLayout))
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
