package main

import (
	"fmt"
	"os"
)

func main() {
	/* The error data type */
	fmt.Println(formattedError(0, 0))

	// Create an empty slice
	aSlice := []float64{}
	// Both length and capacity are 0 because aSlice is empty
	fmt.Println(aSlice, len(aSlice), cap(aSlice))
	// Add elements to a slice
	aSlice = append(aSlice, 1234.56)
	aSlice = append(aSlice, -34.0)
	fmt.Println(aSlice, "with length", len(aSlice))

	//
	// Define two arrays with fixed sizes
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{4, 5, 6}
	sliceResult := concatenateArrays(array1, array2)

	fmt.Println(sliceResult)

}

/* The error data type */
func formattedError(a, b int) error {
	if a == 0 && b == 0 {
		return fmt.Errorf("a %d and b %d. UserID: %d", a, b, os.Getuid())
	}
	return nil
}

func concatenateArrays(array1, array2 [3]int) []int {
	// Convert arrays to slices
	slice1 := array1[:]
	slice2 := array2[:]

	// Use the append function to concatenate slice2 to slice1
	result := append(slice1, slice2...)
	return result
}
