package main

import (
	"fmt"
	"sync"
)

// work to be done alias for func here
type Job func()

var wg sync.WaitGroup

// go routines, number of go routines = number of jobs
func worker(job Job) {
	defer wg.Done()
	job()
}

// FanOut concurrency pattern
func main() {
	// outer loop will run for 5 times -> no. of jobs
	for start := 1; start <= 800001; start += 200000 {
		// preparing jobs and to be passed to respective go routines
		job := func() {
			for i := start; i <= start+199999; i++ {
				fmt.Println("Count : ", i)
			}
		}

		wg.Add(1)
		// spawning a go routine for each job
		go worker(job)
	}

	/**
	calculations :-
	start = 1      => i = 1      -> 200000
	start = 200001 => i = 200001 -> 400000
	start = 400001 => i = 400001 -> 600000
	start = 600001 => i = 600001 -> 800000
	start = 800001 => i = 800001 -> 1000000
	**/

	wg.Wait()
}
