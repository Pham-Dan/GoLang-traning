package main

import (
	"fmt"
	"time"
)

var Company = "Hopee"

func main() {
	
	// Data type
	var name string
	var age int = 26
	var isMale bool = true

   // Get User Input
	fmt.Printf("Please give me your name: ")
   fmt.Scanln(&name)

	// constant
	const PI = 3.14

	// Array
	programmingsLanguage := [...]string{"PHP", "HTML", "CSS", "JS"}

	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
	fmt.Println("Male: ", isMale)
	fmt.Println("Programmings Language: ", programmingsLanguage)

	// Switch case
	day := 1

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	}

	// Loop
	loop(1, 100)
	for _, v := range programmingsLanguage {
		fmt.Println(v)
	}

   // Concurrency
   printNumbers()

}

func loop(start int, end int) {
	for i := start; i <= end; i++ {
		fmt.Println(i)
	}
}

func printNumbers() {
   for i := 1; i <= 5; i++ {
       fmt.Println(i);
       time.Sleep(time.Second);
   }
}
