package main

import (
	"log"
	"runtime"
	"sync"
)

//by using mutex lock and unlock it ensures shared memory avoids data race

type hub struct {
	balance int //shared memory
	mu      sync.Mutex
}

func newHub() *hub {
	return &hub{
		balance: 0,
	}
}

func (h *hub) withdraw(amount int) {
	// h.mu.Lock()
	// defer h.mu.Unlock()
	h.balance -= amount
}
func (h *hub) deposit(amount int) {
	// h.mu.Lock()
	// defer h.mu.Unlock()
	h.balance += amount
}
func (h *hub) getBalance() int {
	return h.balance
}

func main() {
	runtime.GOMAXPROCS(4)

	hb := newHub()
	var wg sync.WaitGroup

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			hb.deposit(1)
		}()
	}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			hb.withdraw(1)
		}()
	}

	wg.Wait()
	log.Println(hb.getBalance())
}
