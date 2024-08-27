package main

import (
	"fmt"
	"mathlib/mathlib"
)

func main() {
	n := 10
	d, s := mathlib.DoubleSquare(n)
	fmt.Println("Double of", n, "is", d)
	fmt.Println("Square of", n, "is", s)

	//
	a := 10
	b := 20
	min, max := mathlib.MinMax(a, b)
	fmt.Println("min:", min)
	fmt.Println("max:", max)

	// Anonymous function
	sum := func(x, y int) int {
		return x + y
	}
	fmt.Println("sum of", n, "and", n, "is", sum(n, n))

	// Callback
	callback := func(s string) {
		fmt.Println("Callback called with:", s)
	}
	doSomething("Hello, World!", callback)

	// Defer
	defer fmt.Println("Goodbye!") // Print a message when the function exits

	fmt.Println("Hello!")

}

func doSomething(s string, callback func(string)) {
	fmt.Println("Doing something...")
	callback(s)
	fmt.Println("Done!")
}
