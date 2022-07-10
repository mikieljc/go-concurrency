package main

import (
	"log"
	"runtime"
	"sync"
)

//by using atomic which is concurrent safe, mainly use for primitive operations such as incrementing a counter, we can avoid data race
//lockless operation
//used for atomic operations on counter
// output should be 50000

func main() {
	runtime.GOMAXPROCS(4)

	var counter uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := 0; c < 1000; c++ {
				// atomic.AddUint64(&counter, 1)
				counter++
			}
		}()
	}

	wg.Wait()
	log.Println("counter:" /*atomic.LoadUint64(&counter))*/, counter)
}
