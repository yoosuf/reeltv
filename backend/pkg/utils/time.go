package utils

import "time"

// NowUTC returns current time in UTC
func NowUTC() time.Time {
	return time.Now().UTC()
}

// ToUTC converts a time to UTC
func ToUTC(t time.Time) time.Time {
	return t.UTC()
}
