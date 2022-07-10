package main

import (
	"log"
	"runtime"
	"sync"
	"time"
)

// Pipelines usually have same parameter and output to maintain compositability of the functions, e.g.  the output of square function can be a parameter to another square function.

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	var sequenceDuration time.Duration
	var concurrentDuration time.Duration

	input := []int{}

	for i := 1; i < 1001; i++ {
		input = append(input, i)
	}

	//CONCURRENT RUNNING
	wg.Add(1)
	go func(inputs ...int) {
		defer wg.Done()
		start := time.Now()

		for range square(generator(inputs...)) {
			// log.Println("Received Concurrent Output", msg)
		}
		duration := time.Since(start)
		log.Println("Concurrent Squaring (Pipeline) Done in", duration)
		concurrentDuration = duration
	}(input...)

	//SEQUENCE RUNNING
	wg.Add(1)
	go func(inputs ...int) {
		defer wg.Done()
		start := time.Now()
		for _, num := range inputs {
			/*result :=*/ squareSequential(num)
			// log.Println("Received Sequence Output", result)
		}
		duration := time.Since(start)
		log.Println("Sequence Squaring Done in", duration)
		sequenceDuration = duration
	}(input...)

	wg.Wait()
	//RESULT

	if sequenceDuration > concurrentDuration {
		log.Println("RESULT: Concurrent Approach is faster")
		return
	}
	log.Println("RESULT: Sequence Approach is faster")
}

// generator() -> square() -> print

// generator - convertes a list of integers to a channel
func generator(nums ...int) <-chan int {

	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out

}

// square - receive on inbound channel
// square the number
// output on outbound channel
func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for msg := range in {
			out <- msg * msg
		}
		close(out)
	}()

	return out
}

func squareSequential(input int) int {
	return input * input
}
