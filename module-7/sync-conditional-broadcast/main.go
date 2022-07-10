package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

//sync conditionals could be useful in situations where multiple readers wait for the shared resources to be available

// Signal function only proccess 1 go routine at any point in time
// Broadcast function process ALL waiting go routine when called

func conditions() bool {
	i := rand.Intn(3)
	if i == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var wg sync.WaitGroup
	var chronologicalProccessed []int
	processedRoutinesTimestamped := make(map[int]time.Duration)

	mu := sync.Mutex{}

	c := sync.NewCond(&mu)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(k int) {
			c.L.Lock()
			defer func() {
				c.L.Unlock()
				wg.Done()
			}()
			//When the condition is not met, enter the sink through the condition variable
			//tcWait() will first release the lock bound to the condition variable, and then enter the sleep state
			timestamp := time.Now()
			for !conditions() {
				c.Wait()
				fmt.Println("waiting go routine", k, "current elapsed time:", time.Since(timestamp))
			}
			chronologicalProccessed = append(chronologicalProccessed, k)
			processedRoutinesTimestamped[k] = time.Since(timestamp)
			log.Println("Updated processed list", processedRoutinesTimestamped)
		}(i)
	}

	go func() {
		for {
			time.Sleep(2 * time.Second)
			// fmt.Println("\n\nBroadcast")
			// c.Broadcast()
			fmt.Println("\n\nSignal")
			c.Signal()
		}

	}()

	wg.Wait()
	log.Println("Chronological Processed Go Routine", chronologicalProccessed)
	log.Println("Whole process elapsed", processedRoutinesTimestamped[chronologicalProccessed[len(chronologicalProccessed)-1]])
}
