package util

func Min(values ...int) int {
	if len(values) == 0 {
		return 0
	}

	min := values[0]
	for _, value := range values[1:] {
		if value < min {
			min = value
		}
	}
	return min
}

func Max(values ...int) int {
	if len(values) == 0 {
		return 0
	}

	max := values[0]
	for _, value := range values[1:] {
		if value > max {
			max = value
		}
	}
	return max
}

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
