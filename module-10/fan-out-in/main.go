package main

import (
	"log"
	"runtime"
	"sync"
	"time"
)

// Pipelines usually have same parameter and output to maintain compositability of the functions, e.g.  the output of square function can be a parameter to another square function.

//conclusion pipeline can be useful for concurrencies/parallelism but might incur more cpu usage when use with not so cpu-heavy processes

var counter int64 = 0

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	var sequenceDuration time.Duration
	var concurrentDuration time.Duration

	input := []int{}

	for i := 1; i < 100001; i++ {
		input = append(input, i)
	}

	in := generator(input...)

	//CONCURRENT RUNNING
	wg.Add(1)
	go func(inp <-chan int) {
		defer wg.Done()
		start := time.Now()

		//Fan-out
		ch1 := square(inp)
		ch2 := square(inp)
		ch3 := square(inp)
		ch4 := square(inp)

		for range merge(ch1, ch2, ch3, ch4) {
			// log.Println("Received Concurrent Output", msg)
			// counter++
		}
		duration := time.Since(start)
		log.Println("Concurrent Squaring (Fan Out-In) Done in", duration)
		concurrentDuration = duration
	}(in)

	//SEQUENCE RUNNING
	wg.Add(1)
	go func(inp <-chan int) {
		defer wg.Done()
		start := time.Now()
		for msg := range inp {
			/*result := */ squareSequential(msg)
			// log.Println("Received Sequence Output", result)
		}
		duration := time.Since(start)
		log.Println("Sequence Squaring Done in", duration)
		sequenceDuration = duration
	}(in)

	wg.Wait()
	//RESULT

	if sequenceDuration > concurrentDuration {
		log.Println("RESULT: Concurrent Approach is faster")
		return
	}
	log.Println("RESULT: Sequence Approach is faster")
}

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

func merge(cs ...<-chan int) <-chan int {
	//Fan-in
	out := make(chan int) // merge channel
	var wg sync.WaitGroup

	wg.Add(len(cs))
	for _, ch := range cs {
		go func(chIn <-chan int) {
			defer wg.Done()

			for msg := range chIn {
				out <- msg
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func squareSequential(input int) int {
	return input * input
}
