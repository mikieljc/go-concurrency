package main

import (
	"log"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go chan1Send(ch1)
	go chan2Send(ch2)

	//multiplex

	select {
	case msg := <-ch1:
		{
			log.Println("Message received from Channel 1", msg)
		}
	case msg := <-ch2:
		{
			log.Println("Message received from Channel 2", msg)
		}
	}

}

func chan1Send(ch1 chan string) {
	time.Sleep(1 * time.Second)
	ch1 <- "one"
}

func chan2Send(ch2 chan string) {
	time.Sleep(1 * time.Second)
	ch2 <- "two"
}
