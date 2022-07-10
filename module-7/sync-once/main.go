package main

import (
	"fmt"
	"sync"
)

// sync Once ensures that given function will be only run once even when called from multiple go routines
// Do method of once is concurrent safe since it is wrapped in lock unlock already
func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	// var once sync.Once

	load := func() {
		fmt.Println("Run only once initialization function")
	}

	wg.Add(10)
	var done bool
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// once.Do(load)

			mu.Lock()
			if !done {
				load()
				done = true
			}
			mu.Unlock()
		}()
	}
	wg.Wait()
}
