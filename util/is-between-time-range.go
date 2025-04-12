package util

import "time"

// IsBetweenTimeRange checks if the given time is between start and end time.
//
// given, start, end must be in UTC time
//
// given, start, end is datetime
func IsBetweenTimeRange(given, start, end time.Time) bool {
	return given.After(start) && given.Before(end)
}
