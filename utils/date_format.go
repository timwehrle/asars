package utils

import (
	"time"

	"github.com/fatih/color"
)

func FormatDate(date string) string {
	if date == "" {
		return color.New(color.Faint).Sprint("None")
	}

	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return color.New(color.FgRed).Sprint("Invalid Date")
	}

	today := time.Now()
	tomorrow := today.Add(24 * time.Hour)
	weekLater := today.Add(6 * 24 * time.Hour)

	green := color.New(color.FgGreen)
	blue := color.New(color.FgBlue)
	red := color.New(color.FgRed)
	white := color.New(color.FgWhite)

	if parsedDate.Equal(today.Truncate(24 * time.Hour)) {
		return green.Sprint("Today")
	}

	if parsedDate.Equal(tomorrow.Truncate(24 * time.Hour)) {
		return green.Sprint("Tomorrow")
	}

	if parsedDate.After(tomorrow) && parsedDate.Before(weekLater) {
		return blue.Sprint(parsedDate.Format("Mon"))
	}

	if parsedDate.Before(today) {
		return red.Sprint(parsedDate.Format("Jan 02, 2006"))
	}

	return white.Sprint(parsedDate.Format("Jan 02, 2006"))
}
