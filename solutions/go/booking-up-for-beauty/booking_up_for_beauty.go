package booking

import (
    "time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    layout := "1/2/2006 15:04:05"
	t, err := time.Parse(layout,date)

    if err != nil {
		panic("Parse failed")
	}

    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
    t, err := time.Parse(layout,date)

	if err != nil {
		panic("Parse failed")
	}

	if time.Now().After(t) {
		return true
	} else {
		return false
	}
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
    t, err := time.Parse(layout,date)

	if err != nil {
		panic("Parse failed")
	}

    if 12 <= t.Hour() && t.Hour() <= 18 {
        return true
    } else {
    	return false
    }
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	t, err := time.Parse(layout, date)

	if err != nil {
		panic("Parse failed")
	}

	return t.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	currentYear := time.Now().UTC().Year()
	return time.Date(currentYear, time.September, 15, 0, 0, 0, 0, time.UTC)
}
