package util

import (
	"time"
)

/*
GivenTimeIsBetweenOnlyTimeRange

// Check t is between start and end time.
*/
func GivenTimeIsBetweenOnlyTimeRange(t, start, end time.Time) bool {
	return !t.Before(start) && !t.After(end)
}
