package util

import "math"

func CalNumOfPages(total, pageSize float64) int {
	if pageSize <= 0 {
		return 0
	}

	return int(math.Ceil(total / pageSize))
}
