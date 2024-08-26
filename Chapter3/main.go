package main

import (
	"fmt"
	"regexp"
)

// Struct
type Person struct {
	name    string
	age     int
	address string
}

func main() {
	// Map

	mapData := map[string]string{
		"name":    "Dan",
		"age":     "Hai muoi",
		"address": "Hue",
	}
	fmt.Println(mapData)

	//
	for key, value := range mapData {
		fmt.Println("Key:", key, "value: ", value)
	}

	person := Person{name: "Dan", age: 20, address: "Hue"}
	fmt.Println(person)

	//Regex
	fmt.Println(matchNameSur("Dan"))

	//
	array1 := [5]int{1, 2, 3, 4, 5}
	result1 := arrayToMap(array1)
	fmt.Println(result1)

}

func matchNameSur(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)

	return re.Match(t)
}

// Write a Go program that converts an existing array into a map.
func arrayToMap(arr [5]int) map[int]int {
	result := make(map[int]int)

	for index, value := range arr {
		result[value] = index
	}

	return result
}
