package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Job func()

// Worker function to execute the job
func worker(job Job) {
	defer wg.Done()
	job()
}

func main() {
	result := make(chan int, 10)

	// Fan-Out: Spawning multiple goroutines
	numJobs := 12
	for start := 1; start <= numJobs; start++ {
		wg.Add(1)
		startCopy := start
		job := func() {
			res := startCopy * 2
			result <- res
		}
		go worker(job)
	}

	// Fan-In: Collect results in a separate goroutine to prevent deadlock
	go func() {
		wg.Wait()
		close(result) // Close channel after all writes are done
	}()

	// Read from result channel
	for i := range result {
		fmt.Println(i)
	}
}
