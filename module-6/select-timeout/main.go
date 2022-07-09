package main

import (
	"log"
	"time"
)

func main() {
	ch := make(chan string)

	go chan1Send(ch)

	select {
	case msg := <-ch:
		{
			log.Println("Message received from Channel 1", msg)
		}
	case <-time.After(3 * time.Second):
		{
			log.Println("Timed out!")
		}
	}

}

func chan1Send(ch1 chan string) {
	time.Sleep(4 * time.Second)
	ch1 <- "one"
}
