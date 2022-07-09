package main

import (
	"log"
	"time"
)

func main() {
	ch := make(chan string)
	go chanSend(ch)

	//do not block when no received message
	for {
		select {
		case msg := <-ch:
			{
				log.Println("Message received from Channel", msg)
			}
		default:
			{
				log.Println("No message received from Channel, proceeding...")
			}
		}

		//do some proccess
		log.Println("processing...")
		time.Sleep(1500 * time.Millisecond)
	}

}

func chanSend(ch chan string) {
	for {
		time.Sleep(3 * time.Second)
		ch <- "message"
	}
}
