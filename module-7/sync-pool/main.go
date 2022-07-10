package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

// sync Pool is usefull for reducing consumption in the perspective of connections (database), memory
// rather than creating new instances of bytes.Buffer for every call of logger function
// we can just reuse

var count = 0

type Pool struct {
	bufPool sync.Pool
}

func NewPool() Pool {
	return Pool{
		bufPool: sync.Pool{
			New: func() interface{} {
				log.Println("Allocating new bytes.Buffer")
				count++
				return new(bytes.Buffer)
			}},
	}
}

func (p *Pool) logger(w io.Writer, val string) {
	b := p.bufPool.Get().(*bytes.Buffer)

	b.Reset()

	b.WriteString(time.Now().Format("15:04:05"))
	b.WriteString(" : ")
	b.WriteString(val)
	b.WriteString("\n")

	w.Write(b.Bytes())

	p.bufPool.Put(b)
}

func main() {

	var wg sync.WaitGroup
	pp := NewPool()

	wg.Add(100)
	timestamp := time.Now()
	for i := 0; i < 100; i++ {
		go func(k int) {
			defer wg.Done()
			pp.logger(os.Stdout, "debug-string"+strconv.Itoa(k))
		}(i)
		time.Sleep(1 * time.Millisecond)
	}

	wg.Wait()
	log.Println("Time elapsed", time.Since(timestamp))
	log.Println("POOL FINAL COUNT", count)
}
