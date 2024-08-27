package main

import (
	"fmt"
	"slices"
)

type AnotherInt int

type AllInts interface {
	~int
}

type Number interface {
	int64 | float64
}

type Numeric interface {
	int | int8 | int16 | int32 | int64 | float64
}

func main() {
	fmt.Println("4 + 3 = ", add(4, 3))
	s := []AnotherInt{10, 1, 2}
	fmt.Println(AddElements(s))

	//Clone 1 array
	s1 := []int{1, 2, -1, -2}
	s2 := slices.Clone(s1)
	s3 := slices.Clone(s1[2:])
	fmt.Println(s1[2], s2[2], s3[0])
	s1[2] = 0
	s1[3] = 0
	fmt.Println(s1[2], s2[2], s3[0])

	s1 = slices.Compact(s1)
	fmt.Println("s1 (compact):", s1)
	fmt.Println(slices.Contains(s1, 2), slices.Contains(s1, -2))

	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  10,
		"second": 20,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  10.0,
		"second": 20.0,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))

}

func add[T Numeric](a, b T) T {
	return a + b
}

func AddElements[T AllInts](s []T) T {
	sum := T(0)
	for _, v := range s {
		sum = sum + v
	}
	return sum
}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// SumIntsOrFloats sums the values of map m. It supports both floats and integers
// as map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
