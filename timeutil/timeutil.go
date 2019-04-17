package timeutil

import "time"

// GetTimeNow returns the current system time
func GetTimeNow() time.Time {
	return timeNow()
}

// GetTimeNowUTC returns the UTC representation of the current system time
func GetTimeNowUTC() time.Time {
	return timeNow().UTC()
}

// FormatDate returns the time in string format "yyyy-MM-dd"
func FormatDate(value time.Time) string {
	return value.Format("2006-01-02")
}

// FormatTime returns the time in string format "HH:mm:ss"
func FormatTime(value time.Time) string {
	return value.Format("15:04:05")
}

// FormatDateTime returns the time in string format "yyyy-MM-ddTHH:mm:ss"
func FormatDateTime(value time.Time) string {
	return value.Format("2006-01-02T15:04:05")
}
