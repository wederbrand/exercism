package booking

import "time"

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	parse, _ := time.Parse("1/02/2006 15:04:05", date)
	return parse
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	parse, _ := time.Parse("January 2, 2006 15:04:05", date)
	return parse.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	parse, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	hour := parse.Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	parse, _ := time.Parse("1/2/2006 15:04:05", date)
	return parse.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	now := time.Now()
	return time.Date(now.Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
