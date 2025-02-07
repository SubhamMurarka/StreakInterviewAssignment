package main

import "fmt"

/**
 stage 1 -> passing all numbers of array/slice to channel A
 stage 2 -> adding 2 to all numbers from channel A and pass to channel B
 stage 3 -> Multiply 2 to all numbers from channel B and pass to channel C
 Finally priniting all numbers from Channel c in main go routine
**/

// Stage 1
func Stage1(num []int) <-chan int {
	ChannelA := make(chan int)
	go func() {
		defer close(ChannelA)
		for i := range num {
			ChannelA <- num[i]
		}
	}()
	return ChannelA
}

// Stage 2
func Stage2(ChannelA <-chan int) <-chan int {
	ChannelB := make(chan int)
	go func() {
		defer close(ChannelB)
		for i := range ChannelA {
			ChannelB <- i + 2
		}
	}()
	return ChannelB
}

// Stage 3
func Stage3(ChannelB <-chan int) <-chan int {
	ChannelC := make(chan int)
	go func() {
		defer close(ChannelC)
		for i := range ChannelB {
			ChannelC <- i * 2
		}
	}()
	return ChannelC
}

// Stage 1 -> ChannelA -> Stage 2 -> ChannelB -> Stage3 -> ChannelC -> main Go Routine
// it implies that
// main go routine waits for Stage 3 go routine => Stage 3 go routine waits for Stage 2 go routine
// Stage 2 go routine waits for Stage 1 go routine until writing channel is closed.
// this is because reading/ recieving from channel is a blocking call.
// Ensuring no pre-exit of main go routine and avoiding force killing of other go routines.

func main() {
	//stage 1
	num := []int{1, 2, 3, 4, 5}
	ChannelA := Stage1(num)
	ChannelB := Stage2(ChannelA)
	ChannelC := Stage3(ChannelB)

	for i := range ChannelC {
		fmt.Println(i)
	}
}
