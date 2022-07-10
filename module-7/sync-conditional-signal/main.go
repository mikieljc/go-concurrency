package main

import (
	"log"
	"sync"
	"time"
)

// using the Signal function, it will signal the longest waiting go routine that the event occured and putting the go routine in local runnable state

var sharedRsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	// c := sync.NewCond(&mu)

	go func() {
		defer wg.Done()

		// c.L.Lock()
		mu.Lock()
		for len(sharedRsc) == 0 {
			mu.Unlock()
			time.Sleep(1 * time.Millisecond)
			mu.Lock()
			// c.Wait()
		}

		log.Println(sharedRsc["rsc1"])
		mu.Unlock()
		// c.L.Unlock()
	}()

	// c.L.Lock()
	mu.Lock()
	sharedRsc["rsc1"] = "foo"
	mu.Unlock()
	// c.Signal()
	// c.L.Unlock()

	wg.Wait()
}
