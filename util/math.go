package util

func SumInt(values ...int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}

func SumFloat(values ...float64) float64 {
	sum := float64(0)
	for _, v := range values {
		sum += v
	}
	return sum
}
