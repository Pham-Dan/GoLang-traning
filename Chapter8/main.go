package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	// Go routine
	go func(x int) {
		fmt.Printf("%d ", x)
	}(10)

	time.Sleep(time.Second)
	fmt.Println("Exiting...")

	for i := 0; i < 10; i++ {
		go func(x int) {
			fmt.Printf("%d ", x)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("\nExiting...")

	//waitGroup
	var waitGroup sync.WaitGroup
	for i := 0; i < 5; i++ {
		waitGroup.Add(1)
		go func(x int) {
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}
	waitGroup.Wait()

	// Channel
	waitGroup.Add(1)
	channel := make(chan int, 1)
	go func(c chan int) {
		defer waitGroup.Done()
		writeToChannel(c, 10)
		fmt.Println("Exit.")
	}(channel)
	fmt.Println("Read:", <-channel)

	urls := []string{
		"https://www.google.com",
		"https://www.example.com",
		"https://www.stackoverflow.com",
		"https://www.github.com",
	}

	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		go func(url string) {
			defer wg.Done()
			fetchUrl(url)
		}(url)
	}

	wg.Wait()
	fmt.Println("All URLs fetched!")
}

func writeToChannel(c chan int, x int) {
	c <- x
	close(c)
}

func printer(ch chan bool) {
	ch <- true
}

func fetchUrl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Fetched %s\n", url)
}
