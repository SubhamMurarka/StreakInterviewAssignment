package main

import (
	"fmt"
	"log"
	"sync"
)

// defining a job alias of function
type Job func()

// pool struct representing the pool/ thread pool
type Pool struct {
	wg         sync.WaitGroup
	job        chan Job
	numWorkers int
}

// creating NewPool
func NewPool(workers int) *Pool {
	pool := &Pool{
		job:        make(chan Job),
		numWorkers: workers,
	}

	pool.wg.Add(workers)

	//preparing Workers for doing any job
	for i := 0; i < workers; i++ {
		go func() {
			defer pool.wg.Done()
			//continuously fetching the job from the channel
			for jb := range pool.job {
				log.Printf("Worker ID : %d\n", i)
				jb()
			}
		}()
	}

	return pool
}

// adding the job to channel
func (p *Pool) AddJob(job Job) {
	p.job <- job
}

// closing the channel to avoid blocking
func (p *Pool) Close() {
	close(p.job)
	p.wg.Wait()
}

func main() {
	pool := NewPool(5)

	/**
	calculations :-
	start = 1      => i = 1      -> 200000
	start = 200001 => i = 200001 -> 400000
	start = 400001 => i = 400001 -> 600000
	start = 600001 => i = 600001 -> 800000
	start = 800001 => i = 800001 -> 1000000
	**/

	// preparing jobs and pushing to channel
	for start := 1; start <= 800001; start += 200000 {
		job := func() {
			for i := start; i <= start+199999; i++ {
				fmt.Printf("Count : %d", i)
			}
		}

		pool.AddJob(job)
	}

	pool.Close()
}
