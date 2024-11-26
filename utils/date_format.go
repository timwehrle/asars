package utils

import (
	"time"
)

func FormatDate(date string) string {
	if date == "" {
		return Faint.Sprint("None")
	}

	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return Red.Sprint("Invalid Date")
	}

	today := time.Now()
	tomorrow := today.Add(24 * time.Hour)
	weekLater := today.Add(6 * 24 * time.Hour)

	if parsedDate.Equal(today.Truncate(24 * time.Hour)) {
		return Green.Sprint("Today")
	}

	if parsedDate.Equal(tomorrow.Truncate(24 * time.Hour)) {
		return Green.Sprint("Tomorrow")
	}

	if parsedDate.After(tomorrow) && parsedDate.Before(weekLater) {
		return Blue.Sprint(parsedDate.Format("Mon"))
	}

	if parsedDate.Before(today) {
		return Red.Sprint(parsedDate.Format("Jan 02, 2006"))
	}

	return White.Sprint(parsedDate.Format("Jan 02, 2006"))
}
