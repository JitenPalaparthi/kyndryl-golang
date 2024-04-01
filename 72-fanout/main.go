package main

import (
	"fmt"
	"sync"
)

func worker(id int, tasks <-chan int, results chan<- int) {
	for task := range tasks {
		// Simulate processing task
		result := task * 2
		// Send result to results channel
		results <- result
		fmt.Printf("Worker %d processed task %d, result: %d\n", id, task, result)
	}
}

func main() {
	// Number of workers
	numWorkers := 3

	// Create channels for tasks and results
	tasks := make(chan int, 10)
	results := make(chan int, 10)

	// Create a WaitGroup to synchronize goroutines
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		go func(workerID int) {
			defer wg.Done()
			worker(workerID, tasks, results)
		}(i)
	}

	// Generate tasks
	for i := 1; i <= 5; i++ {
		tasks <- i
	}
	close(tasks) // Close tasks channel to signal that no more tasks will be sent

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		// Close the results channel after all workers are done
		close(results)
	}()

	// Collect results from workers
	for result := range results {
		fmt.Println("Received result:", result)
	}
}
