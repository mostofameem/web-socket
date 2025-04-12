package util

func Float64ToIntSlice(floatSlice []float64) []int {
	intSlice := make([]int, len(floatSlice))

	for i, v := range floatSlice {
		intSlice[i] = int(v) // Cast float64 to int
	}

	return intSlice
}
