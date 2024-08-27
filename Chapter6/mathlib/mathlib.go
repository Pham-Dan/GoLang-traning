package mathlib

// Function return 2 value
func DoubleSquare(x int) (int, int) {
	return x * 2, x * x
}

// Sorting from smaller to bigger value
func SortTwo(x, y int) (int, int) {
	if x > y {
		return y, x
	}
	return x, y
}

func MinMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
		return min, max
	}
	min = x
	max = y
	return
}
