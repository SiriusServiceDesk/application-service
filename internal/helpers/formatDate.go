package helpers

import "time"

func FormatDate(d time.Time) string {
	return d.Format("02.01.2006")
}

func FormatDateWithTime(d time.Time) string {
	return d.Format("02.01.2006 15:04")
}
