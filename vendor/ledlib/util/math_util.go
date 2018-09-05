package util

func RoundToInt(input float64) int {
	return int(input + 0.5)
}

func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
