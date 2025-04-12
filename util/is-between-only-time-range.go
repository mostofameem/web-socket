package util

import (
	"time"
)

// IsBetweenOnlyTimeRange  Checks given is between start and end time
//
// given , start , end must be in UTC time
//
// given , start , end in dateTime
func IsBetweenOnlyTimeRange(given, start, end time.Time) bool {
	givenTime := convertToTimeOnly(given)
	startTime := convertToTimeOnly(start)
	endTime := convertToTimeOnly(end)
	return !givenTime.Before(startTime) && !givenTime.After(endTime)
}

// convertToTimeOnly  convert t in only time
//
// t in dateTime
func convertToTimeOnly(t time.Time) (onlyTime time.Time) {
	onlyTime = time.Date(0, 1, 1, t.Hour(), t.Minute(), t.Second(), 0, time.UTC)
	return
}
