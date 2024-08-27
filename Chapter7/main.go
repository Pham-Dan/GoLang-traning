package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type UseAll struct {
	Name    string `json:"username"`
	Surname string `json:"surname"`
	Year    int    `json:"created"`
}

func main() {
	ReadFile("text.txt")
	WriteFile("new.txt")

	// Json
	useall := UseAll{Name: "Dan", Surname: "Pham", Year: 2023}
	// Encoding JSON data: Convert Structure to JSON record with fields
	t, err := json.Marshal(&useall)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Value %s\n", t)
	}

	jsonRecord := []byte(t)
	// Create a structure variable to store the result
	temp := UseAll{}
	err = json.Unmarshal(jsonRecord, &temp)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Data type: %T with value %v\n", temp, temp)
	}

}

func ReadFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(string(buf[:n]))
	}
}

func WriteFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := []byte("Hello, world!")
	n, err := file.Write(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Wrote", n, "bytes")
}
