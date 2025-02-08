package main

import (
	"context"
	"log"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

// implementation 1 implementing semaphore from scratch
func Api(i int, sem chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	// blocking a slot in semaphore/(buffered channel)
	sem <- 1
	log.Printf("Api request served : %d\n", i)
	// releasing a slot in semaphore/(buffered channel)
	<-sem
}

// implementation 2 using semaphores from sync
func Api2(i int, sem *semaphore.Weighted, wg *sync.WaitGroup) {
	defer wg.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// blocking a slot in semaphore/(buffered channel)
	if err := sem.Acquire(ctx, 1); err != nil {
		log.Printf("Api request Failed : %d\n", i)
		return
	}

	// can use TryAcquire as well if slot free
	// acquire else return true or false repectively.

	log.Printf("Api request served : %d\n", i)

	// releasing a slot in semaphore/(buffered channel)
	sem.Release(1)
}

// suppose only 3 requests can be send to api concurrently
func main() {
	var wg sync.WaitGroup

	// implementation -> 1 , creating a semaphore == buffered channel
	semaphore1 := make(chan int, 3)
	//implementation -> 2, using sync package
	semaphore2 := semaphore.NewWeighted(3)

	//let total requests be 10 but 3 at a time is served only
	for i := 1; i <= 5; i++ {
		wg.Add(2)
		go Api(i, semaphore1, &wg)
		go Api2(i*10, semaphore2, &wg)
	}

	wg.Wait()
}
