package main

import (
	"fmt"
	"sync"
)

func produceData(out chan<- int, wg *sync.WaitGroup, data ...int) {
	defer wg.Done()
	for _, d := range data {
		out <- d
	}
}

func collectData(in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range in {
		fmt.Println("Received:", val)
	}
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	out := make(chan int)
	var wg sync.WaitGroup

	// Number of producers
	wg.Add(2)

	// Start producer goroutines
	go produceData(out, &wg, data[:len(data)/2]...)
	go produceData(out, &wg, data[len(data)/2:]...)

	// Start collector goroutine
	wg.Add(1)
	go collectData(out, &wg)

	// Wait for all goroutines to finish
	wg.Wait()

	// Close the channel after all data is produced
	close(out)
}
